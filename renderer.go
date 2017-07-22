package protoc_gen_doc

import (
	"bytes"
	"encoding/json"
	"errors"
	html_template "html/template"
	text_template "text/template"
)

type RenderType int8

const (
	_ RenderType = iota
	RenderTypeDocBook
	RenderTypeHtml
	RenderTypeJson
	RenderTypeMarkdown
)

func NewRenderType(renderType string) (RenderType, error) {
	switch renderType {
	case "docbook":
		return RenderTypeDocBook, nil
	case "html":
		return RenderTypeHtml, nil
	case "json":
		return RenderTypeJson, nil
	case "markdown":
		return RenderTypeMarkdown, nil
	}

	return 0, errors.New("Invalid render type")
}

func (rt RenderType) renderer() (Processor, error) {
	tmpl, err := rt.template()
	if err != nil {
		return nil, err
	}

	switch rt {
	case RenderTypeDocBook:
		return &textRenderer{string(tmpl)}, nil
	case RenderTypeHtml:
		return &htmlRenderer{string(tmpl)}, nil
	case RenderTypeJson:
		return new(jsonRenderer), nil
	case RenderTypeMarkdown:
		return &htmlRenderer{string(tmpl)}, nil
	}

	return nil, errors.New("Unable to create a processor")
}

func (rt RenderType) template() ([]byte, error) {
	switch rt {
	case RenderTypeDocBook:
		return fetchResource("docbook.tmpl")
	case RenderTypeHtml:
		return fetchResource("html.tmpl")
	case RenderTypeJson:
		return nil, nil
	case RenderTypeMarkdown:
		return fetchResource("markdown.tmpl")
	}

	return nil, errors.New("Couldn't find template for render type")
}

var funcMap = map[string]interface{}{
	"p":    PFilter,
	"para": ParaFilter,
	"nobr": NoBrFilter,
}

type Processor interface {
	Apply(template *Template) ([]byte, error)
}

func RenderTemplate(kind RenderType, template *Template, inputTemplate string) ([]byte, error) {
	if inputTemplate != "" {
		processor := &textRenderer{inputTemplate}
		return processor.Apply(template)
	}

	processor, err := kind.renderer()
	if err != nil {
		return nil, err
	}

	return processor.Apply(template)
}

type textRenderer struct {
	inputTemplate string
}

func (mr *textRenderer) Apply(template *Template) ([]byte, error) {
	tmpl, err := text_template.New("Text Template").Funcs(funcMap).Parse(mr.inputTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, template); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type htmlRenderer struct {
	inputTemplate string
}

func (mr *htmlRenderer) Apply(template *Template) ([]byte, error) {
	tmpl, err := html_template.New("Text Template").Funcs(funcMap).Parse(mr.inputTemplate)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, template); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type jsonRenderer struct{}

func (r *jsonRenderer) Apply(template *Template) ([]byte, error) {
	return json.MarshalIndent(template, "", "  ")
}
