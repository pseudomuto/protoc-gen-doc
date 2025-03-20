package extensions

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/pseudomuto/protokit"
)

func RegisterAllExtensions(files []*protokit.FileDescriptor) {
	for _, file := range files {
		for _, ext := range file.GetExtensions() {
			registerExtension(ext)
		}

		for _, msg := range file.GetMessages() {
			registerMessageExtensions(msg)
		}
	}
}

func registerMessageExtensions(msg *protokit.Descriptor) {
	for _, ext := range msg.GetExtensions() {
		registerExtension(ext)
	}
	for _, nestedMsg := range msg.GetMessages() {
		registerMessageExtensions(nestedMsg)
	}
}
func registerExtension(ext *protokit.ExtensionDescriptor) {
	extendedType := determineExtendedType(ext.GetExtendee())
	extType := determineExtensionType(ext)
	if extType == nil {
		return
	}

	correctFullName := fmt.Sprintf("%s.%s", ext.GetFile().GetPackage(), ext.GetName())

	proto.RegisterExtension(&proto.ExtensionDesc{
		ExtendedType:  extendedType,
		ExtensionType: extType,
		Field:         int32(ext.GetNumber()),
		Name:          correctFullName,
		Tag:           generateTag(ext),
	})
}






// Helper functions:

func determineExtendedType(typeName string) proto.Message {
	switch typeName {
	case ".google.protobuf.MethodOptions":
		return (*descriptor.MethodOptions)(nil)
	case ".google.protobuf.FieldOptions":
		return (*descriptor.FieldOptions)(nil)
	case ".google.protobuf.MessageOptions":
		return (*descriptor.MessageOptions)(nil)
	case ".google.protobuf.FileOptions":
		return (*descriptor.FileOptions)(nil)
	case ".google.protobuf.EnumOptions":
		return (*descriptor.EnumOptions)(nil)
	case ".google.protobuf.EnumValueOptions":
		return (*descriptor.EnumValueOptions)(nil)
	case ".google.protobuf.ServiceOptions":
		return (*descriptor.ServiceOptions)(nil)
	default:
		panic(fmt.Sprintf("Unsupported extendee type: %s", typeName))
	}
}


func determineExtensionType(ext *protokit.ExtensionDescriptor) interface{} {
	isRepeated := ext.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED

	switch ext.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		if isRepeated {
			return ([]string)(nil)
		}
		return (*string)(nil)
	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		if isRepeated {
			return ([]bool)(nil)
		}
		return (*bool)(nil)
	case descriptor.FieldDescriptorProto_TYPE_INT32, descriptor.FieldDescriptorProto_TYPE_SINT32:
		if isRepeated {
			return ([]int32)(nil)
		}
		return (*int32)(nil)
	case descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_SINT64:
		if isRepeated {
			return ([]int64)(nil)
		}
		return (*int64)(nil)
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		if isRepeated {
			return ([]float32)(nil)
		}
		return (*float32)(nil)
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		if isRepeated {
			return ([]float64)(nil)
		}
		return (*float64)(nil)
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		if isRepeated {
			return ([]int32)(nil) // enums represented as repeated []int32
		}
		return (*int32)(nil) // single enum represented as *int32
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		return nil
	default:
		panic(fmt.Sprintf("Unsupported extension field type: %s", ext.GetType().String()))
	}
}



func generateTag(ext *protokit.ExtensionDescriptor) string {
	var wireType string
	switch ext.GetType() {
	case descriptor.FieldDescriptorProto_TYPE_STRING:
		wireType = "bytes"
	case descriptor.FieldDescriptorProto_TYPE_BOOL,
		descriptor.FieldDescriptorProto_TYPE_INT32,
		descriptor.FieldDescriptorProto_TYPE_INT64,
		descriptor.FieldDescriptorProto_TYPE_SINT32,
		descriptor.FieldDescriptorProto_TYPE_SINT64,
		descriptor.FieldDescriptorProto_TYPE_ENUM:
		wireType = "varint"
	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		wireType = "fixed32"
	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		wireType = "fixed64"
	default:
		panic(fmt.Sprintf("Unsupported tag type for extension: %s", ext.GetType().String()))
	}

	label := "opt"
	if ext.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		label = "rep"
	}

	fieldName := ext.GetName()
	return fmt.Sprintf("%s,%d,%s,name=%s", wireType, ext.GetNumber(), label, fieldName)
}
