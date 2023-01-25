package extensions

import (
	"fmt"
	"github.com/Raiden1974/protoc-gen-doc/extensions"
)

func init() {
	extensions.SetTransformer("google.api.field_behavior", func(payload interface{}) interface{} {
		fmt.Printf("debug field_behavior: %+v\n", payload)
		return payload
	})
}
