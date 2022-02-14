# API Reference

# Table of Contents


- Services
    - [BookingService](#comexamplebookingservice)
  


- Messages
    - [Booking](#booking)
    - [BookingStatus](#bookingstatus)
    - [BookingStatusID](#bookingstatusid)
    - [EmptyBookingMessage](#emptybookingmessage)
  





- Messages
    - [Address](#address)
    - [Customer](#customer)
  





- Messages
    - [Manufacturer](#manufacturer)
    - [Model](#model)
    - [Vehicle](#vehicle)
    - [Vehicle.Category](#vehiclecategory)
  



- [Scalar Value Types](#scalar-value-types)



# BookingService {#comexamplebookingservice}
Service for handling vehicle bookings.

## BookVehicle

> **rpc** BookVehicle([Booking](#booking))
    [BookingStatus](#bookingstatus)

Used to book a vehicle. Pass in a Booking and a BookingStatus will be returned.
## BookingUpdates

> **rpc** BookingUpdates([BookingStatusID](#bookingstatusid))
    [BookingStatus](#bookingstatus)

Used to subscribe to updates of the BookingStatus.
 <!-- end methods -->
 <!-- end services -->

# Messages


## Booking {#booking}
Represents the booking of a vehicle.

Vehicles are some cool shit. But drive carefully!


| Field | Type | Description |
| ----- | ---- | ----------- |
| vehicle_id | [ int32](#int32) | ID of booked vehicle. |
| customer_id | [ int32](#int32) | Customer that booked the vehicle. |
| status | [ BookingStatus](#bookingstatus) | Status of the booking. |
| confirmation_sent | [ bool](#bool) | Has booking confirmation been sent? |
| payment_received | [ bool](#bool) | Has payment been received? |
| color_preference | [ string](#string) | Color preference of the customer. |
 <!-- end Fields -->
 <!-- end HasFields -->


## BookingStatus {#bookingstatus}
Represents the status of a vehicle booking.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int32](#int32) | Unique booking status ID. |
| description | [ string](#string) | Booking status description. E.g. "Active". |
 <!-- end Fields -->
 <!-- end HasFields -->


## BookingStatusID {#bookingstatusid}
Represents the booking status ID.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int32](#int32) | Unique booking status ID. |
 <!-- end Fields -->
 <!-- end HasFields -->


## EmptyBookingMessage {#emptybookingmessage}
An empty message for testing

 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## Address {#address}
Represents a mail address.


| Field | Type | Description |
| ----- | ---- | ----------- |
| address_line_1 | [required string](#string) | First address line. |
| address_line_2 | [optional string](#string) | Second address line. |
| address_line_3 | [optional string](#string) | Second address line. |
| town | [required string](#string) | Address town. |
| county | [optional string](#string) | Address county, if applicable. |
| country | [required string](#string) | Address country. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Customer {#customer}
Represents a customer.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [required int32](#int32) | Unique customer ID. |
| first_name | [required string](#string) | Customer first name. |
| last_name | [required string](#string) | Customer last name. |
| details | [optional string](#string) | Customer details. |
| email_address | [optional string](#string) | Customer e-mail address. |
| phone_number | [repeated string](#string) | Customer phone numbers, primary first. |
| mail_addresses | [repeated Address](#address) | Customer mail addresses, primary first. |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums
 <!-- end Enums -->


 <!-- end services -->

# Messages


## Manufacturer {#manufacturer}
Represents a manufacturer of cars.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [required int32](#int32) | The unique manufacturer ID. |
| code | [required string](#string) | A manufacturer code, e.g. "DKL4P". |
| details | [optional string](#string) | Manufacturer details (minimum orders et.c.). |
| category | [optional Manufacturer.Category](#manufacturercategory) | Manufacturer category. Default: CATEGORY_EXTERNAL |
 <!-- end Fields -->
 <!-- end HasFields -->


## Model {#model}
Represents a vehicle model.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [required string](#string) | The unique model ID. |
| model_code | [required string](#string) | The car model code, e.g. "PZ003". |
| model_name | [required string](#string) | The car model name, e.g. "Z3". |
| daily_hire_rate_dollars | [required sint32](#sint32) | Dollars per day. |
| daily_hire_rate_cents | [required sint32](#sint32) | Cents per day. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Vehicle {#vehicle}
Represents a vehicle that can be hired.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [required int32](#int32) | Unique vehicle ID. |
| model | [required Model](#model) | Vehicle model. |
| reg_number | [required string](#string) | Vehicle registration number. |
| mileage | [optional sint32](#sint32) | Current vehicle mileage, if known. |
| category | [optional Vehicle.Category](#vehiclecategory) | Vehicle category. |
| daily_hire_rate_dollars | [optional sint32](#sint32) | Dollars per day. Default: 50 |
| daily_hire_rate_cents | [optional sint32](#sint32) | Cents per day. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Vehicle.Category {#vehiclecategory}
Represents a vehicle category. E.g. "Sedan" or "Truck".


| Field | Type | Description |
| ----- | ---- | ----------- |
| code | [required string](#string) | Category code. E.g. "S". |
| description | [required string](#string) | Category name. E.g. "Sedan". |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums


## Manufacturer.Category {#manufacturercategory}
Manufacturer category. A manufacturer may be either inhouse or external.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CATEGORY_INHOUSE | 0 | The manufacturer is inhouse. |
| CATEGORY_EXTERNAL | 1 | The manufacturer is external. |


 <!-- end Enums -->
 <!-- end Files -->

# Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <div><h4 id="double" /></div><a name="double" /> double |  | double | double | float |
| <div><h4 id="float" /></div><a name="float" /> float |  | float | float | float |
| <div><h4 id="int32" /></div><a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <div><h4 id="int64" /></div><a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <div><h4 id="uint32" /></div><a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <div><h4 id="uint64" /></div><a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <div><h4 id="sint32" /></div><a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <div><h4 id="sint64" /></div><a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <div><h4 id="fixed32" /></div><a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <div><h4 id="fixed64" /></div><a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <div><h4 id="sfixed32" /></div><a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <div><h4 id="sfixed64" /></div><a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <div><h4 id="bool" /></div><a name="bool" /> bool |  | bool | boolean | boolean |
| <div><h4 id="string" /></div><a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <div><h4 id="bytes" /></div><a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

