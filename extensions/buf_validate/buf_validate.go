package extensions

import (
	"encoding/json"
	"reflect"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
)

// ValidateRule represents a single validator rule from the (validate.rules) method option extension.
type ValidateRule struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// ValidateExtension contains the rules set by the (validate.rules) method option extension.
type ValidateExtension struct {
	*validate.FieldConstraints
	rules []ValidateRule // memoized so that we don't have to use reflection more than we need.
}

// MarshalJSON implements the json.Marshaler interface.
func (v ValidateExtension) MarshalJSON() ([]byte, error) { return json.Marshal(v.Rules()) }

// Rules returns the set of rules for this extension.
func (v ValidateExtension) Rules() []ValidateRule {
	if v.FieldConstraints == nil {
		return nil
	}
	if v.rules == nil {
		v.rules = flattenRules("", reflect.ValueOf(v.FieldConstraints))
	}
	return v.rules
}

func flattenRules(prefix string, vv reflect.Value) (rules []ValidateRule) {
	vv = reflect.Indirect(vv)
	vt := vv.Type()
	switch vt.Kind() {
	case reflect.Struct:
	nextField:
		for i := 0; i < vt.NumField(); i++ {
			f := vt.Field(i)
			ft := f.Type
			fv := vv.Field(i)

			var wasIndirect bool
			for ft.Kind() == reflect.Interface || ft.Kind() == reflect.Ptr {
				if fv.IsNil() {
					continue nextField
				}
				wasIndirect = true
				fv = fv.Elem()
				ft = fv.Type()
			}

			if !wasIndirect && fv.IsZero() {
				continue nextField
			}

			name := prefix
			if tag, ok := f.Tag.Lookup("protobuf"); ok {
				for _, opt := range strings.Split(tag, ",") {
					if strings.HasPrefix(opt, "name=") {
						if name != "" && !strings.HasSuffix(name, ".") {
							name += "."
						}
						name += strings.TrimPrefix(opt, "name=")
						break
					}
				}
			} else if _, ok := f.Tag.Lookup("protobuf_oneof"); !ok {
				continue nextField
			}
			rules = append(rules, flattenRules(name, fv)...)
		}
	case reflect.Slice:
		if vv.Len() == 0 {
			return nil
		}
		fallthrough
	default:
		rules = append(rules, ValidateRule{Name: prefix, Value: vv.Interface()})
	}
	return rules
}

func init() {
	extensions.SetTransformer("buf.validate.field", func(payload interface{}) interface{} {
		rules, ok := payload.(*validate.FieldConstraints)
		if !ok {
			return nil
		}
		return ValidateExtension{FieldConstraints: rules}
	})
}
