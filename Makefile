.PHONY: generate test

generate:
	@go generate

test: generate
	@go test -cover $(shell go list ./... | grep -v -E 'test|tools|vendor')
