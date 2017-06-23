package parser

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type parsedObject struct {
	Name     string
	Comment  string
	Package  string
	IsProto3 bool
	path     string
}

func (po *parsedObject) getPath() string {
	return po.path
}

func (po *parsedObject) setComment(comment string) {
	po.Comment = comment
}

func (po *parsedObject) FullName() string {
	return fmt.Sprintf("%s.%s", po.Package, po.Name)
}

type File struct {
	parsedObject
	Enums      []*Enum
	Extensions []*Extension
	Messages   []*Message
	Services   []*Service
}

func (pf *File) getCommentContainers() []commentContainer {
	containers := []commentContainer{}
	containers = append(containers, pf)

	for _, ext := range pf.Extensions {
		containers = append(containers, ext)
	}

	for _, enum := range pf.Enums {
		containers = append(containers, enum)
		for _, value := range enum.Values {
			containers = append(containers, value)
		}
	}

	for _, msg := range pf.Messages {
		containers = append(containers, msg)
		for _, field := range msg.Fields {
			containers = append(containers, field)
		}

		for _, ext := range msg.Extensions {
			containers = append(containers, ext)
		}
	}

	for _, svc := range pf.Services {
		containers = append(containers, svc)
		for _, method := range svc.Methods {
			containers = append(containers, method)
		}
	}

	return containers
}

func (pf *File) HasEnum(name string) bool {
	return pf.GetEnum(name) != nil
}

func (pf *File) GetEnum(name string) *Enum {
	for _, enum := range pf.Enums {
		if enum.Name == name {
			return enum
		}
	}

	return nil
}

func (pf *File) HasMessage(name string) bool {
	return pf.GetMessage(name) != nil
}

func (pf *File) GetMessage(name string) *Message {
	for _, msg := range pf.Messages {
		if msg.Name == name {
			return msg
		}
	}

	return nil
}

func (pf *File) HasService(name string) bool {
	return pf.GetService(name) != nil
}

func (pf *File) GetService(name string) *Service {
	for _, service := range pf.Services {
		if service.Name == name {
			return service
		}
	}

	return nil
}

type Service struct {
	parsedObject
	Methods []*ServiceMethod
}

type ServiceMethod struct {
	parsedObject
	ClientStreaming bool
	ServerStreaming bool
	RequestType     string
	ResponseType    string
}

type Message struct {
	parsedObject
	Extensions []*Extension
	Fields     []*Field
	enums      []*descriptor.EnumDescriptorProto
}

type Field struct {
	parsedObject
	Type         string
	Label        string
	DefaultValue string
}

type Extension struct {
	Field
	Number         int32
	ContainingType string
	ScopeType      string
}

func (ext *Extension) FullName() string {
	return fmt.Sprintf("%s.%s", ext.ContainingType, ext.Name)
}

type Enum struct {
	parsedObject
	Values []*EnumValue
}

type EnumValue struct {
	parsedObject
	Number int32
}
