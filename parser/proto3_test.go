package parser_test

import (
	"testing"

	"github.com/pseudomuto/protoc-gen-doc/parser"
	"github.com/pseudomuto/protoc-gen-doc/test"
	"github.com/stretchr/testify/suite"
)

var (
	proto3File *parser.File
)

type Proto3ParserTest struct {
	suite.Suite
}

func TestProto3Parser(t *testing.T) {
	suite.Run(t, new(Proto3ParserTest))
}

func (assert *Proto3ParserTest) SetupSuite() {
	codeGenRequest, err := test.MakeCodeGeneratorRequest()
	assert.Nil(err)

	proto3File = parser.ParseCodeRequest(codeGenRequest, nil).GetFile("Vehicle.proto")
}

func (assert *Proto3ParserTest) TestFileProperties() {
	assert.Equal("Vehicle.proto", proto3File.Name)
	assert.Equal("com.example", proto3File.Package)
	assert.Equal("Messages describing manufacturers / vehicles.", proto3File.Comment)
	assert.True(proto3File.IsProto3)
	assert.Equal(0, len(proto3File.Extensions))

	for _, msg := range []string{"EmptyMessage", "Manufacturer", "Model", "Vehicle", "Vehicle.Category"} {
		assert.True(proto3File.HasMessage(msg))
	}

	for _, enum := range []string{"Manufacturer.Category", "Type"} {
		assert.True(proto3File.HasEnum(enum))
	}
}

func (assert *Proto3ParserTest) TestEnumProperties() {
	enum := proto3File.GetEnum("Type")
	assert.Equal("Type", enum.Name)
	assert.Equal("com.example", enum.Package)
	assert.Equal("com.example.Type", enum.FullName())
	assert.Equal("The type of model.", enum.Comment)
	assert.True(enum.IsProto3)

	value := enum.Values[0]
	assert.Equal("COUPE", value.Name)
	assert.Equal(int32(0), value.Number)
	assert.Equal("com.example", value.Package)
	assert.Equal("The type is coupe.", value.Comment)
	assert.True(value.IsProto3)

	value = enum.Values[1]
	assert.Equal("SEDAN", value.Name)
	assert.Equal(int32(1), value.Number)
	assert.Equal("com.example", value.Package)
	assert.Equal("The type is sedan.", value.Comment)
	assert.True(value.IsProto3)
}

func (assert *Proto3ParserTest) TestNestedEnumProperties() {
	enum := proto3File.GetEnum("Manufacturer.Category")
	assert.Equal("Manufacturer.Category", enum.Name)
	assert.Equal("com.example", enum.Package)
	assert.Equal("com.example.Manufacturer.Category", enum.FullName())
	assert.Equal("Manufacturer category. A manufacturer may be either inhouse or external.", enum.Comment)
	assert.True(enum.IsProto3)

	value := enum.Values[0]
	assert.Equal("CATEGORY_INHOUSE", value.Name)
	assert.Equal(int32(0), value.Number)
	assert.Equal("com.example", value.Package)
	assert.Equal("The manufacturer is inhouse.", value.Comment)
	assert.True(value.IsProto3)

	value = enum.Values[1]
	assert.Equal("CATEGORY_EXTERNAL", value.Name)
	assert.Equal(int32(1), value.Number)
	assert.Equal("com.example", value.Package)
	assert.Equal("The manufacturer is external.", value.Comment)
	assert.True(value.IsProto3)
}

func (assert *Proto3ParserTest) TestMessageProperties() {
	msg := proto3File.GetMessage("Vehicle")
	assert.Equal("Vehicle", msg.Name)
	assert.Equal("com.example", msg.Package)
	assert.Equal("com.example.Vehicle", msg.FullName())
	assert.Equal("Represents a vehicle that can be hired.", msg.Comment)
	assert.True(msg.IsProto3)
	assert.Equal(7, len(msg.Fields))
	assert.Equal(0, len(msg.Extensions))

	assert.field(msg.Fields[0], "id", "Unique vehicle ID.", "int32", "")
	assert.field(msg.Fields[1], "model", "Vehicle model.", "com.example.Model", "")
	assert.field(msg.Fields[4], "category", "Vehicle category.", "com.example.Vehicle.Category", "")
	assert.field(msg.Fields[5], "rates", "rates", "sint32", "repeated")

	// maps are just repeated "<Name>Entry" fields
	assert.field(msg.Fields[6], "properties", "bag of properties related to the vehicle.", "com.example.Vehicle.PropertiesEntry", "repeated")
}

func (assert *Proto3ParserTest) TestNestedMessageProperties() {
	msg := proto3File.GetMessage("Vehicle.Category")
	assert.Equal("Vehicle.Category", msg.Name)
	assert.Equal("com.example", msg.Package)
	assert.Equal("com.example.Vehicle.Category", msg.FullName())
	assert.Equal("Represents a vehicle category. E.g. \"Sedan\" or \"Truck\".", msg.Comment)
	assert.True(msg.IsProto3)
	assert.Equal(2, len(msg.Fields))
	assert.Equal(0, len(msg.Extensions))

	assert.field(msg.Fields[0], "code", "Category code. E.g. \"S\".", "string", "")
	assert.field(msg.Fields[1], "description", "Category name. E.g. \"Sedan\".", "string", "")
}

func (assert *Proto3ParserTest) field(field *parser.Field, name, comment, typeName, label string) {
	assert.Equal(name, field.Name)
	assert.Equal(comment, field.Comment)
	assert.Equal(typeName, field.Type)
	assert.Equal(label, field.Label)
	assert.Equal("", field.DefaultValue)
}

func (assert *Proto3ParserTest) TestServiceProperties() {
	service := proto3File.GetService("VehicleService")
	assert.Equal("VehicleService", service.Name)
	assert.Equal("com.example", service.Package)
	assert.Equal("com.example.VehicleService", service.FullName())
	assert.Equal("The vehicle service.\n\nManages vehicles and such...", service.Comment)
	assert.True(service.IsProto3)
	assert.Equal(3, len(service.Methods))

	names := []string{"GetModels", "AddModels", "GetVehicle"}
	comments := []string{"Returns the set of models.", "creates models", "Looks up a vehicle by id."}
	clientStreams := []bool{false, true, false}
	serverStreams := []bool{true, true, false}
	requestTypes := []string{"com.example.EmptyMessage", "com.example.Model", "com.example.FindVehicleById"}
	responseTypes := []string{"com.example.Model", "com.example.Model", "com.example.Vehicle"}

	for idx, method := range service.Methods {
		assert.Equal(names[idx], method.Name)
		assert.Equal(comments[idx], method.Comment)
		assert.Equal(clientStreams[idx], method.ClientStreaming)
		assert.Equal(serverStreams[idx], method.ServerStreaming)
		assert.Equal(requestTypes[idx], method.RequestType)
		assert.Equal(responseTypes[idx], method.ResponseType)
		assert.Equal("com.example", method.Package)
		assert.True(method.IsProto3)
	}
}
