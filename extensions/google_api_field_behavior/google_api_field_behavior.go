package extensions

import (
	"fmt"
	"github.com/pseudomuto/protoc-gen-doc/extensions"
)

func init() {
	extensions.SetTransformer("google.api.field_behavior", func(payload interface{}) interface{} {
		fmt.Printf("debug field_behavior: %+v\n", payload)
		return payload
	})
}
