package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"strconv"
	"strings"
	"regexp"
)

type commentContainer interface {
	getPath() string
	setComment(comment string)
}

type commentExtractor struct {
	comments map[string]string
}

func newCommentExtractor(fd *descriptor.FileDescriptorProto) *commentExtractor {
	comments := make(map[string]string)

	for _, loc := range fd.GetSourceCodeInfo().GetLocation() {
		if loc.LeadingComments == nil && loc.TrailingComments == nil {
			continue
		}

		var p []string
		for _, n := range loc.GetPath() {
			p = append(p, strconv.Itoa(int(n)))
		}

		comments[strings.Join(p, ",")] = strings.Join(
			[]string{loc.GetLeadingComments(), loc.GetTrailingComments()},
			"\n\n",
		)
	}

	return &commentExtractor{comments}
}

func (ce *commentExtractor) extractComments(containers ...commentContainer) {
	for _, container := range containers {
		container.setComment(ce.commentForPath(container.getPath()))
	}
}

func (ce *commentExtractor) commentForPath(path string) string {
	return scrubComment(ce.comments[path])
}

//
// SourceCodeInfo_Location already strips out standard comment markers.
// This function is used to remove additional comment decorations
// at the beginning and end of each line.
// 
// 1) Remove Line Start plus any / or * followed by line end 
//    that is a complete comment decoration line
// 2) Remove Line Start plus any / or * followed by whitepace 
//    that is a comment decoration at start of line
// 3) Remove Whitespace plus any / or * followed by line end 
//    that is a comment decoration at end of line
//
func scrubComment(s string) string {
	var rePrefix = regexp.MustCompile(`^/*(\**|/*)(\s+|$)`)
	var reSuffix = regexp.MustCompile(`\s+(\**|/*)$`)
	lines := strings.Split(s, "\n")

	for idx, line := range lines {
		line = strings.TrimRight(line, " /\n")
		line = rePrefix.ReplaceAllString(line, "")
		line = reSuffix.ReplaceAllString(line, "")
		lines[idx] = strings.TrimLeft(line, " \n")
	}

	return strings.Trim(strings.Join(lines, "\n"), "\n")
}
