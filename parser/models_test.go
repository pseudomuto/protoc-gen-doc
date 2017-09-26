package parser_test

import (
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"github.com/pseudomuto/protoc-gen-doc/test"
	"github.com/stretchr/testify/suite"
	"testing"
)

var file *parser.File

type ModelsTest struct {
	suite.Suite
}

func TestModels(t *testing.T) {
	suite.Run(t, new(ModelsTest))
}

func (assert *ModelsTest) SetupSuite() {
	var err error
	codeGenRequest, err = test.MakeCodeGeneratorRequest()
	assert.Nil(err)

	file = parser.ParseCodeRequest(codeGenRequest).GetFile("Vehicle.proto")
	assert.NotNil(file)
}

func (assert *ModelsTest) TestFileEnums() {
	assert.NotNil(file.GetEnum("Type"))
	assert.True(file.HasEnum("Type"))

	assert.Nil(file.GetEnum("Unknown"))
	assert.False(file.HasEnum("Unknown"))
}

func (assert *ModelsTest) TestFileMessages() {
	assert.NotNil(file.GetMessage("Vehicle"))
	assert.True(file.HasMessage("Vehicle"))

	assert.Nil(file.GetMessage("Unknown"))
	assert.False(file.HasMessage("Unknown"))
}

func (assert *ModelsTest) TestFileServices() {
	assert.NotNil(file.GetService("VehicleService"))
	assert.True(file.HasService("VehicleService"))

	assert.Nil(file.GetService("Unknown"))
	assert.False(file.HasService("Unknown"))
}
