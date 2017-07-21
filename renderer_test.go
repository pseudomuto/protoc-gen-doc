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
	_, err := protoc_gen_doc.RenderTemplate(protoc_gen_doc.RenderTypeDocBook, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestHtmlRenderer() {
	_, err := protoc_gen_doc.RenderTemplate(protoc_gen_doc.RenderTypeHtml, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestJsonRenderer() {
	_, err := protoc_gen_doc.RenderTemplate(protoc_gen_doc.RenderTypeJson, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestMarkdownRenderer() {
	_, err := protoc_gen_doc.RenderTemplate(protoc_gen_doc.RenderTypeMarkdown, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestNewRenderType() {
	expected := []protoc_gen_doc.RenderType{
		protoc_gen_doc.RenderTypeDocBook,
		protoc_gen_doc.RenderTypeHtml,
		protoc_gen_doc.RenderTypeJson,
		protoc_gen_doc.RenderTypeMarkdown,
	}

	supplied := []string{"docbook", "html", "json", "markdown"}

	for idx, input := range supplied {
		rt, err := protoc_gen_doc.NewRenderType(input)
		assert.Nil(err)
		assert.Equal(expected[idx], rt)
	}
}

func (assert *RendererTest) TestNewRenderTypeUnknown() {
	rt, err := protoc_gen_doc.NewRenderType("/some/template.tmpl")
	assert.Zero(rt)
	assert.NotNil(err)
}
