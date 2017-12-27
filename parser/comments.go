package parser

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"regexp"
	"strconv"
	"strings"
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

func scrubComment(s string) string {
	lines := strings.Split(s, "\n")
	re := regexp.MustCompile(`[/*]* (.*)$`)
	for idx, line := range lines {
		if re.MatchString(line) {
			lines[idx] = strings.TrimRight(re.ReplaceAllString(line, "$1"), " ")
		} else {
			lines[idx] = ""
		}
	}

	return strings.Trim(strings.Join(lines, "\n"), "\n")
}
