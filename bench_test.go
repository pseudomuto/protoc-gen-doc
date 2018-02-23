package gendoc_test

import (
	"testing"

	"github.com/pseudomuto/protoc-gen-doc"
	"github.com/pseudomuto/protoc-gen-doc/test"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	codeGenRequest, _ := test.MakeCodeGeneratorRequest()
	plugin := new(gendoc.Plugin)

	for i := 0; i < b.N; i++ {
		plugin.Generate(codeGenRequest)
	}
}
