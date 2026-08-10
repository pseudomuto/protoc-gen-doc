package extensions

import (
	"github.com/pseudomuto/protoc-gen-doc/extensions"
	"google.golang.org/genproto/googleapis/api/annotations"
)

type FieldBehaviorExtension struct {
	Options []string `json:"options"`
}

func init() {
	extensions.SetTransformer("google.api.field_behavior", func(payload interface{}) interface{} {
		fb, ok := payload.([]annotations.FieldBehavior)
		if !ok {
			return nil
		}

		if len(fb) == 0 {
			return nil
		}

		fbs := make([]string, len(fb))
		for i, behavior := range fb {
			fbs[i] = behavior.String()
		}

		return FieldBehaviorExtension{Options: fbs}
	})
}
