package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/plugin"
)

type ParseResult struct {
	Files []*File
}

func (pr *ParseResult) GetFile(name string) *File {
	for _, f := range pr.Files {
		if f.Name == name {
			return f
		}
	}

	return nil
}

func ParseCodeRequest(req *plugin_go.CodeGeneratorRequest) *ParseResult {
	result := new(ParseResult)

	for _, file := range req.GetProtoFile() {
		result.Files = append(result.Files, parseProtoFile(file))
	}

	return result
}
