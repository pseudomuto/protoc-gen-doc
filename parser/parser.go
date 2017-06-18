package parser

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"path"
	"strconv"
	"strings"
)

const (
	// tag numbers in FileDescriptorProto
	packagePath = 2  // package
	messagePath = 4  // message_type
	enumPath    = 5  // enum_type
	servicePath = 6  // service
	syntaxPath  = 12 // syntax

	// tag numbers in DescriptorProto
	messageFieldPath   = 2 // field
	messageMessagePath = 3 // nested_type
	messageEnumPath    = 4 // enum_type

	// tag numbers in EnumDescriptorProto
	enumValuePath = 2 // value

	// tag numbers in ServiceDescriptorProto
	serviceMethodPath = 2 // method
)

type ParseResult struct {
	Files []*File
}

func (pr *ParseResult) GetFile(name string) *File {
	for _, f := range pr.Files {
		if f.Name == name {
			return f
		}
	}

	return nil
}

func ParseCodeRequest(req *plugin_go.CodeGeneratorRequest) *ParseResult {
	result := new(ParseResult)

	for _, file := range req.GetProtoFile() {
		result.Files = append(result.Files, parseFile(file))
	}

	return result
}

func parseFile(fd *descriptor.FileDescriptorProto) *File {
	pf := &File{
		parsedObject: parsedObject{
			Name:     path.Base(fd.GetName()),
			Package:  fd.GetPackage(),
			IsProto3: fd.GetSyntax() == "proto3",
			path:     strconv.Itoa(syntaxPath),
		},
		Messages: parseMessages(fd),
		Services: parseServices(fd),
	}

	pf.Enums = parseEnums(fd, pf.Messages)

	comments := newCommentExtractor(fd)
	comments.extractComments(pf.getCommentContainers()...)

	return pf
}

func parseServices(fd *descriptor.FileDescriptorProto) []*Service {
	descriptors := fd.GetService()
	services := make([]*Service, 0, len(descriptors))

	pkg, isProto3 := fd.GetPackage(), fd.GetSyntax() == "proto3"

	for idx, descriptor := range descriptors {
		path := fmt.Sprintf("%d,%d", servicePath, idx)

		services = append(services, &Service{
			parsedObject: parsedObject{
				Name:     descriptor.GetName(),
				Package:  pkg,
				IsProto3: isProto3,
				path:     path,
			},
			Methods: parseServiceMethods(descriptor, pkg, path, isProto3),
		})
	}

	return services
}

func parseServiceMethods(sd *descriptor.ServiceDescriptorProto, pkg, path string, isProto3 bool) []*ServiceMethod {
	descriptors := sd.GetMethod()
	methods := make([]*ServiceMethod, 0, len(descriptors))

	for idx, method := range descriptors {
		methods = append(methods, &ServiceMethod{
			parsedObject: parsedObject{
				Name:     method.GetName(),
				Package:  pkg,
				IsProto3: isProto3,
				path:     fmt.Sprintf("%s,%d,%d", path, serviceMethodPath, idx),
			},
			ClientStreaming: method.GetClientStreaming(),
			ServerStreaming: method.GetServerStreaming(),
			RequestType:     strings.TrimPrefix(method.GetInputType(), "."),
			ResponseType:    strings.TrimPrefix(method.GetOutputType(), "."),
		})
	}

	return methods
}

func parseMessages(fd *descriptor.FileDescriptorProto) []*Message {
	descriptors := fd.GetMessageType()
	msgs := make([]*Message, 0, len(descriptors)+10)

	for idx, msg := range descriptors {
		msgs = parseDescriptor(msgs, msg, nil, fd, idx)
	}

	return msgs
}

func parseDescriptor(sl []*Message, d *descriptor.DescriptorProto, p *Message, fd *descriptor.FileDescriptorProto, idx int) []*Message {
	sl = append(sl, &Message{
		parsedObject: parsedObject{
			Name:     d.GetName(),
			Package:  fd.GetPackage(),
			IsProto3: fd.GetSyntax() == "proto3",
			path:     fmt.Sprintf("%d,%d", messagePath, idx),
		},
		enums: d.GetEnumType(),
	})

	this := sl[len(sl)-1]
	if p != nil {
		this.Name = fmt.Sprintf("%s.%s", p.Name, this.Name)
		this.path = fmt.Sprintf("%s,%d,%d", p.path, messageMessagePath, idx)
	}

	parseFields(this, d.GetField())

	for i, nested := range d.GetNestedType() {
		sl = parseDescriptor(sl, nested, this, fd, i)
	}

	return sl
}

func parseFields(msg *Message, fields []*descriptor.FieldDescriptorProto) {
	typeName := func(name string) string {
		if strings.HasPrefix(name, ".") {
			return strings.TrimPrefix(name, fmt.Sprintf(".%s.", msg.Package))
		}

		return strings.ToLower(strings.TrimPrefix(name, "TYPE_"))
	}

	for idx, field := range fields {
		msg.Fields = append(msg.Fields, &Field{
			parsedObject: parsedObject{
				Name:     field.GetName(),
				Package:  msg.Package,
				IsProto3: msg.IsProto3,
				path:     fmt.Sprintf("%s,%d,%d", msg.path, messageFieldPath, idx),
			},
			Type: typeName(field.GetTypeName()),
		})

		f := msg.Fields[len(msg.Fields)-1]

		if f.Type == "" {
			f.Type = typeName(field.GetType().String())
		}

		if f.IsProto3 && field.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			f.Label = "repeated"
		}
	}
}

func parseEnums(fd *descriptor.FileDescriptorProto, msgs []*Message) []*Enum {
	enums := make([]*Enum, 0, len(fd.GetEnumType())+10)

	for idx, enum := range fd.GetEnumType() {
		enums = append(enums, parseEnumDescriptor(enum, nil, fd, idx))
	}

	for _, msg := range msgs {
		for idx, enum := range msg.enums {
			enums = append(enums, parseEnumDescriptor(enum, msg, fd, idx))
		}
	}

	return enums
}

func parseEnumDescriptor(e *descriptor.EnumDescriptorProto, p *Message, fd *descriptor.FileDescriptorProto, idx int) *Enum {
	enum := &Enum{
		parsedObject: parsedObject{
			Name:     e.GetName(),
			Package:  fd.GetPackage(),
			IsProto3: fd.GetSyntax() == "proto3",
			path:     fmt.Sprintf("%d,%d", enumPath, idx),
		},
	}

	if p != nil {
		enum.Name = fmt.Sprintf("%s.%s", p.Name, enum.Name)
		enum.path = fmt.Sprintf("%s,%d,%d", p.path, messageEnumPath, idx)
	}

	enum.Values = parseEnumValues(e.GetValue(), enum)
	return enum
}

func parseEnumValues(vd []*descriptor.EnumValueDescriptorProto, enum *Enum) []*EnumValue {
	values := make([]*EnumValue, 0, len(vd))

	for idx, val := range vd {
		values = append(values, &EnumValue{
			parsedObject: parsedObject{
				Name:     val.GetName(),
				Package:  enum.Package,
				IsProto3: enum.IsProto3,
				path:     fmt.Sprintf("%s,%d,%d", enum.path, enumValuePath, idx),
			},
			Number: int32(idx),
		})
	}

	return values
}
