package gendoc

import (
	"encoding/json"
	"fmt"
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"sort"
	"strings"
)

// Template is a type for encapsulating all the parsed files, messages, fields, enums, services, extensions, etc. into
// an object that will be supplied to a go template.
type Template struct {
	// The files that were parsed
	Files []*File `json:"files"`
	// Details about the scalar values and their respective types in supported languages.
	Scalars []*ScalarValue `json:"scalarValueTypes"`
}

// NewTemplate creates a Template object from the ParseResult.
func NewTemplate(pr *parser.ParseResult) *Template {
	files := make([]*File, 0, len(pr.Files))

	for _, f := range pr.Files {
		file := &File{
			Name:          f.Name,
			Description:   description(f.Comment),
			Package:       f.Package,
			HasEnums:      len(f.Enums) > 0,
			HasExtensions: len(f.Extensions) > 0,
			HasMessages:   len(f.Messages) > 0,
			HasServices:   len(f.Services) > 0,
			Enums:         make(orderedEnums, 0, len(f.Enums)),
			Extensions:    make(orderedExtensions, 0, len(f.Extensions)),
			Messages:      make(orderedMessages, 0, len(f.Messages)),
			Services:      make(orderedServices, 0, len(f.Services)),
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

		sort.Sort(file.Enums)
		sort.Sort(file.Extensions)
		sort.Sort(file.Messages)
		sort.Sort(file.Services)

		files = append(files, file)
	}

	return &Template{Files: files, Scalars: makeScalars()}
}

func makeScalars() []*ScalarValue {
	data, _ := fetchResource("scalars.json")

	var scalars []*ScalarValue
	json.Unmarshal(data, &scalars)

	return scalars
}

// File wraps all the relevant parsed info about a proto file. File objects guarantee that their top-level enums,
// extensions, messages, and services are sorted alphabetically based on their "long name". Other values (enum values,
// fields, service methods) will be in the order that they're defined within their respective proto files.
//
// In the case of proto3 files, HasExtensions will always be false, and Extensions will be empty.
type File struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Package     string `json:"package"`

	HasEnums      bool `json:"hasEnums"`
	HasExtensions bool `json:"hasExtensions"`
	HasMessages   bool `json:"hasMessages"`
	HasServices   bool `json:"hasServices"`

	Enums      orderedEnums      `json:"enums"`
	Extensions orderedExtensions `json:"extensions"`
	Messages   orderedMessages   `json:"messages"`
	Services   orderedServices   `json:"services"`
}

// FileExtension contains details about top-level extensions within a proto(2) file.
type FileExtension struct {
	Name               string `json:"name"`
	LongName           string `json:"longName"`
	FullName           string `json:"fullName"`
	Description        string `json:"description"`
	Label              string `json:"label"`
	Type               string `json:"type"`
	LongType           string `json:"longType"`
	FullType           string `json:"fullType"`
	Number             int    `json:"number"`
	DefaultValue       string `json:"defaultValue"`
	ContainingType     string `json:"containingType"`
	ContainingLongType string `json:"containingLongType"`
	ContainingFullType string `json:"containingFullType"`
}

// Message contains details about a protobuf message.
//
// In the case of proto3 files, HasExtensions will always be false, and Extensions will be empty.
type Message struct {
	Name        string `json:"name"`
	LongName    string `json:"longName"`
	FullName    string `json:"fullName"`
	Description string `json:"description"`

	HasExtensions bool `json:"hasExtensions"`
	HasFields     bool `json:"hasFields"`

	Extensions []*MessageExtension `json:"extensions"`
	Fields     []*MessageField     `json:"fields"`
}

// MessageField contains details about an individual field within a message.
//
// In the case of proto3 files, DefaultValue will always be empty. Similarly, label will be empty unless the field is
// repeated (in which case it'll be "repeated").
type MessageField struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	LongType     string `json:"longType"`
	FullType     string `json:"fullType"`
	DefaultValue string `json:"defaultValue"`
}

// MessageExtension contains details about message-scoped extensions in proto(2) files.
type MessageExtension struct {
	FileExtension

	ScopeType     string `json:"scopeType"`
	ScopeLongType string `json:"scopeLongType"`
	ScopeFullType string `json:"scopeFullType"`
}

// Enum contains details about enumerations. These can be either top level enums, or nested (defined within a message).
type Enum struct {
	Name        string       `json:"name"`
	LongName    string       `json:"longName"`
	FullName    string       `json:"fullName"`
	Description string       `json:"description"`
	Values      []*EnumValue `json:"values"`
}

