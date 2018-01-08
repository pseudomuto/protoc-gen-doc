package parser

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"path"
	"strconv"
	"strings"
)

const (
	// tag numbers in FileDescriptorProto
	packagePath   = 2  // package
	messagePath   = 4  // message_type
	enumPath      = 5  // enum_type
	servicePath   = 6  // service
	extensionPath = 7  // extension
	syntaxPath    = 12 // syntax

	// tag numbers in DescriptorProto
	messageFieldPath     = 2 // field
	messageMessagePath   = 3 // nested_type
	messageEnumPath      = 4 // enum_type
	messageExtensionPath = 6 // extension
	messageOneOfPath     = 8 // oneOf

	// tag numbers in EnumDescriptorProto
	enumValuePath = 2 // value

	// tag numbers in ServiceDescriptorProto
	serviceMethodPath = 2 // method
)

type extensionContainer interface {
	GetExtension() []*descriptor.FieldDescriptorProto
}

type protoParser interface {
	Comments() *commentExtractor
	Enums() []*Enum
	Extensions() []*Extension
	FileName() string
	IsProto3() bool
	Messages() []*Message
	Package() string
	Services() []*Service
}

func parseProtoFile(fd *descriptor.FileDescriptorProto) *File {
	pp := newProtoParser(fd)

	file := &File{
		parsedObject: parsedObject{
			Name:     pp.FileName(),
			Package:  pp.Package(),
			IsProto3: pp.IsProto3(),
			path:     strconv.Itoa(syntaxPath),
		},
		Enums:      pp.Enums(),
		Extensions: pp.Extensions(),
		Messages:   pp.Messages(),
		Services:   pp.Services(),
	}

	pp.Comments().extractComments(file.getCommentContainers()...)

	return file
}

type protoFileParser struct {
	fd *descriptor.FileDescriptorProto
}

func newProtoParser(fd *descriptor.FileDescriptorProto) protoParser {
	return &protoFileParser{fd: fd}
}

func (pp *protoFileParser) Comments() *commentExtractor {
	return newCommentExtractor(pp.fd)
}

func (pp *protoFileParser) Enums() []*Enum {
	descriptors := pp.fd.GetEnumType()
	enums := make([]*Enum, 0, len(descriptors)+10)

	for idx, enum := range descriptors {
		enums = append(enums, pp.parseEnumDescriptor(enum, nil, idx))
	}

	for _, msg := range pp.Messages() {
		for idx, enum := range msg.enums {
			enums = append(enums, pp.parseEnumDescriptor(enum, msg, idx))
		}
	}

	return enums
}

func (pp *protoFileParser) parseEnumDescriptor(e *descriptor.EnumDescriptorProto, p *Message, idx int) *Enum {
	enum := &Enum{
		parsedObject: parsedObject{
			Name:     e.GetName(),
			Package:  pp.Package(),
			IsProto3: pp.IsProto3(),
			path:     fmt.Sprintf("%d,%d", enumPath, idx),
		},
	}

	if p != nil {
		enum.Name = fmt.Sprintf("%s.%s", p.Name, enum.Name)
		enum.path = fmt.Sprintf("%s,%d,%d", p.path, messageEnumPath, idx)
	}

	enum.Values = pp.parseEnumValues(e.GetValue(), enum)
	return enum
}

func (pp *protoFileParser) parseEnumValues(vd []*descriptor.EnumValueDescriptorProto, enum *Enum) []*EnumValue {
	values := make([]*EnumValue, 0, len(vd))

	for idx, val := range vd {
		values = append(values, &EnumValue{
			parsedObject: parsedObject{
				Name:     val.GetName(),
				Package:  enum.Package,
				IsProto3: enum.IsProto3,
				path:     fmt.Sprintf("%s,%d,%d", enum.path, enumValuePath, idx),
			},
			Number: val.GetNumber(),
		})
	}

	return values
}

func (pp *protoFileParser) Extensions() []*Extension {
	return pp.parseExtensions(pp.fd, "", strconv.Itoa(extensionPath))
}

func (pp *protoFileParser) FileName() string {
	return path.Base(pp.fd.GetName())
}

func (pp *protoFileParser) IsProto3() bool {
	return pp.fd.GetSyntax() == "proto3"
}

func (pp *protoFileParser) Messages() []*Message {
	descriptors := pp.fd.GetMessageType()
	messages := make([]*Message, 0, len(descriptors))

	for idx, msg := range descriptors {
		messages = pp.parseDescriptor(messages, msg, nil, idx)
	}

	return messages
}

