package parser_test

import (
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"github.com/pseudomuto/protoc-gen-doc/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	codeGenRequest *plugin_go.CodeGeneratorRequest
	subject        *parser.ParseResult
)

type ParserTest struct {
	suite.Suite
}

func TestParser(t *testing.T) {
	suite.Run(t, new(ParserTest))
}

func (assert *ParserTest) SetupSuite() {
	var err error
	codeGenRequest, err = test.MakeCodeGeneratorRequest()
	assert.Nil(err)

	subject = parser.ParseCodeRequest(codeGenRequest)
}

func (assert *ParserTest) TestGetFile() {
	assert.NotNil(subject.GetFile("Vehicle.proto"))
	assert.Nil(subject.GetFile("Unknown.proto"))
}

func (assert *ParserTest) TestFileProperties() {
	file := subject.GetFile("Vehicle.proto")
	assert.Equal("Vehicle.proto", file.Name)
	assert.Equal("com.example", file.Package)
	assert.Equal("Messages describing manufacturers / vehicles.", file.Comment)
	assert.True(file.IsProto3)

	for _, msg := range []string{"EmptyMessage", "Manufacturer", "Model", "Vehicle", "Vehicle.Category"} {
		assert.True(file.HasMessage(msg))
	}

	for _, enum := range []string{"Manufacturer.Category", "Type"} {
		assert.True(file.HasEnum(enum))
	}
}

func (assert *ParserTest) TestEnumProperties() {
	enum := subject.GetFile("Vehicle.proto").GetEnum("Type")
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

func (assert *ParserTest) TestNestedEnumProperties() {
	enum := subject.GetFile("Vehicle.proto").GetEnum("Manufacturer.Category")
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

func (assert *ParserTest) TestMessageProperties() {
	msg := subject.GetFile("Vehicle.proto").GetMessage("Vehicle")
	assert.Equal("Vehicle", msg.Name)
	assert.Equal("com.example", msg.Package)
	assert.Equal("com.example.Vehicle", msg.FullName())
	assert.Equal("Represents a vehicle that can be hired.", msg.Comment)
	assert.True(msg.IsProto3)
	assert.Equal(7, len(msg.Fields))

	assert.field(msg.Fields[0], "id", "Unique vehicle ID.", "int32", "")
	assert.field(msg.Fields[1], "model", "Vehicle model.", "Model", "")
	assert.field(msg.Fields[4], "category", "Vehicle category.", "Vehicle.Category", "")
	assert.field(msg.Fields[5], "rates", "rates", "sint32", "repeated")

	// maps are just repeated "<Name>Entry" fields
	assert.field(msg.Fields[6], "properties", "bag of properties related to the vehicle.", "Vehicle.PropertiesEntry", "repeated")
}

func (assert *ParserTest) TestNestedMessageProperties() {
	msg := subject.GetFile("Vehicle.proto").GetMessage("Vehicle.Category")
	assert.Equal("Vehicle.Category", msg.Name)
	assert.Equal("com.example", msg.Package)
	assert.Equal("com.example.Vehicle.Category", msg.FullName())
	assert.Equal("Represents a vehicle category. E.g. \"Sedan\" or \"Truck\".", msg.Comment)
	assert.True(msg.IsProto3)
	assert.Equal(2, len(msg.Fields))

	assert.field(msg.Fields[0], "code", "Category code. E.g. \"S\".", "string", "")
	assert.field(msg.Fields[1], "description", "Category name. E.g. \"Sedan\".", "string", "")
}

func (assert *ParserTest) field(field *parser.Field, name, comment, typeName, label string) {
	assert.Equal(name, field.Name)
	assert.Equal(comment, field.Comment)
	assert.Equal(typeName, field.Type)
	assert.Equal(label, field.Label)
	assert.Equal("", field.DefaultValue)
}

func (assert *ParserTest) TestServiceProperties() {
	service := subject.GetFile("Vehicle.proto").GetService("VehicleService")
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
