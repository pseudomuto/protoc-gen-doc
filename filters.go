package gendoc

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"

  "github.com/gosimple/slug"
)

var (
	paraPattern         = regexp.MustCompile(`(\n|\r|\r\n)\s*`)
	spacePattern        = regexp.MustCompile("( )+")
	multiNewlinePattern = regexp.MustCompile(`(\r\n|\r|\n){2,}`)
	specialCharsPattern = regexp.MustCompile(`[^a-zA-Z0-9_-]`)

	capDigitsRegex      = regexp.MustCompile(`([A-Z]{2,})(\d+)`)
	lowerCapDigitsRegex = regexp.MustCompile(`([a-z\d]+)([A-Z]{2,})`)
	lowerUpperRegex     = regexp.MustCompile(`([a-z\d])([A-Z])`)
  // `[a-rt-z]` matches all lowercase characters except `s`.
  // This avoids matching plural acronyms like `APIs`.
	pluralAcronymsRegex = regexp.MustCompile(`([A-Z]+)([A-Z][a-rt-z\d]+)`)
)

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

// NoBrFilter removes single CR and LF from content.
func NoBrFilter(content string) string {
	normalized := strings.Replace(content, "\r\n", "\n", -1)
	paragraphs := multiNewlinePattern.Split(normalized, -1)
	for i, p := range paragraphs {
		withoutCR := strings.Replace(p, "\r", " ", -1)
		withoutLF := strings.Replace(withoutCR, "\n", " ", -1)
		paragraphs[i] = spacePattern.ReplaceAllString(withoutLF, " ")
	}
	return strings.Join(paragraphs, "\n\n")
}

func decamelize(input string) string {
	// Replace the matched patterns with spaces to decamelize the string.
	input = capDigitsRegex.ReplaceAllString(input, "$1 $2")
	input = lowerCapDigitsRegex.ReplaceAllString(input, "$1 $2")
	input = lowerUpperRegex.ReplaceAllString(input, "$1 $2")
	input = pluralAcronymsRegex.ReplaceAllString(input, "$1 $2")

	return input
}

func AnchorFilter(str string) string {
  return slug.Make(decamelize(str))
}
