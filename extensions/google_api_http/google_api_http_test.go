package extensions_test

import (
	"net/http"
	"testing"

	"github.com/pseudomuto/protoc-gen-doc/extensions"
	. "github.com/pseudomuto/protoc-gen-doc/extensions/google_api_http"
	"github.com/stretchr/testify/suite"
	"google.golang.org/genproto/googleapis/api/annotations"
)

var rule *annotations.HttpRule

type HTTPRuleTest struct {
	suite.Suite
}

func TestHTTPRule(t *testing.T) {
	suite.Run(t, new(HTTPRuleTest))
}

func (assert *HTTPRuleTest) SetupSuite() {
	rule = &annotations.HttpRule{
		Pattern: &annotations.HttpRule_Get{Get: "/api/v1/method"},
		AdditionalBindings: []*annotations.HttpRule{
			{Pattern: &annotations.HttpRule_Put{Put: "/api/v1/method_alt"}, Body: "*"},
			{Pattern: &annotations.HttpRule_Post{Post: "/api/v1/method_alt"}, Body: "*"},
			{Pattern: &annotations.HttpRule_Delete{Delete: "/api/v1/method_alt"}},
			{Pattern: &annotations.HttpRule_Patch{Patch: "/api/v1/method_alt"}, Body: "*"},
			{Pattern: &annotations.HttpRule_Custom{Custom: &annotations.CustomHttpPattern{
				Kind: http.MethodOptions,
				Path: "/api/v1/method_alt",
			}}},
		},
	}
}

func (assert *HTTPRuleTest) TestTransform() {
	transformed := extensions.Transform(map[string]interface{}{
		"google.api.http": rule,
	})
	assert.NotEmpty(transformed)
	if assert.Contains(transformed, "google.api.http") {
		rules := transformed["google.api.http"].(HTTPExtension).Rules
		assert.Equal(rules, []HTTPRule{
			{Method: http.MethodGet, Pattern: "/api/v1/method"},
			{Method: http.MethodPut, Pattern: "/api/v1/method_alt", Body: "*"},
			{Method: http.MethodPost, Pattern: "/api/v1/method_alt", Body: "*"},
			{Method: http.MethodDelete, Pattern: "/api/v1/method_alt"},
			{Method: http.MethodPatch, Pattern: "/api/v1/method_alt", Body: "*"},
			{Method: http.MethodOptions, Pattern: "/api/v1/method_alt"},
		})
	}
}
