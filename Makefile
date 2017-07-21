.PHONY: bench test build

generate:
	@go generate

test: generate
	@go test -cover $(shell go list ./... | grep -v -E 'build|cmd|test|tools|vendor')

bench:
	@go test -bench=.

build: generate
	@go build ./cmd/...

examples: build
	@rm -f examples/doc/*
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=docbook,example.docbook examples/proto/*.proto
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=html,example.html examples/proto/*.proto
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=json,example.json examples/proto/*.proto
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=markdown,example.md examples/proto/*.proto
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=examples/templates/asciidoc.tmpl,example.txt examples/proto/*.proto
