.PHONY: bench setup test build dist docker

EXAMPLE_DIR=$(shell pwd)/examples
DOCS_DIR=$(EXAMPLE_DIR)/doc
PROTOS_DIR=$(EXAMPLE_DIR)/proto

setup:
	$(info Synching dev tools and dependencies...)
	@if test -z $(which retool); then go get github.com/twitchtv/retool; fi
	@retool sync
	@retool do dep ensure

generate:
	@go generate

lint:
	@golint -set_exit_status ./build/... && \
		golint -set_exit_status ./cmd/... && \
		golint -set_exit_status ./parser/... && \
		golint -set_exit_status ./test/... && \
		golint -set_exit_status .

test: generate
	@go test -cover $(shell go list ./... | grep -v -E 'build|cmd|test|tools|vendor')

bench:
	@go test -bench=.

build: setup generate
	@go build ./cmd/...

examples: build
	@rm -f examples/doc/*
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=docbook,example.docbook:Ignore* examples/proto/*.proto
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=html,example.html:Ignore* examples/proto/*.proto
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=json,example.json:Ignore* examples/proto/*.proto
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=markdown,example.md:Ignore* examples/proto/*.proto
	@protoc --plugin=protoc-gen-doc --doc_out=examples/doc --doc_opt=examples/templates/asciidoc.tmpl,example.txt:Ignore* examples/proto/*.proto

dist:
	@script/dist.sh

docker:
	@script/push_to_docker.sh

docker_test: docker
	@rm -f examples/doc/*
	@docker run --rm -v $(DOCS_DIR):/out:rw -v $(PROTOS_DIR):/protos:ro pseudomuto/protoc-gen-doc --doc_opt=docbook,example.docbook:Ignore*
	@docker run --rm -v $(DOCS_DIR):/out:rw -v $(PROTOS_DIR):/protos:ro pseudomuto/protoc-gen-doc --doc_opt=html,example.html:Ignore*
	@docker run --rm -v $(DOCS_DIR):/out:rw -v $(PROTOS_DIR):/protos:ro pseudomuto/protoc-gen-doc --doc_opt=json,example.json:Ignore*
	@docker run --rm -v $(DOCS_DIR):/out:rw -v $(PROTOS_DIR):/protos:ro pseudomuto/protoc-gen-doc --doc_opt=markdown,example.md:Ignore*
	@docker run --rm \
		-v $(DOCS_DIR):/out:rw \
		-v $(PROTOS_DIR):/protos:ro \
		-v $(EXAMPLE_DIR)/templates:/templates:ro \
		pseudomuto/protoc-gen-doc --doc_opt=/templates/asciidoc.tmpl,example.txt:Ignore*