// EnumValue contains details about an individual value within an enumeration.
type EnumValue struct {
	Name        string `json:"name"`
	Number      string `json:"number"`
	Description string `json:"description"`
}

// Service contains details about a service definition within a proto file.
type Service struct {
	Name        string           `json:"name"`
	LongName    string           `json:"longName"`
	FullName    string           `json:"fullName"`
	Description string           `json:"description"`
	Methods     []*ServiceMethod `json:"methods"`
}

// ServiceMethod contains details about an individual method within a service.
type ServiceMethod struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	RequestType      string `json:"requestType"`
	RequestLongType  string `json:"requestLongType"`
	RequestFullType  string `json:"requestFullType"`
	ResponseType     string `json:"responseType"`
	ResponseLongType string `json:"responseLongType"`
	ResponseFullType string `json:"responseFullType"`
}

// ScalarValue contains information about scalar value types in protobuf. The common use case for this type is to know
// which language specific type maps to the protobuf type.
//
// For example, the protobuf type `int64` maps to `long` in C#, and `Bignum` in Ruby. For the full list, take a look at
// https://developers.google.com/protocol-buffers/docs/proto3#scalar
type ScalarValue struct {
	ProtoType  string `json:"protoType"`
	Notes      string `json:"notes"`
	CppType    string `json:"cppType"`
	CSharp     string `json:"csType"`
	GoType     string `json:"goType"`
	JavaType   string `json:"javaType"`
	PhpType    string `json:"phpType"`
	PythonType string `json:"pythonType"`
	RubyType   string `json:"rubyType"`
}

func parseEnum(pe *parser.Enum) *Enum {
	enum := &Enum{
		Name:        baseName(pe.Name),
		LongName:    strings.TrimPrefix(pe.FullName(), pe.Package+"."),
		FullName:    pe.FullName(),
		Description: description(pe.Comment),
	}

	for _, val := range pe.Values {
		enum.Values = append(enum.Values, &EnumValue{
			Name:        val.Name,
			Number:      fmt.Sprint(val.Number),
			Description: description(val.Comment),
		})
	}

	return enum
}

func parseFileExtension(pe *parser.Extension) *FileExtension {
	return &FileExtension{
		Name:               baseName(pe.Name),
		LongName:           strings.TrimPrefix(pe.FullName(), pe.Package+"."),
		FullName:           pe.FullName(),
		Description:        description(pe.Comment),
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
		Description:   description(pm.Comment),
		HasExtensions: len(pm.Extensions) > 0,
		HasFields:     len(pm.Fields) > 0,
		Extensions:    make([]*MessageExtension, 0, len(pm.Extensions)),
		Fields:        make([]*MessageField, 0, len(pm.Fields)),
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
		Description:  description(pf.Comment),
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
		Description: description(ps.Comment),
	}

	for _, sm := range ps.Methods {
		service.Methods = append(service.Methods, parseServiceMethod(sm))
	}

	return service
}

func parseServiceMethod(pm *parser.ServiceMethod) *ServiceMethod {
	return &ServiceMethod{
		Name:             pm.Name,
		Description:      description(pm.Comment),
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

func description(comment string) string {
	if strings.HasPrefix(comment, "@exclude") {
		return ""
	}

	return comment
}

type orderedEnums []*Enum

func (oe orderedEnums) Len() int           { return len(oe) }
func (oe orderedEnums) Swap(i, j int)      { oe[i], oe[j] = oe[j], oe[i] }
func (oe orderedEnums) Less(i, j int) bool { return oe[i].LongName < oe[j].LongName }

type orderedExtensions []*FileExtension

func (oe orderedExtensions) Len() int           { return len(oe) }
func (oe orderedExtensions) Swap(i, j int)      { oe[i], oe[j] = oe[j], oe[i] }
func (oe orderedExtensions) Less(i, j int) bool { return oe[i].LongName < oe[j].LongName }

type orderedMessages []*Message

func (om orderedMessages) Len() int           { return len(om) }
func (om orderedMessages) Swap(i, j int)      { om[i], om[j] = om[j], om[i] }
func (om orderedMessages) Less(i, j int) bool { return om[i].LongName < om[j].LongName }

type orderedServices []*Service

func (os orderedServices) Len() int           { return len(os) }
func (os orderedServices) Swap(i, j int)      { os[i], os[j] = os[j], os[i] }
func (os orderedServices) Less(i, j int) bool { return os[i].LongName < os[j].LongName }
