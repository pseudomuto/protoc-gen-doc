package main_test

import (
	"bytes"
	"testing"

	. "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
	"github.com/stretchr/testify/suite"
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
		f := ParseFlags(new(bytes.Buffer), test.args)
		assert.Equal(test.result, HandleFlags(f))
	}
}
