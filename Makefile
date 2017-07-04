.PHONY: bench test

generate:
	@go generate

test: generate
	@go test -cover $(shell go list ./... | grep -v -E 'build|test|tools|vendor')

bench:
	@go test -bench=.
