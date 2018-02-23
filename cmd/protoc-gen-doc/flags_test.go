package main_test

import (
	"github.com/stretchr/testify/suite"

	"bytes"
	"fmt"
	"testing"

	"github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)

type FlagsTest struct {
	suite.Suite
}

func TestFlags(t *testing.T) {
	suite.Run(t, new(FlagsTest))
}

func (assert *FlagsTest) TestCode() {
	f := main.ParseFlags(nil, []string{"app", "-help"})
	assert.Equal(0, f.Code())

	f = main.ParseFlags(nil, []string{"app", "-whoawhoawhoa"})
	assert.Equal(1, f.Code())
}

func (assert *FlagsTest) TestHasMatch() {
	f := main.ParseFlags(nil, []string{"app", "-help"})
	assert.True(f.HasMatch())

	f = main.ParseFlags(nil, []string{"app", "-version"})
	assert.True(f.HasMatch())

	f = main.ParseFlags(nil, []string{"app", "-watthewhat"})
	assert.True(f.HasMatch())

	f = main.ParseFlags(nil, []string{"app"})
	assert.False(f.HasMatch())
}

func (assert *FlagsTest) TestShowHelp() {
	f := main.ParseFlags(nil, []string{"app", "-help"})
	assert.True(f.ShowHelp())

	f = main.ParseFlags(nil, []string{"app", "-version"})
	assert.False(f.ShowHelp())
}

func (assert *FlagsTest) TestShowVersion() {
	f := main.ParseFlags(nil, []string{"app", "-version"})
	assert.True(f.ShowVersion())

	f = main.ParseFlags(nil, []string{"app", "-help"})
	assert.False(f.ShowVersion())
}

func (assert *FlagsTest) TestPrintHelp() {
	buf := new(bytes.Buffer)

	f := main.ParseFlags(buf, []string{"app"})
	f.PrintHelp()

	result := buf.String()
	assert.Contains(result, "Usage of app:\n\n")
	assert.Contains(result, "FLAGS\n")
	assert.Contains(result, "-help")
	assert.Contains(result, "-version")
}

func (assert *FlagsTest) TestPrintVersion() {
	buf := new(bytes.Buffer)

	f := main.ParseFlags(buf, []string{"app"})
	f.PrintVersion()

	// Normally, I'm not a fan of using constants like this in tests. However, having this break everytime the version
	// changes is kinda poop, so I've used VERSION here.
	assert.Equal(fmt.Sprintf("app version %s\n", main.Version()), buf.String())
}

func (assert *FlagsTest) TestInvalidFlags() {
	buf := new(bytes.Buffer)

	f := main.ParseFlags(buf, []string{"app", "-wat"})
	assert.Contains(buf.String(), "flag provided but not defined: -wat\n")
	assert.True(f.HasMatch())
	assert.True(f.ShowHelp())
}
