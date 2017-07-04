package protoc_gen_doc

//go:generate go build ./build/cmd/gen_fixtures
//go:generate protoc --plugin=protoc-gen-doc=./gen_fixtures --doc_out=. fixtures/Booking.proto fixtures/Vehicle.proto
//go:generate rm gen_fixtures

//go:generate go run build/cmd/resources/main.go -in templates -out resources.go -pkg protoc_gen_doc
