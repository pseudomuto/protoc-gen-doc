package gendoc_test

import (
	"testing"

	"github.com/pseudomuto/protoc-gen-doc/parser"
	"github.com/pseudomuto/protoc-gen-doc/test"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	codeGenRequest, _ := test.MakeCodeGeneratorRequest()

	for i := 0; i < b.N; i++ {
		parser.ParseCodeRequest(codeGenRequest, nil)
	}
}
