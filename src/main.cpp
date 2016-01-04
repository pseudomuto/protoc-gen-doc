/*
  Copyright 2014, Elvis Stansvik

  Redistribution and use in source and binary forms, with or without modification,
  are permitted provided that the following conditions are met:

    Redistributions of source code must retain the above copyright notice,
    this list of conditions and the following disclaimer.

    Redistributions in binary form must reproduce the above copyright notice,
    this list of conditions and the following disclaimer in the documentation
    and/or other materials provided with the distribution.
*/

#include "mustache.h"

#include <algorithm>
#include <iostream>
#include <string>

#include <QDir>
#include <QFile>
#include <QFileInfo>
#include <QFileInfoList>
#include <QIODevice>
#include <QJsonDocument>
#include <QJsonArray>
#include <QRegularExpression>
#include <QString>
#include <QStringList>
#include <QVariant>
#include <QVariantHash>
#include <QVariantList>

#include <google/protobuf/compiler/plugin.h>
#include <google/protobuf/compiler/code_generator.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/io/zero_copy_stream.h>
#include <google/protobuf/io/printer.h>

namespace gp = google::protobuf;
namespace ms = Mustache;

/**
 * Context class for the documentation generator.
 */
class DocGeneratorContext {
public:
    QString template_;      /**< Mustache template. */
    QString outputFileName; /**< Output filename. */
    QVariantList files;     /**< List of files to render. */
};

/**
 * Returns the "long" name of the message or enum described by @p descriptor.
 *
 * The long name is the name of the message or enum itself, preceeded by the
 * names of its enclosing types, separated by dots. E.g. for "Baz" it could be
 * "Foo.Bar.Baz".
 */
template<typename T>
static QString longName(const T *descriptor)
{
    if (!descriptor) {
        return QString();
    } else if (!descriptor->containing_type()) {
        return QString::fromStdString(descriptor->name());
    }
    return longName(descriptor->containing_type()) + "." +
            QString::fromStdString(descriptor->name());
}

/**
 * Returns true if the variant @p v1 is less than @p v2.
 *
 * It is assumed that the variants each contain a QVariantHash with either
 * a "message_long_name" or "message_long_name" key. This comparator is used
 * when sorting the message and enum lists.
 */
static inline bool longNameLessThan(const QVariant &v1, const QVariant &v2)
{
    if (v1.toHash()["message_long_name"].toString() < v2.toHash()["message_long_name"].toString())
        return true;
    if (v1.toHash()["enum_long_name"].toString() < v2.toHash()["enum_long_name"].toString())
        return true;
    return v1.toHash()["extension_long_name"].toString() < v2.toHash()["extension_long_name"].toString();
}

/**
 * Returns the description of a message, enum, enum value or field.
 *
 * The description is taken as the leading comments followed by the trailing
 * comments. If present, a single space is removed from the start of each line.
 * Whitespace is trimmed from the final result before it is returned.
 */
template<typename T>
static QString descriptionOf(const T *descriptor)
{
    QString description;

    gp::SourceLocation sourceLocation;
    descriptor->GetSourceLocation(&sourceLocation);

    // Check for leading documentation comments.
    QString leading = QString::fromStdString(sourceLocation.leading_comments);
    if (leading.startsWith('*') || leading.startsWith('/')) {
        leading = leading.mid(1);
        leading.replace(QRegularExpression("^ ", QRegularExpression::MultilineOption), "");
        description += leading;
    }

    // Check for trailing documentation comments.
    QString trailing = QString::fromStdString(sourceLocation.trailing_comments);
    if (trailing.startsWith('*') || trailing.startsWith('/')) {
        trailing = trailing.mid(1);
        trailing.replace(QRegularExpression("^ ", QRegularExpression::MultilineOption), "");
        description += trailing;
    }

    // Return trimmed result.
    return description.trimmed();
}

/**
 * Returns the name of the scalar field type @p type.
 */
