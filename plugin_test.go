package protoc_gen_doc_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protoc-gen-doc"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PluginTest struct {
	suite.Suite
}

func TestPlugin(t *testing.T) {
	suite.Run(t, new(PluginTest))
}

func (assert *PluginTest) TestParseOptionsForBuiltinTemplates() {
	results := map[string]string{
		"docbook":  "output.xml",
		"html":     "output.html",
		"json":     "output.json",
		"markdown": "output.md",
	}

	for kind, file := range results {
		req := new(plugin_go.CodeGeneratorRequest)
		req.Parameter = proto.String(kind + "," + file)

		options, err := protoc_gen_doc.ParseOptions(req)
		assert.Nil(err)

		renderType, err := protoc_gen_doc.NewRenderType(kind)
		assert.Nil(err)

		assert.Equal(renderType, options.Type)
		assert.Equal("", options.TemplateFile)
		assert.Equal(file, options.OutputFile)
	}
}

func (assert *PluginTest) TestParseOptionsForCustomTemplate() {
	req := new(plugin_go.CodeGeneratorRequest)
	req.Parameter = proto.String("/path/to/template.tmpl,/base/name/only/output.md")

	options, err := protoc_gen_doc.ParseOptions(req)
	assert.Nil(err)

	assert.Equal(protoc_gen_doc.RenderTypeHtml, options.Type)
	assert.Equal("/path/to/template.tmpl", options.TemplateFile)
	assert.Equal("output.md", options.OutputFile)
}

func (assert *PluginTest) TestParseOptionsWithInvalidValues() {
	badValues := []string{
		"markdown",
		"html",
		"/some/path.tmpl",
		"more,than,1,comma",
	}

	for _, value := range badValues {
		req := new(plugin_go.CodeGeneratorRequest)
		req.Parameter = proto.String(value)

		_, err := protoc_gen_doc.ParseOptions(req)
		assert.NotNil(err)
	}
}

func (assert *PluginTest) TestRunPluginForBuiltinTemplate() {
	req := new(plugin_go.CodeGeneratorRequest)
	req.Parameter = proto.String("markdown,/base/name/only/output.md")

	resp, err := protoc_gen_doc.RunPlugin(req)
	assert.Nil(err)

	assert.Equal(1, len(resp.File))
	assert.Equal("output.md", resp.File[0].GetName())
	assert.NotEmpty(resp.File[0].GetContent())
}

func (assert *PluginTest) TestRunPluginForCustomTemplate() {
	req := new(plugin_go.CodeGeneratorRequest)
	req.Parameter = proto.String("templates/html.tmpl,/base/name/only/output.html")

	resp, err := protoc_gen_doc.RunPlugin(req)
	assert.Nil(err)

	assert.Equal(1, len(resp.File))
	assert.Equal("output.html", resp.File[0].GetName())
	assert.NotEmpty(resp.File[0].GetContent())
}

func (assert *PluginTest) TestRunPluginWithInvalidOptions() {
	req := new(plugin_go.CodeGeneratorRequest)
	req.Parameter = proto.String("html")

	_, err := protoc_gen_doc.RunPlugin(req)
	assert.NotNil(err)
}
