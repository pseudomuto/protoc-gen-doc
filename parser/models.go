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

// File represents a parsed file object. All information about a proto file will be encapsulated in a File object.
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

// HasEnum indicated whether or not a file-level enum exists.
func (pf *File) HasEnum(name string) bool {
	return pf.GetEnum(name) != nil
}

// GetEnum finds an enum object by name and returns it. It will return `nil` if not found.
func (pf *File) GetEnum(name string) *Enum {
	for _, enum := range pf.Enums {
		if enum.Name == name {
			return enum
		}
	}

	return nil
}

// HasMessage indicated whether or not this file contains the named message.
func (pf *File) HasMessage(name string) bool {
	return pf.GetMessage(name) != nil
}

// GetMessage finds a message by name and returns it. It will return `nil` if not found.
func (pf *File) GetMessage(name string) *Message {
	for _, msg := range pf.Messages {
		if msg.Name == name {
			return msg
		}
	}

	return nil
}

// HasService indicated whether or not this file contains the named service.
func (pf *File) HasService(name string) bool {
	return pf.GetService(name) != nil
}

// GetService finds a service by name and returns it. It will return `nil` if not found.
func (pf *File) GetService(name string) *Service {
	for _, service := range pf.Services {
		if service.Name == name {
			return service
		}
	}

	return nil
}

// A Service object to encasulate service details.
type Service struct {
	parsedObject
	Methods []*ServiceMethod
}

// A ServiceMethod object to encapsulate service method details.
type ServiceMethod struct {
	parsedObject
	ClientStreaming bool
	ServerStreaming bool
	RequestType     string
	ResponseType    string
}

// A Message object to encapsulate message details.
type Message struct {
	parsedObject
	Extensions []*Extension
	Fields     []*Field
	enums      []*descriptor.EnumDescriptorProto
}

// A Field object to encapsulate message field details.
type Field struct {
	parsedObject
	Type         string
	Label        string
	DefaultValue string
	IsMap        bool
}

// An Extension object to encapsulate extension details.
type Extension struct {
	Field
	Label          string
	Number         int32
	ContainingType string
	ScopeType      string
}

// FullName returns the full name of this extension including the containing type
func (ext *Extension) FullName() string {
	return fmt.Sprintf("%s.%s", ext.ContainingType, ext.Name)
}

// An Enum object to encapsulate enum details.
type Enum struct {
	parsedObject
	Values []*EnumValue
}

// An EnumValue object to encapsulate enum value details.
type EnumValue struct {
	parsedObject
	Number int32
}
