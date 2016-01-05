# Protocol Documentation
<a name="top"/>

## Table of Contents
* [Booking.proto](#Booking.proto)
 * [Booking](#com.example.Booking)
 * [BookingStatus](#com.example.BookingStatus)
* [Customer.proto](#Customer.proto)
 * [Address](#com.example.Address)
 * [Customer](#com.example.Customer)
* [Vehicle.proto](#Vehicle.proto)
 * [Manufacturer](#com.example.Manufacturer)
 * [Model](#com.example.Model)
 * [Vehicle](#com.example.Vehicle)
 * [Vehicle.Category](#com.example.Vehicle.Category)
 * [Manufacturer.Category](#com.example.Manufacturer.Category)
 * [Manufacturer.country](#com.example.country)
* [Scalar Value Types](#scalar-value-types)

<a name="Booking.proto"/>
<p align="right"><a href="#top">Top</a></p>

## Booking.proto

<a name="com.example.Booking"/>
### Booking
Represents the booking of a vehicle.

Vehicles are some cool shit. But drive carefully!

| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| vehicle_id | [int32](#int32) | required | ID of booked vehicle. |
| customer_id | [int32](#int32) | required | Customer that booked the vehicle. |
| status | [BookingStatus](#com.example.BookingStatus) | required | Status of the booking. |
| confirmation_sent | [bool](#bool) | required | Has booking confirmation been sent? |
| payment_received | [bool](#bool) | required | Has payment been received? |

<a name="com.example.BookingStatus"/>
### BookingStatus
Represents the status of a vehicle booking.

| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) | required | Unique booking status ID. |
| description | [string](#string) | required | Booking status description. E.g. &quot;Active&quot;. |




<a name="Customer.proto"/>
<p align="right"><a href="#top">Top</a></p>

## Customer.proto

<a name="com.example.Address"/>
### Address
Represents a mail address.

| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address_line_1 | [string](#string) | required | First address line. |
| address_line_2 | [string](#string) | optional | Second address line. |
| address_line_3 | [string](#string) | optional | Second address line. |
| town | [string](#string) | required | Address town. |
| county | [string](#string) | optional | Address county, if applicable. |
| country | [string](#string) | required | Address country. |

<a name="com.example.Customer"/>
### Customer
Represents a customer.

| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) | required | Unique customer ID. |
| first_name | [string](#string) | required | Customer first name. |
| last_name | [string](#string) | required | Customer last name. |
| details | [string](#string) | optional | Customer details. |
| email_address | [string](#string) | optional | Customer e-mail address. |
| phone_number | [string](#string) | repeated | Customer phone numbers, primary first. |
| mail_addresses | [Address](#com.example.Address) | repeated | Customer mail addresses, primary first. |




<a name="Vehicle.proto"/>
<p align="right"><a href="#top">Top</a></p>

## Vehicle.proto

<a name="com.example.Manufacturer"/>
### Manufacturer
Represents a manufacturer of cars.

| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) | required | The unique manufacturer ID. |
| code | [string](#string) | required | A manufacturer code, e.g. &quot;DKL4P&quot;. |
| category | [Manufacturer.Category](#com.example.Manufacturer.Category) | required | Manufacturer category. |
| details | [string](#string) | optional | Manufacturer details (minimum orders et.c.). |

<a name="com.example.Model"/>
### Model
Represents a vehicle model.

| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) | required | The unique model ID. |
| model_code | [string](#string) | required | The car model code, e.g. &quot;PZ003&quot;. |
| model_name | [string](#string) | required | The car model name, e.g. &quot;Z3&quot;. |
| daily_hire_rate_dollars | [sint32](#sint32) | required | Dollars per day. |
| daily_hire_rate_cents | [sint32](#sint32) | required | Cents per day. |

<a name="com.example.Vehicle"/>
### Vehicle
Represents a vehicle that can be hired.

| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| series | string | Model | 100 | Vehicle model series. |


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) | required | Unique vehicle ID. |
| model | [Model](#com.example.Model) | required | Vehicle model. |
| reg_number | [string](#string) | required | Vehicle registration number. |
| mileage | [sint32](#sint32) | optional | Current vehicle mileage, if known. |
| category | [Vehicle.Category](#com.example.Vehicle.Category) | optional | Vehicle category. |
| daily_hire_rate_dollars | [sint32](#sint32) | optional | Dollars per day.Taken from model if unspecified. |
| daily_hire_rate_cents | [sint32](#sint32) | optional | Cents per day.Taken from model if unspecified. |

<a name="com.example.Vehicle.Category"/>
### Vehicle.Category
Represents a vehicle category. E.g. &quot;Sedan&quot; or &quot;Truck&quot;.

| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) | required | Category code. E.g. &quot;S&quot;. |
| description | [string](#string) | required | Category name. E.g. &quot;Sedan&quot;. |


<a name="com.example.Manufacturer.Category"/>
### Manufacturer.Category
Manufacturer category. A manufacturer may be either inhouse or external.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CATEGORY_INHOUSE | 0 | The manufacturer is inhouse. |
| CATEGORY_EXTERNAL | 1 | The manufacturer is external. |


<a name="com.example.country"/>
### Manufacturer.country
Manufacturer country.

country extends Manufacturer at number 100


<a name="scalar-value-types"/>
## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double"/> double |  | double | double | float |
| <a name="float"/> float |  | float | float | float |
| <a name="int32"/> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64"/> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32"/> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64"/> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32"/> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64"/> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32"/> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64"/> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32"/> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64"/> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool"/> bool |  | bool | boolean | boolean |
| <a name="string"/> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes"/> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
