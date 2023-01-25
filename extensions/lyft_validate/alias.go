package extensions

import evp "github.com/Raiden1974/protoc-gen-doc/extensions/envoyproxy_validate"

// ValidateRule represents a single validator rule from the (validate.rules) method option extension.
type ValidateRule = evp.ValidateRule

// ValidateExtension contains the rules set by the (validate.rules) method option extension.
type ValidateExtension = evp.ValidateExtension
