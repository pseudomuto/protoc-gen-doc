package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protoc-gen-doc"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Could not read contents from stdin")
	}

	req := new(plugin_go.CodeGeneratorRequest)
	if err = proto.Unmarshal(input, req); err != nil {
		log.Fatal(err)
	}

	resp, err := protoc_gen_doc.RunPlugin(req)
	if err != nil {
		log.Fatal(err)
	}

	data, err := proto.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stdout.Write(data); err != nil {
		log.Fatal(err)
	}
}
