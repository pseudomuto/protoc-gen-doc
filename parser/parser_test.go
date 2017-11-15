package parser_test

import (
	"regexp"
	"testing"

	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"github.com/pseudomuto/protoc-gen-doc/test"
	"github.com/stretchr/testify/suite"
)

var (
	codeGenRequest             *plugin_go.CodeGeneratorRequest
	subject                    *parser.ParseResult
	subjectWithExcludePatterns *parser.ParseResult
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

	subject = parser.ParseCodeRequest(codeGenRequest, nil)
	excludePattern, _ := regexp.Compile("Booking*")
	subjectWithExcludePatterns = parser.ParseCodeRequest(codeGenRequest, []*regexp.Regexp{excludePattern})
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

func (assert *ParserTest) TestGetFileExcluded() {
	file := subjectWithExcludePatterns.GetFile("Vehicle.proto")
	assert.NotNil(file)
	assert.True(file.IsProto3)

	// Booking* excluded
	assert.Nil(subjectWithExcludePatterns.GetFile("Booking.proto"))
}
