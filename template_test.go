package gendoc_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	. "github.com/pseudomuto/protoc-gen-doc"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
	"github.com/pseudomuto/protokit"
	"github.com/pseudomuto/protokit/utils"
	"github.com/stretchr/testify/suite"
)

var (
	template    *Template
	bookingFile *File
	vehicleFile *File
)

type TemplateTest struct {
	suite.Suite
}

func TestTemplate(t *testing.T) {
	suite.Run(t, new(TemplateTest))
}

func (assert *TemplateTest) SetupSuite() {
	registerTestExtensions()

	set, err := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	assert.NoError(err)

	req := utils.CreateGenRequest(set, "Booking.proto", "Vehicle.proto")
	result := protokit.ParseCodeGenRequest(req)
	template = NewTemplate(result)
	bookingFile = template.Files[0]
	vehicleFile = template.Files[1]
}

func identity(payload interface{}) interface{} { return payload }

var E_ExtendFile = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         20000,
	Name:          "com.pseudomuto.protokit.v1.extend_file",
	Tag:           "varint,20000,opt,name=extend_file,json=extendFile",
	Filename:      "extend.proto",
}

var E_ExtendService = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         20000,
	Name:          "com.pseudomuto.protokit.v1.extend_service",
	Tag:           "varint,20000,opt,name=extend_service,json=extendService",
	Filename:      "extend.proto",
}

var E_ExtendMethod = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         20000,
	Name:          "com.pseudomuto.protokit.v1.extend_method",
	Tag:           "varint,20000,opt,name=extend_method,json=extendMethod",
	Filename:      "extend.proto",
}

var E_ExtendEnum = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         20000,
	Name:          "com.pseudomuto.protokit.v1.extend_enum",
	Tag:           "varint,20000,opt,name=extend_enum,json=extendEnum",
	Filename:      "extend.proto",
}

var E_ExtendEnumValue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         20000,
	Name:          "com.pseudomuto.protokit.v1.extend_enum_value",
	Tag:           "varint,20000,opt,name=extend_enum_value,json=extendEnumValue",
	Filename:      "extend.proto",
}

var E_ExtendMessage = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         20000,
	Name:          "com.pseudomuto.protokit.v1.extend_message",
	Tag:           "varint,20000,opt,name=extend_message,json=extendMessage",
	Filename:      "extend.proto",
}

var E_ExtendField = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         20000,
	Name:          "com.pseudomuto.protokit.v1.extend_field",
	Tag:           "varint,20000,opt,name=extend_field,json=extendField",
	Filename:      "extend.proto",
}

func registerTestExtensions() {
	proto.RegisterExtension(E_ExtendFile)
	extensions.SetTransformer(E_ExtendFile.Name, identity)
	proto.RegisterExtension(E_ExtendService)
	extensions.SetTransformer(E_ExtendService.Name, identity)
	proto.RegisterExtension(E_ExtendMethod)
	extensions.SetTransformer(E_ExtendMethod.Name, identity)
	proto.RegisterExtension(E_ExtendEnum)
	extensions.SetTransformer(E_ExtendEnum.Name, identity)
	proto.RegisterExtension(E_ExtendEnumValue)
	extensions.SetTransformer(E_ExtendEnumValue.Name, identity)
	proto.RegisterExtension(E_ExtendMessage)
	extensions.SetTransformer(E_ExtendMessage.Name, identity)
	proto.RegisterExtension(E_ExtendField)
	extensions.SetTransformer(E_ExtendField.Name, identity)
}

func (assert *TemplateTest) TestTemplateProperties() {
	assert.Equal(2, len(template.Files))
}

