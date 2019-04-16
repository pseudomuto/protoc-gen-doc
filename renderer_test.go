package gendoc_test

import (
	"os"
	"testing"

	. "github.com/pseudomuto/protoc-gen-doc"
	"github.com/pseudomuto/protokit"
	"github.com/pseudomuto/protokit/utils"
	"github.com/stretchr/testify/suite"
)

const tempTestDir = "./tmp"

var renderTemplate *Template

type RendererTest struct {
	suite.Suite
}

func TestRenderer(t *testing.T) {
	suite.Run(t, new(RendererTest))
}

func (assert *RendererTest) SetupSuite() {
	set, err := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	assert.NoError(err)

	os.Mkdir(tempTestDir, os.ModePerm)

	req := utils.CreateGenRequest(set, "Booking.proto", "Vehicle.proto")
	result := protokit.ParseCodeGenRequest(req)
	renderTemplate = NewTemplate(result)
}

func (assert *RendererTest) TestDocBookRenderer() {
	_, err := RenderTemplate(RenderTypeDocBook, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestHtmlRenderer() {
	_, err := RenderTemplate(RenderTypeHTML, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestJsonRenderer() {
	_, err := RenderTemplate(RenderTypeJSON, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestMarkdownRenderer() {
	_, err := RenderTemplate(RenderTypeMarkdown, renderTemplate, "")
	assert.Nil(err)
}

func (assert *RendererTest) TestNewRenderType() {
	expected := []RenderType{
		RenderTypeDocBook,
		RenderTypeHTML,
		RenderTypeJSON,
		RenderTypeMarkdown,
	}

	supplied := []string{"docbook", "html", "json", "markdown"}

	for idx, input := range supplied {
		rt, err := NewRenderType(input)
		assert.Nil(err)
		assert.Equal(expected[idx], rt)
	}
}

func (assert *RendererTest) TestNewRenderTypeUnknown() {
	rt, err := NewRenderType("/some/template.tmpl")
	assert.Zero(rt)
	assert.NotNil(err)
}
