package gendoc

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

var paraPattern = regexp.MustCompile("(\\n|\\r|\\r\\n)\\s*")

// PFilter splits the content by new lines and wraps each one in a <p> tag.
func PFilter(content string) template.HTML {
	paragraphs := paraPattern.Split(content, -1)
	return template.HTML(fmt.Sprintf("<p>%s</p>", strings.Join(paragraphs, "</p><p>")))
}

// ParaFilter splits the content by new lines and wraps each one in a <para> tag.
func ParaFilter(content string) string {
	paragraphs := paraPattern.Split(content, -1)
	return fmt.Sprintf("<para>%s</para>", strings.Join(paragraphs, "</para><para>"))
}

// NoBrFilter removes CR and LF from content
func NoBrFilter(content string) string {
	withoutCR := strings.Replace(content, "\r", "", -1)
	withoutLF := strings.Replace(withoutCR, "\n", "", -1)

	return strings.Replace(withoutLF, "\r\n", "", -1)
}
