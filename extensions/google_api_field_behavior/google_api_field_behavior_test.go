package extensions_test

import (
	"testing"

	"github.com/pseudomuto/protoc-gen-doc/extensions"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/api/annotations"
)

func TestTransform(t *testing.T) {
	behavior := []annotations.FieldBehavior{annotations.FieldBehavior_REQUIRED}

	transformed := extensions.Transform(map[string]interface{}{"google.api.field_behavior": behavior})
	require.NotEmpty(t, transformed)

	options := transformed["google.api.field_behavior"].(FieldBehaviorExtension).Options()
	require.Equal(t, options, []string{annotations.FieldBehavior_REQUIRED.String()})
}
