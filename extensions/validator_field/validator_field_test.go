package extensions_test

import (
	"github.com/golang/protobuf/proto"
	validator "github.com/mwitkow/go-proto-validators"
	"github.com/stretchr/testify/suite"

	"testing"

	"github.com/pseudomuto/protoc-gen-doc/extensions"
	validator_field "github.com/pseudomuto/protoc-gen-doc/extensions/validator_field"
)

var fieldValidator *validator.FieldValidator

type ValidatorTest struct {
	suite.Suite
}

func TestValidator(t *testing.T) {
	suite.Run(t, new(ValidatorTest))
}

func (assert *ValidatorTest) SetupSuite() {
	fieldValidator = &validator.FieldValidator{
		StringNotEmpty: proto.Bool(true),
	}
}

func (assert *ValidatorTest) TestTransform() {
	transformed := extensions.Transform(map[string]interface{}{
		"validator.field": fieldValidator,
	})
	assert.NotEmpty(transformed)
	if assert.Contains(transformed, "validator.field") {
		rules := transformed["validator.field"].(validator_field.ValidatorExtension).Rules()
		assert.Equal(rules, []validator_field.ValidatorRule{
			{Name: "string_not_empty", Value: true},
		})
	}
}
