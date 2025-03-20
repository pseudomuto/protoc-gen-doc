// Package generic provides a mechanism to register and transform any extension
// without knowing its specifics in advance
package generic

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
	"regexp"
)

// ExtensionPattern defines a pattern for registering multiple extensions at once
type ExtensionPattern struct {
	Name   string
	Number int32
	Type   proto.Message
}

// Initialize extension patterns - can be extended for more types
func init() {
	// Register a generic transformer for all extensions
	extensions.SetDefaultTransformer(identityTransformer)
}


// formatFieldNumber formats a field number for use in extension names
func formatFieldNumber(num int32) string {
	return fmt.Sprintf("field_%d", num)
}

// identityTransformer returns the extension value as-is
func identityTransformer(payload interface{}) interface{} {
	return payload
}

// ExtractOptionName extracts a cleaner name from a full extension name
func ExtractOptionName(fullName string) string {
	// Match the base name pattern (everything before the last dot and numbers)
	re := regexp.MustCompile(`(.*?)\.[\w_]+\d+$`)
	matches := re.FindStringSubmatch(fullName)
	if len(matches) > 1 {
		return matches[1]
	}
	return fullName
}
