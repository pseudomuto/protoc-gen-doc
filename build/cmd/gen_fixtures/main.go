// Functions like a normal code generator, except that it won't output anything. Is it used to store the
// CodeGeneratorRequest sent in by protoc for use in unit tests.
//
// Example:
//   protoc --plugin=protoc-gen-doc=./gen_fixtures --doc_out=. fixtures/proto3/*.proto
package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Could not read contents from stdin")
	}

	ioutil.WriteFile("fixtures/generator_request.dat", data, 0666)
}