func (assert *TemplateTest) TestFileProperties() {
	assert.Equal("Booking.proto", bookingFile.Name)
	assert.Equal("Booking related messages.\n\nThis file is really just an example. The data model is completely\nfictional.", bookingFile.Description)
	assert.Equal("com.example", bookingFile.Package)
	assert.True(bookingFile.HasEnums)
	assert.True(bookingFile.HasExtensions)
	assert.True(bookingFile.HasMessages)
	assert.True(bookingFile.HasServices)
	if assert.NotNil(bookingFile.Options) {
		assert.NotEmpty(bookingFile.Options)
	}
	if assert.NotNil(bookingFile.Option(E_ExtendFile.Name)) {
		assert.True(*bookingFile.Option(E_ExtendFile.Name).(*bool))
	}
}

func (assert *TemplateTest) TestFileEnumProperties() {
	enum := findEnum("BookingStatus.StatusCode", bookingFile)
	assert.Equal("StatusCode", enum.Name)
	assert.Equal("BookingStatus.StatusCode", enum.LongName)
	assert.Equal("com.example.BookingStatus.StatusCode", enum.FullName)
	assert.Equal("A flag for the status result.", enum.Description)
	assert.Equal(2, len(enum.Values))

	expectedValues := []*EnumValue{
		{Name: "OK", Number: "200", Description: "OK result."},
		{Name: "BAD_REQUEST", Number: "400", Description: "BAD result."},
	}

	for idx, value := range enum.Values {
		assert.Equal(expectedValues[idx], value)
	}

	enum = findEnum("BookingType", bookingFile)
	if assert.NotNil(enum.Options) {
		assert.NotEmpty(enum.Options)
	}
	if assert.NotNil(enum.Option(E_ExtendEnum.Name)) {
		assert.True(*enum.Option(E_ExtendEnum.Name).(*bool))
	}
	assert.Contains(enum.ValueOptions(), E_ExtendEnumValue.Name)
	assert.NotEmpty(enum.ValuesWithOption(E_ExtendEnumValue.Name))

	for _, value := range enum.Values {
		if value.Name == "FUTURE" {
			if assert.NotNil(value.Options) {
				assert.NotEmpty(value.Options)
			}
			if assert.NotNil(value.Option(E_ExtendEnumValue.Name)) {
				assert.True(*value.Option(E_ExtendEnumValue.Name).(*bool))
			}
		}
	}
}

func (assert *TemplateTest) TestFileExtensionProperties() {
	ext := findExtension("BookingStatus.country", bookingFile)
	assert.Equal("country", ext.Name)
	assert.Equal("BookingStatus.country", ext.LongName)
	assert.Equal("com.example.BookingStatus.country", ext.FullName)
	assert.Equal("The country the booking occurred in.", ext.Description)
	assert.Equal("optional", ext.Label)
	assert.Equal("string", ext.Type)
	assert.Equal("string", ext.LongType)
	assert.Equal("string", ext.FullType)
	assert.Equal(100, ext.Number)
	assert.Equal("china", ext.DefaultValue)
	assert.Equal("BookingStatus", ext.ContainingType)
	assert.Equal("BookingStatus", ext.ContainingLongType)
	assert.Equal("com.example.BookingStatus", ext.ContainingFullType)
}

func (assert *TemplateTest) TestMessageProperties() {
	msg := findMessage("Vehicle", vehicleFile)
	assert.Equal("Vehicle", msg.Name)
	assert.Equal("Vehicle", msg.LongName)
	assert.Equal("com.example.Vehicle", msg.FullName)
	assert.Equal("Represents a vehicle that can be hired.", msg.Description)
	assert.False(msg.HasExtensions)
	assert.True(msg.HasFields)
	if assert.NotNil(msg.Options) {
		assert.NotEmpty(msg.Options)
	}
	if assert.NotNil(msg.Option(E_ExtendMessage.Name)) {
		assert.True(*msg.Option(E_ExtendMessage.Name).(*bool))
	}
	assert.Contains(msg.FieldOptions(), E_ExtendField.Name)
	assert.NotEmpty(msg.FieldsWithOption(E_ExtendField.Name))
}

