package protoc_gen_doc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"io/ioutil"
	"path"
	"strings"
)

// <docbook|html|json|markdown|TEMPLATE_FILE>,<OUTPUT>
type PluginOptions struct {
	Type         RenderType
	TemplateFile string
	OutputFile   string
}

func ParseOptions(req *plugin_go.CodeGeneratorRequest) (*PluginOptions, error) {
	options := &PluginOptions{
		Type:         RenderTypeHtml,
		TemplateFile: "",
		OutputFile:   "index.html",
	}

	params := req.GetParameter()
	if params == "" {
		return options, nil
	}

	if !strings.Contains(params, ",") {
		return nil, fmt.Errorf("Invalid parameter: %s", params)
	}

	parts := strings.Split(params, ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid parameter: %s", params)
	}

	options.TemplateFile = parts[0]
	options.OutputFile = path.Base(parts[1])

	renderType, err := NewRenderType(options.TemplateFile)
	if err == nil {
		options.Type = renderType
		options.TemplateFile = ""
	}

	return options, nil
}

func RunPlugin(request *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	result := parser.ParseCodeRequest(request)
	template := NewTemplate(result)

	options, err := ParseOptions(request)
	if err != nil {
		return nil, err
	}

	customTemplate := ""

	if options.TemplateFile != "" {
		data, err := ioutil.ReadFile(options.TemplateFile)
		if err != nil {
			return nil, err
		}

		customTemplate = string(data)
	}

	output, err := RenderTemplate(options.Type, template, customTemplate)
	if err != nil {
		return nil, err
	}

	resp := new(plugin_go.CodeGeneratorResponse)
	resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
		Name:    proto.String(options.OutputFile),
		Content: proto.String(string(output)),
	})

	return resp, nil
}
