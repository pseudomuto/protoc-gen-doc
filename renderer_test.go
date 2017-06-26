package protoc_gen_doc_test

import (
	"github.com/pseudomuto/protoc-gen-doc"
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"github.com/pseudomuto/protoc-gen-doc/test"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

const tempTestDir = "./tmp"

var renderTemplate *protoc_gen_doc.Template

type RendererTest struct {
	suite.Suite
}

func TestRenderer(t *testing.T) {
	suite.Run(t, new(RendererTest))
}

func (assert *RendererTest) SetupSuite() {
	codeGenRequest, err := test.MakeCodeGeneratorRequest()
	assert.Nil(err)

	assert.Nil(os.Mkdir(tempTestDir, os.ModePerm))

	result := parser.ParseCodeRequest(codeGenRequest)
	renderTemplate = protoc_gen_doc.NewTemplate(result)
}

func (assert *RendererTest) TearDownSuite() {
	assert.Nil(os.RemoveAll(tempTestDir))
}

func (assert *RendererTest) TestJsonRenderer() {
	err := protoc_gen_doc.RenderTemplate(
		protoc_gen_doc.RenderTypeJson,
		renderTemplate,
		"",
		tempTestDir+"/output.json",
	)

	assert.Nil(err)
}
