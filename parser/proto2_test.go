package parser_test

import (
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"github.com/pseudomuto/protoc-gen-doc/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	proto2File *parser.File
)

type Proto2ParserTest struct {
	suite.Suite
}

func TestProto2Parser(t *testing.T) {
	suite.Run(t, new(Proto2ParserTest))
}

func (assert *Proto2ParserTest) SetupSuite() {
	codeGenRequest, err := test.MakeCodeGeneratorRequest()
	assert.Nil(err)

	proto2File = parser.ParseCodeRequest(codeGenRequest).GetFile("Booking.proto")
}

func (assert *Proto2ParserTest) TestFileProperties() {
	assert.Equal("Booking.proto", proto2File.Name)
	assert.Equal("com.example", proto2File.Package)
	assert.Equal("Booking related messages.\n\nThis file is really just an example. The data model is completely\nfictional.", proto2File.Comment)
	assert.False(proto2File.IsProto3)
	assert.Equal(1, len(proto2File.Extensions))

	for _, msg := range []string{"Booking", "BookingStatus"} {
		assert.True(proto2File.HasMessage(msg))
	}

	for _, enum := range []string{"BookingType", "BookingStatus.StatusCode"} {
		assert.True(proto2File.HasEnum(enum))
	}
}

func (assert *Proto2ParserTest) TestFileExtensionProperties() {
	ext := proto2File.Extensions[0]
	assert.Equal(int32(100), ext.Number)
	assert.Equal("country", ext.Name)
	assert.Equal("com.example", ext.Package)
	assert.Equal("com.example.BookingStatus.country", ext.FullName())
	assert.Equal("The country the booking occurred in.", ext.Comment)
	assert.Equal("china", ext.DefaultValue)
	assert.Equal("com.example.BookingStatus", ext.ContainingType)
	assert.Equal("", ext.ScopeType)
	assert.False(ext.IsProto3)
}

func (assert *Proto2ParserTest) TestEnumProperties() {
	enum := proto2File.GetEnum("BookingType")
	assert.Equal("BookingType", enum.Name)
	assert.Equal("com.example", enum.Package)
	assert.Equal("com.example.BookingType", enum.FullName())
	assert.Equal("The type of booking.", enum.Comment)
	assert.False(enum.IsProto3)

	value := enum.Values[0]
	assert.Equal("IMMEDIATE", value.Name)
	assert.Equal(int32(100), value.Number)
	assert.Equal("com.example", value.Package)
	assert.Equal("Immediate booking.", value.Comment)
	assert.False(value.IsProto3)

	value = enum.Values[1]
	assert.Equal("FUTURE", value.Name)
	assert.Equal(int32(101), value.Number)
	assert.Equal("com.example", value.Package)
	assert.Equal("Future booking.", value.Comment)
	assert.False(value.IsProto3)
}

func (assert *Proto2ParserTest) TestNestedEnumProperties() {
	enum := proto2File.GetEnum("BookingStatus.StatusCode")
	assert.Equal("BookingStatus.StatusCode", enum.Name)
	assert.Equal("com.example", enum.Package)
	assert.Equal("com.example.BookingStatus.StatusCode", enum.FullName())
	assert.Equal("A flag for the status result.", enum.Comment)
	assert.False(enum.IsProto3)

	value := enum.Values[0]
	assert.Equal("OK", value.Name)
	assert.Equal(int32(200), value.Number)
	assert.Equal("com.example", value.Package)
	assert.Equal("OK result.", value.Comment)
	assert.False(value.IsProto3)

	value = enum.Values[1]
	assert.Equal("BAD_REQUEST", value.Name)
	assert.Equal(int32(400), value.Number)
	assert.Equal("com.example", value.Package)
	assert.Equal("BAD result.", value.Comment)
	assert.False(value.IsProto3)
}

func (assert *Proto2ParserTest) TestMessageProperties() {
	msg := proto2File.GetMessage("BookingStatus")
	assert.Equal("BookingStatus", msg.Name)
	assert.Equal("com.example", msg.Package)
	assert.Equal("com.example.BookingStatus", msg.FullName())
	assert.Equal("Represents the status of a vehicle booking.", msg.Comment)
	assert.False(msg.IsProto3)
	assert.Equal(3, len(msg.Fields))

	assert.field(msg.Fields[0], "id", "Unique booking status ID.", "int32", "required")
	assert.field(msg.Fields[2], "status_code", "The status of this status?", "BookingStatus.StatusCode", "optional")
}

func (assert *Proto2ParserTest) TestMessageExtensionProperties() {
	ext := proto2File.GetMessage("Booking").Extensions[0]
	assert.Equal(int32(101), ext.Number)
	assert.Equal("optional_field_1", ext.Name)
	assert.Equal("com.example", ext.Package)
	assert.Equal("com.example.BookingStatus.optional_field_1", ext.FullName())
	assert.Equal("An optional field to be used however you please.", ext.Comment)
	assert.Equal("", ext.DefaultValue)
	assert.Equal("com.example.BookingStatus", ext.ContainingType)
	assert.Equal("com.example.Booking", ext.ScopeType)
	assert.False(ext.IsProto3)
}

func (assert *Proto2ParserTest) TestServiceProperties() {
	service := proto2File.GetService("BookingService")
	assert.Equal("BookingService", service.Name)
	assert.Equal("com.example", service.Package)
	assert.Equal("com.example.BookingService", service.FullName())
	assert.Equal("Service for handling vehicle bookings.", service.Comment)
	assert.False(service.IsProto3)
	assert.Equal(1, len(service.Methods))

	method := service.Methods[0]
	assert.Equal("BookVehicle", method.Name)
	assert.Equal("Used to book a vehicle. Pass in a Booking and a BookingStatus will be returned.", method.Comment)
	assert.False(method.ClientStreaming)
	assert.False(method.ServerStreaming)
	assert.Equal("com.example.Booking", method.RequestType)
	assert.Equal("com.example.BookingStatus", method.ResponseType)
	assert.Equal("com.example", method.Package)
	assert.False(method.IsProto3)
}

func (assert *Proto2ParserTest) field(field *parser.Field, name, comment, typeName, label string) {
	assert.Equal(name, field.Name)
	assert.Equal(comment, field.Comment)
	assert.Equal(typeName, field.Type)
	assert.Equal(label, field.Label)
	assert.Equal("", field.DefaultValue)
}
