package protoc_gen_doc_test

import (
	"github.com/pseudomuto/protoc-gen-doc"
	"github.com/stretchr/testify/suite"
	html "html/template"
	"testing"
)

type FilterTest struct {
	suite.Suite
}

func TestFilter(t *testing.T) {
	suite.Run(t, new(FilterTest))
}

func (assert *FilterTest) TestPFilter() {
	tests := map[string]string{
		"Some content.":                          "<p>Some content.</p>",
		"Some content.\nRight here.":             "<p>Some content.</p><p>Right here.</p>",
		"Some content.\r\nRight here.":           "<p>Some content.</p><p>Right here.</p>",
		"Some content.\n\tRight here.":           "<p>Some content.</p><p>Right here.</p>",
		"Some content.\r\n\n  \r\n  Right here.": "<p>Some content.</p><p>Right here.</p>",
	}

	for input, output := range tests {
		assert.Equal(html.HTML(output), protoc_gen_doc.PFilter(input))
	}
}

func (assert *FilterTest) TestParaFilter() {
	tests := map[string]string{
		"Some content.":                          "<para>Some content.</para>",
		"Some content.\nRight here.":             "<para>Some content.</para><para>Right here.</para>",
		"Some content.\r\nRight here.":           "<para>Some content.</para><para>Right here.</para>",
		"Some content.\n\tRight here.":           "<para>Some content.</para><para>Right here.</para>",
		"Some content.\r\n\n  \r\n  Right here.": "<para>Some content.</para><para>Right here.</para>",
	}

	for input, output := range tests {
		assert.Equal(output, protoc_gen_doc.ParaFilter(input))
	}
}

func (assert *FilterTest) TestNoBrFilter() {
	tests := map[string]string{
		"My content":                     "My content",
		"My content \r\nHere.":           "My content Here.",
		"My\n content\r right\r\n here.": "My content right here.",
	}

	for input, output := range tests {
		assert.Equal(output, protoc_gen_doc.NoBrFilter(input))
	}
}
