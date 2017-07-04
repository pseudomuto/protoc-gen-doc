package protoc_gen_doc

import (
	"bytes"
	"encoding/json"
	html_template "html/template"
	"io/ioutil"
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

type Processor interface {
	Apply(template *Template) ([]byte, error)
}

func RenderTemplate(kind RenderType, template *Template, outputPath string) error {
	var processor Processor

	switch kind {
	case RenderTypeDocBook:
		res, err := fetchResource("docbook.tmpl")
		if err != nil {
			return err
		}

		processor = &textRenderer{string(res)}
	case RenderTypeHtml:
		res, err := fetchResource("html.tmpl")
		if err != nil {
			return err
		}

		processor = &htmlRenderer{string(res)}
	case RenderTypeJson:
		processor = &jsonRenderer{}
	case RenderTypeMarkdown:
		res, err := fetchResource("markdown.tmpl")
		if err != nil {
			return err
		}

		processor = &htmlRenderer{string(res)}
	}

	result, err := processor.Apply(template)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, result, 0644)
}

type textRenderer struct {
	inputTemplate string
}

func (mr *textRenderer) Apply(template *Template) ([]byte, error) {
	tmpl, err := text_template.New("Text Template").Parse(mr.inputTemplate)
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
	tmpl, err := html_template.New("Text Template").Parse(mr.inputTemplate)
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
	return json.MarshalIndent(template.Files, "", "  ")
}
