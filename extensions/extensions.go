// Package extensions implements a system for working with extended options.
package extensions

import (
	"strings"
	"github.com/golang/protobuf/proto"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Transformer functions for transforming payloads of an extension option into
// something that can be rendered by a template.
type Transformer func(payload interface{}) interface{}

var transformers = make(map[string]Transformer)
var defaultTransformer Transformer
var transformNameAliases = make(map[string]string)

// SetTransformer sets the transformer function for the given extension name
func SetTransformer(extensionName string, f Transformer) {
	transformers[extensionName] = f
}

// SetNameAlias sets an alias for an extension name, allowing different keys
// to map to the same transformer
func SetNameAlias(fullName, aliasTo string) {
	transformNameAliases[fullName] = aliasTo
}

// SetDefaultTransformer sets the default transformer function for any
// extension that doesn't have a specific transformer registered
func SetDefaultTransformer(f Transformer) {
	defaultTransformer = f
}

// Generic Transform function
func Transform(extensions map[string]interface{}) map[string]interface{} {
	if extensions == nil {
		return nil
	}

	out := make(map[string]interface{}, len(extensions))

	for originalKey, payload := range extensions {
		transformedName := originalKey // fallback
		
		// Resolve registered extensions explicitly
		if extDescs, err := proto.ExtensionDescs((*descriptor.MethodOptions)(nil)); err == nil {
			for _, extDesc := range extDescs {
				protoKey := fmt.Sprintf(".google.protobuf.MethodOptions.%s", extDesc.Name[strings.LastIndex(extDesc.Name, ".")+1:])
				if protoKey == originalKey {
					transformedName = extDesc.Name
					break
				}
			}
		}

		// Apply the transformer if one exists
		transform, ok := transformers[transformedName]
		if ok {
			out[transformedName] = transform(payload)
		} else if defaultTransformer != nil {
			out[transformedName] = defaultTransformer(payload)
		} else {
			out[transformedName] = payload
		}
	}

	return out
}