static QString scalarTypeName(gp::FieldDescriptor::Type type)
{
    switch (type) {
        case gp::FieldDescriptor::TYPE_BOOL:
            return "bool";
        case gp::FieldDescriptor::TYPE_BYTES:
            return "bytes";
        case gp::FieldDescriptor::TYPE_DOUBLE:
            return "double";
        case gp::FieldDescriptor::TYPE_FIXED32:
            return "fixed32";
        case gp::FieldDescriptor::TYPE_FIXED64:
            return "fixed64";
        case gp::FieldDescriptor::TYPE_FLOAT:
            return "float";
        case gp::FieldDescriptor::TYPE_INT32:
            return "int32";
        case gp::FieldDescriptor::TYPE_INT64:
            return "int64";
        case gp::FieldDescriptor::TYPE_SFIXED32:
            return "sfixed32";
        case gp::FieldDescriptor::TYPE_SFIXED64:
            return "sfixed64";
        case gp::FieldDescriptor::TYPE_SINT32:
            return "sint32";
        case gp::FieldDescriptor::TYPE_SINT64:
            return "sint64";
        case gp::FieldDescriptor::TYPE_STRING:
            return "string";
        case gp::FieldDescriptor::TYPE_UINT32:
            return "uint32";
        case gp::FieldDescriptor::TYPE_UINT64:
            return "uint64";
        default:
            return "<unknown>";
    }
}

/**
 * Returns the name of the field label @p label.
 */
static QString labelName(gp::FieldDescriptor::Label label)
{
    switch(label) {
        case gp::FieldDescriptor::LABEL_OPTIONAL:
            return "optional";
        case gp::FieldDescriptor::LABEL_REPEATED:
            return "repeated";
        case gp::FieldDescriptor::LABEL_REQUIRED:
            return "required";
        default:
            return "<unknown>";
    }
}

/**
 * Add field to variant list.
 *
 * Adds the field described by @p fieldDescriptor to the variant list @p fields.
 */
static void addField(const gp::FieldDescriptor *fieldDescriptor, QVariantList *fields)
{
    QString description = descriptionOf(fieldDescriptor);

    if (description.startsWith("@exclude")) {
        return;
    }

    QVariantHash field;

    // Add basic info.
    field["field_name"] = QString::fromStdString(fieldDescriptor->name());
    field["field_description"] = description;
    field["field_label"] = labelName(fieldDescriptor->label());

    // Add type information.
    gp::FieldDescriptor::Type type = fieldDescriptor->type();
    if (type == gp::FieldDescriptor::TYPE_MESSAGE || type == gp::FieldDescriptor::TYPE_GROUP) {
        // Field is of message / group type.
        const gp::Descriptor *descriptor = fieldDescriptor->message_type();
        field["field_type"] = QString::fromStdString(descriptor->name());
        field["field_long_type"] = longName(descriptor);
        field["field_full_type"] = QString::fromStdString(descriptor->full_name());
    } else if (type == gp::FieldDescriptor::TYPE_ENUM) {
        // Field is of enum type.
        const gp::EnumDescriptor *descriptor = fieldDescriptor->enum_type();
        field["field_type"] = QString::fromStdString(descriptor->name());
        field["field_long_type"] = longName(descriptor);
        field["field_full_type"] = QString::fromStdString(descriptor->full_name());
    } else {
        // Field is of scalar type.
        QString typeName(scalarTypeName(type));
        field["field_type"] = typeName;
        field["field_long_type"] = typeName;
        field["field_full_type"] = typeName;
    }

    fields->append(field);
}

/**
 * Add extension to variant list.
 *
 * Adds the extension described by @p fieldDescriptor to the variant list @p extensions.
 */
