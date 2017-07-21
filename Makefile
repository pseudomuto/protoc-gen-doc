.PHONY: bench test build

generate:
	@go generate

test: generate
	@go test -cover $(shell go list ./... | grep -v -E 'build|cmd|test|tools|vendor')

bench:
	@go test -bench=.

build:
	@go build ./cmd/...
