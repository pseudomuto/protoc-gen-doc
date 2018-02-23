package main_test

import (
	"github.com/stretchr/testify/suite"

	"bytes"
	"testing"

	"github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)

type MainTest struct {
	suite.Suite
}

func TestMain(t *testing.T) {
	suite.Run(t, new(MainTest))
}

func (assert *MainTest) TestHandleFlags() {
	tests := []struct {
		args   []string
		result bool
	}{
		{[]string{"app", "-help"}, true},
		{[]string{"app", "-version"}, true},
		{[]string{"app", "-wjat"}, true},
	}

	for _, test := range tests {
		f := main.ParseFlags(new(bytes.Buffer), test.args)
		assert.Equal(test.result, main.HandleFlags(f))
	}
}