static void addExtension(const gp::FieldDescriptor *fieldDescriptor, QVariantList *extensions)
{
    QString description = descriptionOf(fieldDescriptor);

    if (description.startsWith("@exclude")) {
        return;
    }

    QVariantHash extension;

    // Add basic info.
    extension["extension_name"] = QString::fromStdString(fieldDescriptor->name());
    extension["extension_full_name"] = QString::fromStdString(fieldDescriptor->full_name());
    extension["extension_long_name"] = longName(fieldDescriptor);
    extension["extension_description"] = description;
    extension["extension_label"] = labelName(fieldDescriptor->label());
    extension["extension_number"] = QString::number(fieldDescriptor->number());

    if (fieldDescriptor->is_extension()) {
        const gp::Descriptor *descriptor = fieldDescriptor->extension_scope();
        if (descriptor != NULL) {
            extension["extension_scope_type"] = QString::fromStdString(descriptor->name());
            extension["extension_scope_long_type"] = longName(descriptor);
            extension["extension_scope_full_type"] = QString::fromStdString(descriptor->full_name());
        }

        descriptor = fieldDescriptor->containing_type();
        if (descriptor != NULL) {
            extension["extension_containing_type"] = QString::fromStdString(descriptor->name());
            extension["extension_containing_long_type"] = longName(descriptor);
            extension["extension_containing_full_type"] = QString::fromStdString(descriptor->full_name());
        }
    }

    // Add type information.
    gp::FieldDescriptor::Type type = fieldDescriptor->type();
    if (type == gp::FieldDescriptor::TYPE_MESSAGE || type == gp::FieldDescriptor::TYPE_GROUP) {
        // Field is of message / group type.
        const gp::Descriptor *descriptor = fieldDescriptor->message_type();
        extension["extension_type"] = QString::fromStdString(descriptor->name());
        extension["extension_long_type"] = longName(descriptor);
        extension["extension_full_type"] = QString::fromStdString(descriptor->full_name());
    } else if (type == gp::FieldDescriptor::TYPE_ENUM) {
        // Field is of enum type.
        const gp::EnumDescriptor *descriptor = fieldDescriptor->enum_type();
        extension["extension_type"] = QString::fromStdString(descriptor->name());
        extension["extension_long_type"] = longName(descriptor);
        extension["extension_full_type"] = QString::fromStdString(descriptor->full_name());
    } else {
        // Field is of scalar type.
        QString typeName(scalarTypeName(type));
        extension["extension_type"] = typeName;
        extension["extension_long_type"] = typeName;
        extension["extension_full_type"] = typeName;
    }

    extensions->append(extension);
}

/**
 * Adds the enum described by @p enumDescriptor to the variant list @p enums.
 */
static void addEnum(const gp::EnumDescriptor *enumDescriptor, QVariantList *enums)
{
    QString description = descriptionOf(enumDescriptor);

    if (description.startsWith("@exclude")) {
        return;
    }

    QVariantHash enum_;

    // Add basic info.
    enum_["enum_name"] = QString::fromStdString(enumDescriptor->name());
    enum_["enum_long_name"] = longName(enumDescriptor);
    enum_["enum_full_name"] = QString::fromStdString(enumDescriptor->full_name());
    enum_["enum_description"] = description;

    // Add enum values.
    QVariantList values;
    for (int i = 0; i < enumDescriptor->value_count(); ++i) {
        const gp::EnumValueDescriptor *valueDescriptor = enumDescriptor->value(i);

        QString description = descriptionOf(valueDescriptor);

        if (description.startsWith("@exclude")) {
            continue;
        }

        QVariantHash value;
        value["value_name"] = QString::fromStdString(valueDescriptor->name());
        value["value_number"] = valueDescriptor->number();
        value["value_description"] = description;
        values.append(value);
    }
    enum_["enum_values"] = values;

    enums->append(enum_);
}

/**
 * Add messages to variant list.
 *
 * Adds the message described by @p descriptor and all its nested messages and
 * enums to the variant list @p messages and @p enums, respectively.
 */
static void addMessages(const gp::Descriptor *descriptor,
                        QVariantList *messages,
                        QVariantList *enums)
{
    QString description = descriptionOf(descriptor);

    if (description.startsWith("@exclude")) {
        return;
    }

    QVariantHash message;

    // Add basic info.
    message["message_name"] = QString::fromStdString(descriptor->name());
    message["message_long_name"] = longName(descriptor);
    message["message_full_name"] = QString::fromStdString(descriptor->full_name());
    message["message_description"] = description;

    // Add fields.
    QVariantList fields;
    for (int i = 0; i < descriptor->field_count(); ++i) {
        addField(descriptor->field(i), &fields);
    }
    message["message_fields"] = fields;

    // Add inlined extensions
    QVariantList extensions;
    for (int i = 0; i < descriptor->extension_count(); ++i) {
        addExtension(descriptor->extension(i), &extensions);
    }
    message["message_extensions"] = extensions;


    messages->append(message);

    // Add nested messages and enums.
    for (int i = 0; i < descriptor->nested_type_count(); ++i) {
        addMessages(descriptor->nested_type(i), messages, enums);
    }
    for (int i = 0; i < descriptor->enum_type_count(); ++i) {
        addEnum(descriptor->enum_type(i), enums);
    }
}

