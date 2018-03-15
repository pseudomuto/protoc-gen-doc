package gendoc_test

import (
	"github.com/pseudomuto/protokit"
	"github.com/pseudomuto/protokit/utils"
	"github.com/stretchr/testify/suite"

	"os"
	"testing"

	"github.com/ilackarms/protoc-gen-doc"
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
	set, err := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	assert.NoError(err)

	os.Mkdir(tempTestDir, os.ModePerm)

	req := utils.CreateGenRequest(set, "Booking.proto", "Vehicle.proto")
	result := protokit.ParseCodeGenRequest(req)
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
