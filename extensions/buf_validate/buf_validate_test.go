package extensions_test

import (
	"testing"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/golang/protobuf/proto"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
	. "github.com/pseudomuto/protoc-gen-doc/extensions/buf_validate"
	"github.com/stretchr/testify/require"
)

func TestTransform(t *testing.T) {
	fieldRules := &validate.FieldConstraints{
		Type: &validate.FieldConstraints_String_{
			String_: &validate.StringRules{

				MinLen: proto.Uint64(1),
				NotIn:  []string{"invalid"},
			},
		},
	}

	transformed := extensions.Transform(map[string]interface{}{"buf.validate.field": fieldRules})
	require.NotEmpty(t, transformed)

	rules := transformed["buf.validate.field"].(ValidateExtension).Rules()
	require.Equal(t, []ValidateRule{
		{Name: "string.min_len", Value: uint64(1)},
		{Name: "string.not_in", Value: []string{"invalid"}},
	}, rules)
}