/**
 * Add file to variant list.
 *
 * Adds the file described by @p fileDescriptor to the variant list @p files.
 */
static void addFile(const gp::FileDescriptor *fileDescriptor, QVariantList *files)
{
    QVariantHash file;

    // Add basic info.
    file["file_name"] = QFileInfo(QString::fromStdString(fileDescriptor->name())).fileName();
    file["file_package"] = QString::fromStdString(fileDescriptor->package());

    QVariantList messages;
    QVariantList enums;
    QVariantList extensions;

    // Add messages.
    for (int i = 0; i < fileDescriptor->message_type_count(); ++i) {
        addMessages(fileDescriptor->message_type(i), &messages, &enums);
    }
    std::sort(messages.begin(), messages.end(), &longNameLessThan);
    file["file_messages"] = messages;

    // Add enums.
    for (int i = 0; i < fileDescriptor->enum_type_count(); ++i) {
        addEnum(fileDescriptor->enum_type(i), &enums);
    }
    std::sort(enums.begin(), enums.end(), &longNameLessThan);
    file["file_enums"] = enums;

    // Add file-level extensions
    for (int i = 0; i < fileDescriptor->extension_count(); ++i) {
        addExtension(fileDescriptor->extension(i), &extensions);
    }
    std::sort(extensions.begin(), extensions.end(), &longNameLessThan);
    file["file_extensions"] = extensions;

    files->append(file);
}

/**
 * Return a formatted template rendering error.
 *
 * @param template_ Template in which the error occurred.
 * @param renderer Template renderer that failed.
 * @return Formatted single-line error.
 */
static std::string formattedError(const QString &template_, const ms::Renderer &renderer)
{
    QString location = template_;
    if (!renderer.errorPartial().isEmpty()) {
        location += " in partial " + renderer.errorPartial();
    }
    return QString("%1:%2: %3")
            .arg(location)
            .arg(renderer.errorPos())
            .arg(renderer.error()).toStdString();
}

/**
 * Returns the list of formats that are supported out of the box.
 */
static QStringList supportedFormats()
{
    QStringList formats;
    QStringList filter = QStringList() << "*.mustache";
    QFileInfoList entries = QDir(":/templates").entryInfoList(filter);
    for (const QFileInfo &entry : entries) {
        formats.append(entry.baseName());
    }
    return formats;
}

/**
 * Returns the template specified by @p name.
 *
 * The @p name parameter may be either a template file name, or the name of a
 * supported format ("html", "docbook", ...). If an error occured, @p error is
 * set to point to an error message and QString() returned.
 */
static QString readTemplate(const QString &name, std::string *error)
{
    QString builtInFileName = QString(":/templates/%1.mustache").arg(name);
    QString fileName = supportedFormats().contains(name) ? builtInFileName : name;
    QFile file(fileName);

    if (!file.open(QIODevice::ReadOnly)) {
        *error = QString("%1: %2").arg(fileName).arg(file.errorString()).toStdString();
        return QString();
    } else {
        return file.readAll();
    }
}

/**
 * Parses the plugin parameter string.
 *
 * @param parameter Plugin parameter string.
 * @param generatorContext Documentation generator context to parse into.
 * @param error Pointer to error if parsing failed.
 * @return true on success, otherwise false.
 */
static bool parseParameter(const std::string &parameter,
                           DocGeneratorContext &generatorContext,
                           std::string *error)
{
    QStringList tokens = QString::fromStdString(parameter).split(",");

    if (tokens.size() != 2) {
        QString usage("Usage: --doc_out=%1|<TEMPLATE_FILENAME>,<OUT_FILENAME>:<OUT_DIR>");
        *error = QString(usage).arg(supportedFormats().join("|")).toStdString();
        return false;
    }
    generatorContext.template_ = readTemplate(tokens.at(0), error);
    generatorContext.outputFileName = tokens.at(1);

    return true;
}

/**
 * Template filter for breaking paragraphs into HTML `<p>` elements.
 *
 * Renders @p text with @p renderer in @p context and returns the result with
 * paragraphs enclosed in `<p>..</p>`.
 *
 */
