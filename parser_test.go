package protoc_gen_doc_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ParserTest struct {
	suite.Suite
}

func TestParser(t *testing.T) {
	suite.Run(t, new(ParserTest))
}

func (assert *ParserTest) TestTheTruth() {
	assert.True(true)
}
