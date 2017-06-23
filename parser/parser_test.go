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
	file := subject.GetFile("Vehicle.proto")
	assert.NotNil(file)
	assert.True(file.IsProto3)

	file = subject.GetFile("Booking.proto")
	assert.NotNil(file)
	assert.False(file.IsProto3)

	assert.Nil(subject.GetFile("Unknown.proto"))
}
