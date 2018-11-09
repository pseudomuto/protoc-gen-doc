package fixtures

//go:generate protoc --descriptor_set_out=fileset.pb --include_imports --include_source_info -I. -I../vendor Booking.proto Vehicle.proto
