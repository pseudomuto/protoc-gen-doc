# All the Types
<a name="top"/>

## Table of Contents
* [Scalar Value Types](#scalar-value-types)

<a name="scalar-value-types"/>
## Scalar Value Types

| .proto Type | Notes | C++ Type | C# Type | Go Type | Java Type | PHP Type | Python Type | Ruby Type |
| ----------- | ----- | -------- | ------- | --------| --------- | -------- | ----------- | --------- |
| <a name="double"/> double |  | double | double | float64 | double | float | float | Float |
| <a name="float"/> float |  | float | float | float32 | float | float | float | Float |
| <a name="int32"/> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int32 | int | integer | int | Bignum or Fixnum (as required) |
| <a name="int64"/> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int64 | long | integer/string | int/long | Bignum |
| <a name="uint32"/> uint32 | Uses variable-length encoding. | uint32 | uint | uint32 | int | integer | int/long | Bignum or Fixnum (as required) |
| <a name="uint64"/> uint64 | Uses variable-length encoding. | uint64 | ulong | uint64 | long | integer/string | int/long | Bignum or Fixnum (as required) |
| <a name="sint32"/> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int32 | int | integer | int | Bignum or Fixnum (as required) |
| <a name="sint64"/> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int64 | long | integer/string | int/long | Bignum |
| <a name="fixed32"/> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | uint | uint32 | int | integer | int | Bignum or Fixnum (as required) |
| <a name="fixed64"/> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | ulong | uint64 | long | integer/string | int/long | Bignum |
| <a name="sfixed32"/> sfixed32 | Always four bytes. | int32 | int | int32 | int | integer | int | Bignum or Fixnum (as required) |
| <a name="sfixed64"/> sfixed64 | Always eight bytes. | int64 | long | int64 | long | integer/string | int/long | Bignum |
| <a name="bool"/> bool |  | bool | bool | bool | boolean | boolean | boolean | TrueClass/FalseClass |
| <a name="string"/> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | string | string | String | string | str/unicode | String (UTF-8) |
| <a name="bytes"/> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | []byte | ByteString | string | str | String (ASCII-8BIT) |
