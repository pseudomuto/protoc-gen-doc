package protoc_gen_doc

import (
	"encoding/json"
	"fmt"
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"sort"
	"strings"
)

type Template struct {
	Files   []*File        `json:"files"`
	Scalars []*ScalarValue `json:"scalarValueTypes"`
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

	scalars := make([]*ScalarValue, 0)
	json.Unmarshal(data, &scalars)

	return scalars
}

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

type MessageField struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	LongType     string `json:"longType"`
	FullType     string `json:"fullType"`
	DefaultValue string `json:"defaultValue"`
}

type MessageExtension struct {
	FileExtension

	ScopeType     string `json:"scopeType"`
	ScopeLongType string `json:"scopeLongType"`
	ScopeFullType string `json:"scopeFullType"`
}

type Enum struct {
	Name        string       `json:"name"`
	LongName    string       `json:"longName"`
	FullName    string       `json:"fullName"`
	Description string       `json:"description"`
	Values      []*EnumValue `json:"values"`
}

type EnumValue struct {
	Name        string `json:"name"`
	Number      string `json:"number"`
	Description string `json:"description"`
}

type Service struct {
	Name        string           `json:"name"`
	LongName    string           `json:"longName"`
	FullName    string           `json:"fullName"`
	Description string           `json:"description"`
	Methods     []*ServiceMethod `json:"methods"`
}

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
