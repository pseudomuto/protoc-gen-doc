package gendoc_test

import (
	"github.com/pseudomuto/protoc-gen-doc"
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"github.com/pseudomuto/protoc-gen-doc/test"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

const tempTestDir = "./tmp"

var renderTemplate *gendoc.Template

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
	renderTemplate = gendoc.NewTemplate(result)
}

func (assert *RendererTest) TestDocBookRenderer() {
	_, err := gendoc.RenderTemplate(gendoc.RenderTypeDocBook, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestHtmlRenderer() {
	_, err := gendoc.RenderTemplate(gendoc.RenderTypeHTML, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestJsonRenderer() {
	_, err := gendoc.RenderTemplate(gendoc.RenderTypeJSON, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestMarkdownRenderer() {
	_, err := gendoc.RenderTemplate(gendoc.RenderTypeMarkdown, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestNewRenderType() {
	expected := []gendoc.RenderType{
		gendoc.RenderTypeDocBook,
		gendoc.RenderTypeHTML,
		gendoc.RenderTypeJSON,
		gendoc.RenderTypeMarkdown,
	}

	supplied := []string{"docbook", "html", "json", "markdown"}

	for idx, input := range supplied {
		rt, err := gendoc.NewRenderType(input)
		assert.Nil(err)
		assert.Equal(expected[idx], rt)
	}
}

func (assert *RendererTest) TestNewRenderTypeUnknown() {
	rt, err := gendoc.NewRenderType("/some/template.tmpl")
	assert.Zero(rt)
	assert.NotNil(err)
}
