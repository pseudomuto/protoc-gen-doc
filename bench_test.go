package protoc_gen_doc

import (
	"github.com/pseudomuto/protoc-gen-doc/parser"
	"github.com/pseudomuto/protoc-gen-doc/test"
	"testing"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	codeGenRequest, _ := test.MakeCodeGeneratorRequest()

	for i := 0; i < b.N; i++ {
		parser.ParseCodeRequest(codeGenRequest)
	}
}
