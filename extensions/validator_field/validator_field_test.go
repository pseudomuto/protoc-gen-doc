package extensions_test

import (
	"testing"

	validator "github.com/mwitkow/go-proto-validators"
	"github.com/stretchr/testify/require"

	"github.com/pseudomuto/protoc-gen-doc/extensions"
	. "github.com/pseudomuto/protoc-gen-doc/extensions/validator_field"
)

func TestTransform(t *testing.T) {
	notEmpty := true
	fieldValidator := &validator.FieldValidator{
		StringNotEmpty: &notEmpty,
	}

	transformed := extensions.Transform(map[string]any{"validator.field": fieldValidator})
	require.NotEmpty(t, transformed)

	rules := transformed["validator.field"].(ValidatorExtension).Rules()
	require.Equal(t, rules, []ValidatorRule{
		{Name: "string_not_empty", Value: true},
	})
}
