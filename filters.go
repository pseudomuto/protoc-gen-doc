package protoc_gen_doc

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

var paraPattern = regexp.MustCompile("(\\n|\\r|\\r\\n)\\s*")

func PFilter(content string) template.HTML {
	paragraphs := paraPattern.Split(content, -1)
	return template.HTML(fmt.Sprintf("<p>%s</p>", strings.Join(paragraphs, "</p><p>")))
}

func ParaFilter(content string) string {
	paragraphs := paraPattern.Split(content, -1)
	return fmt.Sprintf("<para>%s</para>", strings.Join(paragraphs, "</para><para>"))
}

func NoBrFilter(content string) string {
	withoutCR := strings.Replace(content, "\r", "", -1)
	withoutLF := strings.Replace(withoutCR, "\n", "", -1)

	return strings.Replace(withoutLF, "\r\n", "", -1)
}