func (pp *protoFileParser) parseDescriptor(sl []*Message, d *descriptor.DescriptorProto, p *Message, idx int) []*Message {
	sl = append(sl, &Message{
		parsedObject: parsedObject{
			Name:     d.GetName(),
			Package:  pp.Package(),
			IsProto3: pp.IsProto3(),
			path:     fmt.Sprintf("%d,%d", messagePath, idx),
		},
		enums: d.GetEnumType(),
	})

	this := sl[len(sl)-1]
	if p != nil { // nested?
		this.Name = fmt.Sprintf("%s.%s", p.Name, this.Name)
		this.path = fmt.Sprintf("%s,%d,%d", p.path, messageMessagePath, idx)
	}

	this.Fields = pp.parseFields(d, d.GetField(), this.path)
	this.Extensions = pp.parseExtensions(d, this.FullName(), fmt.Sprintf("%s,%d", this.path, messageExtensionPath))

	// parse nested message types
	for i, nested := range d.GetNestedType() {
		sl = pp.parseDescriptor(sl, nested, this, i)
	}

	return sl
}

func (pp *protoFileParser) parseFields(fc *descriptor.DescriptorProto, fdp []*descriptor.FieldDescriptorProto, basePath string) []*Field {
	fields := make([]*Field, 0, len(fdp))

	for idx, field := range fdp {
		fields = append(fields, &Field{
			parsedObject: parsedObject{
				Name:     field.GetName(),
				Package:  pp.Package(),
				IsProto3: pp.IsProto3(),
				path:     fmt.Sprintf("%s,%d,%d", basePath, messageFieldPath, idx),
			},
			Label: pp.labelName(field.GetLabel()),
			Type:  fmt.Sprintf("%s.%s", pp.Package(), pp.typeName(field.GetTypeName())),
		})

		f := fields[len(fields)-1]

		if(field.OneofIndex != nil){
			f.OneOf = &OneOf{
				parsedObject: parsedObject{
					Name:     fc.GetOneofDecl()[(*field.OneofIndex)].GetName(),
					Package:  pp.Package(),
					IsProto3: pp.IsProto3(),
					path:     fmt.Sprintf("%s,%d,%d", basePath, messageOneOfPath, *field.OneofIndex),
				},
			}
		}

		if f.Type == pp.Package()+"." {
			f.Type = pp.typeName(field.GetType().String())
		}
	}

	return fields
}


func (pp *protoFileParser) typeName(name string) string {
	if strings.HasPrefix(name, ".") {
		return strings.TrimPrefix(name, fmt.Sprintf(".%s.", pp.Package()))
	}

	return strings.ToLower(strings.TrimPrefix(name, "TYPE_"))
}

func (pp *protoFileParser) parseExtensions(ec extensionContainer, scopeType, basePath string) []*Extension {
	descriptors := ec.GetExtension()
	extensions := make([]*Extension, 0, len(descriptors))

	for idx, ext := range descriptors {
		extensions = append(extensions, &Extension{
			Field: Field{
				parsedObject: parsedObject{
					Name:     ext.GetName(),
					Package:  pp.Package(),
					IsProto3: pp.IsProto3(),
					path:     fmt.Sprintf("%s,%d", basePath, idx),
				},
				DefaultValue: ext.GetDefaultValue(),
				Type:         pp.typeName(ext.GetTypeName()),
			},
			ContainingType: strings.TrimPrefix(ext.GetExtendee(), "."),
			ScopeType:      scopeType,
			Label:          pp.labelName(ext.GetLabel()),
			Number:         ext.GetNumber(),
		})

		e := extensions[len(extensions)-1]
		if e.Type == "" {
			e.Type = pp.typeName(ext.GetType().String())
		}
	}

	return extensions
}

func (pp *protoFileParser) labelName(fd descriptor.FieldDescriptorProto_Label) string {
	if pp.IsProto3() && fd != descriptor.FieldDescriptorProto_LABEL_REPEATED {
		return ""
	}

	return strings.ToLower(strings.TrimPrefix(fd.String(), "LABEL_"))
}

func (pp *protoFileParser) Package() string {
	return pp.fd.GetPackage()
}

func (pp *protoFileParser) Services() []*Service {
	descriptors := pp.fd.GetService()
	services := make([]*Service, 0, len(descriptors))

	for idx, descriptor := range descriptors {
		path := fmt.Sprintf("%d,%d", servicePath, idx)

		services = append(services, &Service{
			parsedObject: parsedObject{
				Name:     descriptor.GetName(),
				Package:  pp.Package(),
				IsProto3: pp.IsProto3(),
				path:     path,
			},
			Methods: pp.parseServiceMethods(descriptor, path),
		})
	}

	return services
}

func (pp *protoFileParser) parseServiceMethods(sd *descriptor.ServiceDescriptorProto, basePath string) []*ServiceMethod {
	descriptors := sd.GetMethod()
	methods := make([]*ServiceMethod, 0, len(descriptors))

	for idx, method := range descriptors {
		methods = append(methods, &ServiceMethod{
			parsedObject: parsedObject{
				Name:     method.GetName(),
				Package:  pp.Package(),
				IsProto3: pp.IsProto3(),
				path:     fmt.Sprintf("%s,%d,%d", basePath, serviceMethodPath, idx),
			},
			ClientStreaming: method.GetClientStreaming(),
			ServerStreaming: method.GetServerStreaming(),
			RequestType:     strings.TrimPrefix(method.GetInputType(), "."),
			ResponseType:    strings.TrimPrefix(method.GetOutputType(), "."),
		})
	}

	return methods
}
