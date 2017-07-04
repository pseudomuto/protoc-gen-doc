package protoc_gen_doc

import (
	"fmt"
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"strings"
)

type Template struct {
	Files   []*File        `json:"files"`
	Scalars []*ScalarValue `json:"scalar_value_types"`
}

func NewTemplate(pr *parser.ParseResult) *Template {
	files := make([]*File, 0, len(pr.Files))

	for _, f := range pr.Files {
		file := &File{
			Name:          f.Name,
			Description:   f.Comment,
			Package:       f.Package,
			HasEnums:      len(f.Enums) > 0,
			HasExtensions: len(f.Extensions) > 0,
			HasMessages:   len(f.Messages) > 0,
			HasServices:   len(f.Services) > 0,
		}

		for _, e := range f.Enums {
			file.Enums = append(file.Enums, parseEnum(e))
		}

		for _, e := range f.Extensions {
			file.Extensions = append(file.Extensions, parseFileExtension(e))
		}

		for _, m := range f.Messages {
			file.Messages = append(file.Messages, parseMessage(m))
		}

		for _, s := range f.Services {
			file.Services = append(file.Services, parseService(s))
		}

		files = append(files, file)
	}

	return &Template{Files: files, Scalars: makeScalars()}
}

type File struct {
	Name        string `json:"file_name"`
	Description string `json:"file_description"`
	Package     string `json:"file_package"`

	HasEnums      bool `json:"file_has_enums"`
	HasExtensions bool `json:"file_has_extensions"`
	HasMessages   bool `json:"file_has_messages"`
	HasServices   bool `json:"file_has_services"`

	Enums      []*Enum          `json:"file_enums"`
	Extensions []*FileExtension `json:"file_extensions"`
	Messages   []*Message       `json:"file_messages"`
	Services   []*Service       `json:"file_services"`
}

type FileExtension struct {
	Name               string `json:"extension_name"`
	LongName           string `json:"extension_long_name"`
	FullName           string `json:"extension_full_name"`
	Description        string `json:"extension_description"`
	Label              string `json:"extension_label"`
	Type               string `json:"extension_type"`
	LongType           string `json:"extension_long_type"`
	FullType           string `json:"extension_full_type"`
	Number             int    `json:"extension_number"`
	DefaultValue       string `json:"extension_default_value"`
	ContainingType     string `json:"extension_containing_type"`
	ContainingLongType string `json:"extension_containing_long_type"`
	ContainingFullType string `json:"extension_containing_full_type"`
}

type Message struct {
	Name        string `json:"message_name"`
	LongName    string `json:"message_long_name"`
	FullName    string `json:"message_full_name"`
	Description string `json:"message_description"`

	HasExtensions bool `json:"message_has_extensions"`
	HasFields     bool `json:"message_has_extensions"`

	Extensions []*MessageExtension `json:"message_extensions"`
	Fields     []*MessageField     `json:"message_fields"`
}

type MessageField struct {
	Name         string `json:"field_name"`
	Description  string `json:"field_description"`
	Label        string `json:"field_label"`
	Type         string `json:"field_type"`
	LongType     string `json:"field_long_type"`
	FullType     string `json:"field_full_type"`
	DefaultValue string `json:"field_default_value"`
}

type MessageExtension struct {
	FileExtension

	ScopeType     string `json:"extension_scope_type"`
	ScopeLongType string `json:"extension_scope_long_type"`
	ScopeFullType string `json:"extension_scope_full_type"`
}

type Enum struct {
	Name        string       `json:"enum_name"`
	LongName    string       `json:"enum_long_name"`
	FullName    string       `json:"enum_full_name"`
	Description string       `json:"enum_description"`
	Values      []*EnumValue `json:"enum_values"`
}

type EnumValue struct {
	Name        string `json:"value_name"`
	Number      string `json:"value_number"`
	Description string `json:"value_description"`
}

