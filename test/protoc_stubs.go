package test

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"io/ioutil"
	"path"
	"runtime"
)

func MakeCodeGeneratorRequest() (*plugin_go.CodeGeneratorRequest, error) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "../fixtures/generator_request.dat")

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	req := new(plugin_go.CodeGeneratorRequest)
	if err = proto.Unmarshal(data, req); err != nil {
		return nil, err
	}

	return req, nil
}
