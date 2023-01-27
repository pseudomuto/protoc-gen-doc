package gendoc

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
)

// PluginOptions encapsulates options for the plugin. The type of renderer, template file, and the name of the output
// file are included.
type PluginOptions struct {
	Type            RenderType
	TemplateFile    string
	OutputFile      string
	ExcludePatterns []*regexp.Regexp
	SourceRelative  bool
}

// SupportedFeatures describes a flag setting for supported features.
var SupportedFeatures = uint64(plugin_go.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

// Plugin describes a protoc code generate plugin. It's an implementation of Plugin from github.com/pseudomuto/protokit
type Plugin struct{}

// Generate compiles the documentation and generates the CodeGeneratorResponse to send back to protoc. It does this
// by rendering a template based on the options parsed from the CodeGeneratorRequest.
func (p *Plugin) Generate(r *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	options, err := ParseOptions(r)
	if err != nil {
		return nil, err
	}

	result := excludeUnwantedProtos(protokit.ParseCodeGenRequest(r), options.ExcludePatterns)

	customTemplate := ""

	if options.TemplateFile != "" {
		data, err := ioutil.ReadFile(options.TemplateFile)
		if err != nil {
			return nil, err
		}

		customTemplate = string(data)
	}

	resp := new(plugin_go.CodeGeneratorResponse)
	fdsGroup := groupProtosByDirectory(result, options.SourceRelative)

	// Sort the directory groups by directory key first...
	dirs := make([]string, 0, len(fdsGroup))
	for dir := range fdsGroup {
		dirs = append(dirs, dir)
	}
	sort.Slice(dirs, func(i, j int) bool {
		return pathCompare(dirs[i], dirs[j], 0)
	})
	// ...then, for each directory group, sort the files by their path (excluding last element, which is the filename)
	for _, dir := range dirs {
		fds := fdsGroup[dir]
		sort.Slice(fds, func(i, j int) bool {
			return pathCompare(fds[i].GetName(), fds[j].GetName(), 1)
		})

		template := NewTemplate(fds)

		output, err := RenderTemplate(options.Type, template, customTemplate)
		if err != nil {
			return nil, err
		}

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    proto.String(filepath.Join(dir, options.OutputFile)),
			Content: proto.String(string(output)),
		})
	}

	resp.SupportedFeatures = proto.Uint64(SupportedFeatures)

	return resp, nil
}

func groupProtosByDirectory(fds []*protokit.FileDescriptor, sourceRelative bool) map[string][]*protokit.FileDescriptor {
	fdsGroup := make(map[string][]*protokit.FileDescriptor)

	for _, fd := range fds {
		dir := ""
		if sourceRelative {
			dir, _ = filepath.Split(fd.GetName())
		}
		if dir == "" {
			dir = "./"
		}
		fdsGroup[dir] = append(fdsGroup[dir], fd)
	}
	return fdsGroup
}

func excludeUnwantedProtos(fds []*protokit.FileDescriptor, excludePatterns []*regexp.Regexp) []*protokit.FileDescriptor {
	descs := make([]*protokit.FileDescriptor, 0)

OUTER:
	for _, d := range fds {
		for _, p := range excludePatterns {
			if p.MatchString(d.GetName()) {
				continue OUTER
			}
		}

		descs = append(descs, d)
	}

	return descs
}

// A comparison function taking two paths with "/" as separator and a number of elements to ignore at the end.
// It sorts a list of paths "top-level-first", such that they end up forming a clean directory tree, like:
// - /top1/sub1/
// - /top1/sub1/bot1/
// - /top1/sub1/bot2/
// - /top1/sub2/
// - /top2/
// - /top2/sub1/
func pathCompare(i, j string, ignoreAtEnd int) bool {
	pi, pj := strings.Split(i, "/"), strings.Split(j, "/")

	minLen := len(pi)
	if minLen > len(pj) {
		minLen = len(pj)
	}
	minLen -= ignoreAtEnd

	for k := 0; k < minLen; k++ {
		comp := strings.Compare(pi[k], pj[k])
		if comp != 0 {
			return comp < 0
		}
	}

	return minLen < (len(pj) - ignoreAtEnd)
}

// ParseOptions parses plugin options from a CodeGeneratorRequest. It does this by splitting the `Parameter` field from
// the request object and parsing out the type of renderer to use and the name of the file to be generated.
//
// The parameter (`--doc_opt`) must be of the format <TYPE|TEMPLATE_FILE>,<OUTPUT_FILE>[,default|source_relative]:<EXCLUDE_PATTERN>,<EXCLUDE_PATTERN>*.
// The file will be written to the directory specified with the `--doc_out` argument to protoc.
func ParseOptions(req *plugin_go.CodeGeneratorRequest) (*PluginOptions, error) {
	options := &PluginOptions{
		Type:           RenderTypeHTML,
		TemplateFile:   "",
		OutputFile:     "index.html",
		SourceRelative: false,
	}

	params := req.GetParameter()
	if strings.Contains(params, ":") {
		// Parse out exclude patterns if any
		parts := strings.Split(params, ":")
		for _, pattern := range strings.Split(parts[1], ",") {
			r, err := regexp.Compile(pattern)
			if err != nil {
				return nil, err
			}
			options.ExcludePatterns = append(options.ExcludePatterns, r)
		}
		// The first part is parsed below
		params = parts[0]
	}
	if params == "" {
		return options, nil
	}

	if !strings.Contains(params, ",") {
		return nil, fmt.Errorf("Invalid parameter: %s", params)
	}

	parts := strings.Split(params, ",")
	if len(parts) < 2 || len(parts) > 3 {
		return nil, fmt.Errorf("Invalid parameter: %s", params)
	}

	options.TemplateFile = parts[0]
	options.OutputFile = path.Base(parts[1])
	if len(parts) > 2 {
		switch parts[2] {
		case "source_relative":
			options.SourceRelative = true
		case "default":
			options.SourceRelative = false
		default:
			return nil, fmt.Errorf("Invalid parameter: %s", params)
		}
	}
	options.SourceRelative = len(parts) > 2 && parts[2] == "source_relative"

	renderType, err := NewRenderType(options.TemplateFile)
	if err == nil {
		options.Type = renderType
		options.TemplateFile = ""
	}

	return options, nil
}