type Service struct {
	Name        string           `json:"service_name"`
	LongName    string           `json:"service_long_name"`
	FullName    string           `json:"service_full_name"`
	Description string           `json:"service_description"`
	Methods     []*ServiceMethod `json:"service_methods"`
}

type ServiceMethod struct {
	Name             string `json:"method_name"`
	Description      string `json:"method_description"`
	RequestType      string `json:"method_request_type"`
	RequestLongType  string `json:"method_request_long_type"`
	RequestFullType  string `json:"method_request_full_type"`
	ResponseType     string `json:"method_response_type"`
	ResponseLongType string `json:"method_response_long_type"`
	ResponseFullType string `json:"method_response_full_type"`
}

type ScalarValue struct {
	ProtoType  string `json:"scalar_value_proto_type"`
	Notes      string `json:"scalar_value_notes"`
	CppType    string `json:"scalar_value_cpp_type"`
	CSharp     string `json:"scalar_value_cs_type"`
	GoType     string `json:"scalar_value_go_type"`
	JavaType   string `json:"scalar_value_java_type"`
	PhpType    string `json:"scalar_value_php_type"`
	PythonType string `json:"scalar_value_python_type"`
	RubyType   string `json:"scalar_value_ruby_type"`
}

func parseEnum(pe *parser.Enum) *Enum {
	enum := &Enum{
		Name:        baseName(pe.Name),
		LongName:    strings.TrimPrefix(pe.FullName(), pe.Package+"."),
		FullName:    pe.FullName(),
		Description: pe.Comment,
	}

	for _, val := range pe.Values {
		enum.Values = append(enum.Values, &EnumValue{
			Name:        val.Name,
			Number:      fmt.Sprint(val.Number),
			Description: val.Comment,
		})
	}

	return enum
}

func parseFileExtension(pe *parser.Extension) *FileExtension {
	return &FileExtension{
		Name:               baseName(pe.Name),
		LongName:           strings.TrimPrefix(pe.FullName(), pe.Package+"."),
		FullName:           pe.FullName(),
		Description:        pe.Comment,
		Label:              pe.Label,
		Type:               baseName(pe.Type),
		LongType:           strings.TrimPrefix(pe.Type, pe.Package+"."),
		FullType:           pe.Type,
		Number:             int(pe.Number),
		DefaultValue:       pe.DefaultValue,
		ContainingType:     baseName(pe.ContainingType),
		ContainingLongType: strings.TrimPrefix(pe.ContainingType, pe.Package+"."),
		ContainingFullType: pe.ContainingType,
	}
}

func parseMessage(pm *parser.Message) *Message {
	msg := &Message{
		Name:          baseName(pm.Name),
		LongName:      pm.Name,
		FullName:      pm.FullName(),
		Description:   pm.Comment,
		HasExtensions: len(pm.Extensions) > 0,
		HasFields:     len(pm.Fields) > 0,
	}

	for _, ext := range pm.Extensions {
		msg.Extensions = append(msg.Extensions, parseMessageExtension(ext))
	}

	for _, f := range pm.Fields {
		msg.Fields = append(msg.Fields, parseMessageField(f))
	}

	return msg
}

func parseMessageExtension(pe *parser.Extension) *MessageExtension {
	return &MessageExtension{
		FileExtension: *parseFileExtension(pe),
		ScopeType:     baseName(pe.ScopeType),
		ScopeLongType: strings.TrimPrefix(pe.ScopeType, pe.Package+"."),
		ScopeFullType: pe.ScopeType,
	}
}

func parseMessageField(pf *parser.Field) *MessageField {
	return &MessageField{
		Name:         pf.Name,
		Description:  pf.Comment,
		Label:        pf.Label,
		Type:         baseName(pf.Type),
		LongType:     strings.TrimPrefix(pf.Type, pf.Package+"."),
		FullType:     pf.Type,
		DefaultValue: pf.DefaultValue,
	}
}

