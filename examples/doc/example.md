# Protocol Documentation
<a name="top"/>

## Table of Contents
* [Booking.proto](#Booking.proto)
 * [Booking](#com.example.Booking)
 * [BookingStatus](#com.example.BookingStatus)
 * [BookingService](#com.example.BookingService)
* [Customer.proto](#Customer.proto)
 * [Address](#com.example.Address)
 * [Customer](#com.example.Customer)
* [Vehicle.proto](#Vehicle.proto)
 * [Manufacturer](#com.example.Manufacturer)
 * [Model](#com.example.Model)
 * [Vehicle](#com.example.Vehicle)
 * [Vehicle.Category](#com.example.Vehicle.Category)
 * [Manufacturer.Category](#com.example.Manufacturer.Category)
 * [File-level Extensions](#Vehicle.proto-extensions)
* [Scalar Value Types](#scalar-value-types)

<a name="Booking.proto"/>
<p align="right"><a href="#top">Top</a></p>

## Booking.proto

Booking related messages.

This file is really just an example. The data model is completely
fictional.

Author: Elvis Stansvik

<a name="com.example.Booking"/>
### Booking
Represents the booking of a vehicle.

Vehicles are some cool shit. But drive carefully!

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

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) | required | Unique booking status ID. |
| description | [string](#string) | required | Booking status description. E.g. "Active". |





<a name="com.example.BookingService"/>
### BookingService
Service for handling vehicle bookings.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| BookVehicle | [Booking](#com.example.Booking) | [BookingStatus](#com.example.BookingStatus) | Used to book a vehicle. Pass in a Booking and a BookingStatus will be returned. |


<a name="Customer.proto"/>
<p align="right"><a href="#top">Top</a></p>

## Customer.proto

This file has messages for describing a customer.

Author: Elvis Stansvik

<a name="com.example.Address"/>
### Address
Represents a mail address.

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

Messages describing manufacturers / vehicles.

<a name="com.example.Manufacturer"/>
### Manufacturer
Represents a manufacturer of cars.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) | required | The unique manufacturer ID. |
| code | [string](#string) | required | A manufacturer code, e.g. "DKL4P". |
| details | [string](#string) | optional | Manufacturer details (minimum orders et.c.). |
| category | [Manufacturer.Category](#com.example.Manufacturer.Category) | optional | Manufacturer category. Default: CATEGORY_EXTERNAL |


<a name="com.example.Model"/>
### Model
Represents a vehicle model.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) | required | The unique model ID. |
| model_code | [string](#string) | required | The car model code, e.g. "PZ003". |
| model_name | [string](#string) | required | The car model name, e.g. "Z3". |
| daily_hire_rate_dollars | [sint32](#sint32) | required | Dollars per day. |
| daily_hire_rate_cents | [sint32](#sint32) | required | Cents per day. |


<a name="com.example.Vehicle"/>
### Vehicle
Represents a vehicle that can be hired.

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) | required | Unique vehicle ID. |
| model | [Model](#com.example.Model) | required | Vehicle model. |
| reg_number | [string](#string) | required | Vehicle registration number. |
| mileage | [sint32](#sint32) | optional | Current vehicle mileage, if known. |
| category | [Vehicle.Category](#com.example.Vehicle.Category) | optional | Vehicle category. |
| daily_hire_rate_dollars | [sint32](#sint32) | optional | Dollars per day. Default: 50 |
| daily_hire_rate_cents | [sint32](#sint32) | optional | Cents per day. |

| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| series | string | Model | 100 | Vehicle model series. |

<a name="com.example.Vehicle.Category"/>
### Vehicle.Category
Represents a vehicle category. E.g. "Sedan" or "Truck".

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) | required | Category code. E.g. "S". |
| description | [string](#string) | required | Category name. E.g. "Sedan". |



<a name="com.example.Manufacturer.Category"/>
### Manufacturer.Category
Manufacturer category. A manufacturer may be either inhouse or external.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CATEGORY_INHOUSE | 0 | The manufacturer is inhouse. |
| CATEGORY_EXTERNAL | 1 | The manufacturer is external. |


<a name="Vehicle.proto-extensions"/>
### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| country | string | Manufacturer | 100 | Manufacturer country. Default: &quot;China&quot; |



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
