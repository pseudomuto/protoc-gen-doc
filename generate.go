package gendoc

//go:generate go build ./build/cmd/gen_fixtures
//go:generate protoc --plugin=protoc-gen-doc=./gen_fixtures -Itest/fixtures --doc_out=./test test/fixtures/Booking.proto test/fixtures/Vehicle.proto
//go:generate rm gen_fixtures

//go:generate go run build/cmd/resources/main.go -in resources -out resources.go -pkg gendoc
