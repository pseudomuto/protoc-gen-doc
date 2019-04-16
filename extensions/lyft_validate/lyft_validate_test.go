package extensions_test

import (
	"testing"

	"github.com/envoyproxy/protoc-gen-validate/validate"
	"github.com/golang/protobuf/proto"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
	. "github.com/pseudomuto/protoc-gen-doc/extensions/lyft_validate"
	"github.com/stretchr/testify/suite"
)

var fieldRules *validate.FieldRules

type ValidateTest struct {
	suite.Suite
}

func TestValidate(t *testing.T) {
	suite.Run(t, new(ValidateTest))
}

func (assert *ValidateTest) SetupSuite() {
	fieldRules = &validate.FieldRules{
		Type: &validate.FieldRules_String_{
			String_: &validate.StringRules{
				MinLen: proto.Uint64(1),
			},
		},
	}
}

func (assert *ValidateTest) TestTransform() {
	transformed := extensions.Transform(map[string]interface{}{
		"validate.rules": fieldRules,
	})
	assert.NotEmpty(transformed)
	if assert.Contains(transformed, "validate.rules") {
		rules := transformed["validate.rules"].(ValidateExtension).Rules()
		assert.Equal(rules, []ValidateRule{
			{Name: "string.min_len", Value: uint64(1)},
		})
	}
}