static QString pFilter(const QString &text, ms::Renderer* renderer, ms::Context* context)
{
    QRegularExpression re("(\\n|\\r|\\r\\n)\\s*(\\n|\\r|\\r\\n)");
    return "<p>" + renderer->render(text, context).split(re).join("</p><p>") + "</p>";
}

/**
 * Template filter for breaking paragraphs into DocBook `<para>` elements.
 *
 * Renders @p text with @p renderer in @p context and returns the result with
 * paragraphs enclosed in `<para>..</para>`.
 *
 */
static QString paraFilter(const QString &text, ms::Renderer* renderer, ms::Context* context)
{
    QRegularExpression re("(\\n|\\r|\\r\\n)\\s*(\\n|\\r|\\r\\n)");
    return "<para>" + renderer->render(text, context).split(re).join("</para><para>") + "</para>";
}

/**
 * Template filter for removing line breaks.
 *
 * Renders @p text with @p renderer in @p context and returns the result with
 * all occurrances of `\r\n`, `\n`, `\r` removed in that order.
 */
static QString nobrFilter(const QString &text, ms::Renderer* renderer, ms::Context* context)
{
    QString result = renderer->render(text, context);
    result.remove("\r\n");
    result.remove("\r");
    result.remove("\n");
    return result;
}

/**
 * Renders the list of files.
 *
 * Renders files in the @p generatorContext to the directory specified in
 * @p context. If an error occurred, @p error is set to point to an error
 * message and no output is written.
 *
 * @param generatorContext Documentation generator context.
 * @param context Compiler generator context specifying the output directory.
 * @param error Pointer to error if rendering failed.
 * @return true on success, otherwise false.
 */
static bool render(const DocGeneratorContext &generatorContext,
                   gp::compiler::GeneratorContext *context, std::string *error)
{
    QVariantHash args;

    // Add filters.
    args["p"] = QVariant::fromValue(ms::QtVariantContext::fn_t(pFilter));
    args["para"] = QVariant::fromValue(ms::QtVariantContext::fn_t(paraFilter));
    args["nobr"] = QVariant::fromValue(ms::QtVariantContext::fn_t(nobrFilter));

    // Add files list.
    args["files"] = generatorContext.files;

    // Add scalar value types table.
    QString fileName(":/templates/scalar_value_types.json");
    QFile file(fileName);
    if (!file.open(QIODevice::ReadOnly)) {
        *error = QString("%1: %2").arg(fileName).arg(file.errorString()).toStdString();
        return false;
    }
    QJsonDocument document(QJsonDocument::fromJson(file.readAll()));
    args["scalar_value_types"] = document.array().toVariantList();

    // Render template.
    ms::Renderer renderer;
    ms::QtVariantContext variantContext(args);
    QString result = renderer.render(generatorContext.template_, &variantContext);

    // Check for errors.
    if (!renderer.error().isEmpty()) {
        *error = formattedError(generatorContext.template_, renderer);
        return false;
    }

    // Write output.
    std::string outputFileName = generatorContext.outputFileName.toStdString();
    gp::io::ZeroCopyOutputStream *stream = context->Open(outputFileName);
    gp::io::Printer printer(stream, '$');
    printer.PrintRaw(result.toStdString());

    return true;
}

/// Documentation generator context instance.
static DocGeneratorContext generatorContext;

/**
 * Documentation generator class.
 */
class DocGenerator : public gp::compiler::CodeGenerator
{
    /// Implements google::protobuf::compiler::CodeGenerator.
    bool Generate(
            const gp::FileDescriptor *fileDescriptor,
            const std::string &parameter,
            gp::compiler::GeneratorContext *context,
            std::string *error) const
    {
        addFile(fileDescriptor, &generatorContext.files);

        // Don't render until last file has been parsed.
        std::vector<const gp::FileDescriptor *> parsedFiles;
        context->ListParsedFiles(&parsedFiles);
        if (fileDescriptor != parsedFiles.back()) {
            return true;
        }

        // Parse the plugin parameter.
        if (!parseParameter(parameter, generatorContext, error)) {
            return false;
        }

        // Render files
        if (!render(generatorContext, context, error)) {
            return false;
        }

        return true;
    }
};

int main(int argc, char *argv[])
{
    // Instantiate and invoke the generator plugin.
    DocGenerator generator;
    return google::protobuf::compiler::PluginMain(argc, argv, &generator);
}