func (assert *TemplateTest) TestNestedMessageProperties() {
	msg := findMessage("Vehicle.Category", vehicleFile)
	assert.Equal("Category", msg.Name)
	assert.Equal("Vehicle.Category", msg.LongName)
	assert.Equal("com.example.Vehicle.Category", msg.FullName)
	assert.Equal("Represents a vehicle category. E.g. \"Sedan\" or \"Truck\".", msg.Description)
	assert.False(msg.HasExtensions)
	assert.True(msg.HasFields)
}

func (assert *TemplateTest) TestMultiplyNestedMessages() {
	assert.NotNil(findEnum("Vehicle.Engine.FuelType", vehicleFile))
	assert.NotNil(findMessage("Vehicle.Engine.Stats", vehicleFile))
}

func (assert *TemplateTest) TestMessageExtensionProperties() {
	msg := findMessage("Booking", bookingFile)
	assert.Equal(1, len(msg.Extensions))

	ext := msg.Extensions[0]
	assert.Equal("optional_field_1", ext.Name)
	assert.Equal("BookingStatus.optional_field_1", ext.LongName)
	assert.Equal("com.example.BookingStatus.optional_field_1", ext.FullName)
	assert.Equal("An optional field to be used however you please.", ext.Description)
	assert.Equal("optional", ext.Label)
	assert.Equal("string", ext.Type)
	assert.Equal("string", ext.LongType)
	assert.Equal("string", ext.FullType)
	assert.Equal(101, ext.Number)
	assert.Equal("", ext.DefaultValue)
	assert.Equal("BookingStatus", ext.ContainingType)
	assert.Equal("BookingStatus", ext.ContainingLongType)
	assert.Equal("com.example.BookingStatus", ext.ContainingFullType)
	assert.Equal("Booking", ext.ScopeType)
	assert.Equal("Booking", ext.ScopeLongType)
	assert.Equal("com.example.Booking", ext.ScopeFullType)
}

func (assert *TemplateTest) TestFieldProperties() {
	msg := findMessage("BookingStatus", bookingFile)

	field := findField("id", msg)
	assert.Equal("id", field.Name)
	assert.Equal("Unique booking status ID.", field.Description)
	assert.Equal("required", field.Label)
	assert.Equal("int32", field.Type)
	assert.Equal("int32", field.LongType)
	assert.Equal("int32", field.FullType)
	assert.Equal("", field.DefaultValue)
	if assert.NotNil(field.Options) {
		assert.NotEmpty(field.Options)
	}
	if assert.NotNil(field.Option(E_ExtendField.Name)) {
		assert.True(*field.Option(E_ExtendField.Name).(*bool))
	}

	field = findField("status_code", msg)
	assert.Equal("status_code", field.Name)
	assert.Equal("The status of this status?", field.Description)
	assert.Equal("optional", field.Label)
	assert.Equal("StatusCode", field.Type)
	assert.Equal("BookingStatus.StatusCode", field.LongType)
	assert.Equal("com.example.BookingStatus.StatusCode", field.FullType)
	assert.Equal("", field.DefaultValue)

	field = findField("category", findMessage("Vehicle", vehicleFile))
	assert.Equal("category", field.Name)
	assert.Equal("Vehicle category.", field.Description)
	assert.Equal("", field.Label) // proto3, neither required, nor optional are valid
	assert.Equal("Category", field.Type)
	assert.Equal("Vehicle.Category", field.LongType)
	assert.Equal("com.example.Vehicle.Category", field.FullType)
	assert.Equal("", field.DefaultValue)

	field = findField("properties", findMessage("Vehicle", vehicleFile))
	assert.Equal("properties", field.Name)
	assert.Equal("repeated", field.Label)
	assert.Equal("PropertiesEntry", field.Type)
	assert.Equal("Vehicle.PropertiesEntry", field.LongType)
	assert.Equal("com.example.Vehicle.PropertiesEntry", field.FullType)
	assert.Equal("", field.DefaultValue)
	assert.True(field.IsMap)

	field = findField("rates", findMessage("Vehicle", vehicleFile))
	assert.Equal("rates", field.Name)
	assert.Equal("repeated", field.Label)
	assert.Equal("sint32", field.Type)
	assert.Equal("sint32", field.LongType)
	assert.Equal("sint32", field.FullType)
	assert.False(field.IsMap)
}

