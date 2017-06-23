package test

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"io/ioutil"
)

func MakeCodeGeneratorRequest() (*plugin_go.CodeGeneratorRequest, error) {
	data, err := ioutil.ReadFile("../fixtures/generator_request.dat")
	if err != nil {
		return nil, err
	}

	req := new(plugin_go.CodeGeneratorRequest)
	if err = proto.Unmarshal(data, req); err != nil {
		return nil, err
	}

	return req, nil
}
