package protoc_gen_doc

import (
	"encoding/json"
	"io/ioutil"
)

type RenderType int8

const (
	_ RenderType = iota
	RenderTypeJson
)

type Processor interface {
	Apply(template *Template) ([]byte, error)
}

func RenderTemplate(kind RenderType, template *Template, inputTemplate, outputPath string) error {
	var processor Processor

	switch kind {
	case RenderTypeJson:
		processor = &jsonRenderer{}
	}

	result, err := processor.Apply(template)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, result, 0644)
}

type jsonRenderer struct{}

func (r *jsonRenderer) Apply(template *Template) ([]byte, error) {
	return json.Marshal(template)
}