func (assert *TemplateTest) TestServiceProperties() {
	service := findService("VehicleService", vehicleFile)
	assert.Equal("VehicleService", service.Name)
	assert.Equal("VehicleService", service.LongName)
	assert.Equal("com.example.VehicleService", service.FullName)
	assert.Equal("The vehicle service.\n\nManages vehicles and such...", service.Description)
	assert.Equal(3, len(service.Methods))
	if assert.NotNil(service.Options) {
		assert.NotEmpty(service.Options)
	}
	if assert.NotNil(service.Option(E_ExtendService.Name)) {
		assert.True(*service.Option(E_ExtendService.Name).(*bool))
	}
	assert.Contains(service.MethodOptions(), E_ExtendMethod.Name)
	assert.NotEmpty(service.MethodsWithOption(E_ExtendMethod.Name))
}

func (assert *TemplateTest) TestServiceMethodProperties() {
	service := findService("VehicleService", vehicleFile)

	method := findServiceMethod("AddModels", service)
	assert.Equal("AddModels", method.Name)
	assert.Equal("creates models", method.Description)
	assert.Equal("Model", method.RequestType)
	assert.Equal("Model", method.RequestLongType)
	assert.Equal("com.example.Model", method.RequestFullType)
	assert.Equal(true, method.RequestStreaming)
	assert.Equal("Model", method.ResponseType)
	assert.Equal("Model", method.ResponseLongType)
	assert.Equal("com.example.Model", method.ResponseFullType)
	assert.Equal(true, method.ResponseStreaming)

	method = findServiceMethod("GetVehicle", service)
	assert.Equal("GetVehicle", method.Name)
	assert.Equal("Looks up a vehicle by id.", method.Description)
	assert.Equal("FindVehicleById", method.RequestType)
	assert.Equal("FindVehicleById", method.RequestLongType)
	assert.Equal("com.example.FindVehicleById", method.RequestFullType)
	assert.Equal(false, method.RequestStreaming)
	assert.Equal("Vehicle", method.ResponseType)
	assert.Equal("Vehicle", method.ResponseLongType)
	assert.Equal("com.example.Vehicle", method.ResponseFullType)
	assert.Equal(false, method.ResponseStreaming)
	if assert.NotNil(method.Options) {
		assert.NotEmpty(method.Options)
	}
	if assert.NotNil(method.Option(E_ExtendMethod.Name)) {
		assert.True(*method.Option(E_ExtendMethod.Name).(*bool))
	}
}

func (assert *TemplateTest) TestExcludedComments() {
	message := findMessage("ExcludedMessage", vehicleFile)
	assert.Empty(message.Description)
	assert.Empty(findField("name", message).Description)
	assert.Empty(findField("value", message).Description)

	// just checking that it doesn't exclude everything
	assert.Equal("the id of this message.", findField("id", message).Description)
}

func findService(name string, f *File) *Service {
	for _, s := range f.Services {
		if s.Name == name {
			return s
		}
	}

	return nil
}

func findServiceMethod(name string, s *Service) *ServiceMethod {
	for _, m := range s.Methods {
		if m.Name == name {
			return m
		}
	}

	return nil
}

func findEnum(name string, f *File) *Enum {
	for _, enum := range f.Enums {
		if enum.LongName == name {
			return enum
		}
	}

	return nil
}

func findExtension(name string, f *File) *FileExtension {
	for _, ext := range f.Extensions {
		if ext.LongName == name {
			return ext
		}
	}

	return nil
}

func findMessage(name string, f *File) *Message {
	for _, m := range f.Messages {
		if m.LongName == name {
			return m
		}
	}

	return nil
}

func findField(name string, m *Message) *MessageField {
	for _, f := range m.Fields {
		if f.Name == name {
			return f
		}
	}

	return nil
}
