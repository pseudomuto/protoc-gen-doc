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

	os.Mkdir(tempTestDir, os.ModePerm)

	result := parser.ParseCodeRequest(codeGenRequest)
	renderTemplate = protoc_gen_doc.NewTemplate(result)
}

func (assert *RendererTest) TestDocBookRenderer() {
	err := protoc_gen_doc.RenderTemplate(
		protoc_gen_doc.RenderTypeDocBook,
		renderTemplate,
		tempTestDir+"/output.docbook",
	)

	assert.Nil(err)
}

func (assert *RendererTest) TestHtmlRenderer() {
	err := protoc_gen_doc.RenderTemplate(
		protoc_gen_doc.RenderTypeHtml,
		renderTemplate,
		tempTestDir+"/output.html",
	)

	assert.Nil(err)
}

func (assert *RendererTest) TestJsonRenderer() {
	err := protoc_gen_doc.RenderTemplate(
		protoc_gen_doc.RenderTypeJson,
		renderTemplate,
		tempTestDir+"/output.json",
	)

	assert.Nil(err)
}

func (assert *RendererTest) TestMarkdownRenderer() {
	err := protoc_gen_doc.RenderTemplate(
		protoc_gen_doc.RenderTypeMarkdown,
		renderTemplate,
		tempTestDir+"/output.md",
	)

	assert.Nil(err)
}