func parseService(ps *parser.Service) *Service {
	service := &Service{
		Name:        ps.Name,
		LongName:    ps.Name,
		FullName:    fmt.Sprintf("%s.%s", ps.Package, ps.Name),
		Description: ps.Comment,
	}

	for _, sm := range ps.Methods {
		service.Methods = append(service.Methods, parseServiceMethod(sm))
	}

	return service
}

func parseServiceMethod(pm *parser.ServiceMethod) *ServiceMethod {
	return &ServiceMethod{
		Name:             pm.Name,
		Description:      pm.Comment,
		RequestType:      baseName(pm.RequestType),
		RequestLongType:  strings.TrimPrefix(pm.RequestType, pm.Package+"."),
		RequestFullType:  pm.RequestType,
		ResponseType:     baseName(pm.ResponseType),
		ResponseLongType: strings.TrimPrefix(pm.ResponseType, pm.Package+"."),
		ResponseFullType: pm.ResponseType,
	}
}

func baseName(name string) string {
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
}

func makeScalars() []*ScalarValue {
	return []*ScalarValue{
		{
			"double",
			"",
			"double",
			"double",
			"float64",
			"double",
			"float",
			"float",
			"Float",
		},
		{
			"float",
			"",
			"float",
			"float",
			"float32",
			"float",
			"float",
			"float",
			"Float",
		},
		{
			"int32",
			"Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead.",
			"int32",
			"int",
			"int32",
			"int",
			"integer",
			"int",
			"Bignum or Fixnum (as required)",
		},
		{
			"int64",
			"Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead.",
			"int64",
			"long",
			"int64",
			"long",
			"integer/string",
			"int/long",
			"Bignum",
		},
		{
			"uint32",
			"Uses variable-length encoding.",
			"uint32",
			"uint",
			"uint32",
			"int",
			"integer",
			"int/long",
			"Bignum or Fixnum (as required)",
		},
		{
			"uint64",
			"Uses variable-length encoding.",
			"uint64",
			"ulong",
			"uint64",
			"long",
			"integer/string",
			"int/long",
			"Bignum or Fixnum (as required)",
		},
		{
			"sint32",
			"Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.",
			"int32",
			"int",
			"int32",
			"int",
			"integer",
			"int",
			"Bignum or Fixnum (as required)",
		},
		{
			"sint64",
			"Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.",
			"int64",
			"long",
			"int64",
			"long",
			"integer/string",
			"int/long",
			"Bignum",
		},
		{
			"fixed32",
			"Always four bytes. More efficient than uint32 if values are often greater than 2^28.",
			"uint32",
			"uint",
			"uint32",
			"int",
			"integer",
			"int",
			"Bignum or Fixnum (as required)",
		},
		{
			"fixed64",
			"Always eight bytes. More efficient than uint64 if values are often greater than 2^56.",
			"uint64",
			"ulong",
			"uint64",
			"long",
			"integer/string",
			"int/long",
			"Bignum",
		},
		{
			"sfixed32",
			"Always four bytes.",
			"int32",
			"int",
			"int32",
			"int",
			"integer",
			"int",
			"Bignum or Fixnum (as required)",
		},
		{
			"sfixed64",
			"Always eight bytes.",
			"int64",
			"long",
			"int64",
			"long",
			"integer/string",
			"int/long",
			"Bignum",
		},
		{
			"bool",
			"",
			"bool",
			"bool",
			"bool",
			"boolean",
			"boolean",
			"boolean",
			"TrueClass/FalseClass",
		},
		{
			"string",
			"A string must always contain UTF-8 encoded or 7-bit ASCII text.",
			"string",
			"string",
			"string",
			"String",
			"string",
			"str/unicode",
			"String (UTF-8)",
		},
		{
			"bytes",
			"May contain any arbitrary sequence of bytes.",
			"string",
			"ByteString",
			"[]byte",
			"ByteString",
			"string",
			"str",
			"String (ASCII-8BIT)",
		},
	}
}
