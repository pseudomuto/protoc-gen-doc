all: html/example.html

html/example.html: *.proto
	protoc --plugin=../protoc-gen-doc --doc_out=html,example.html:html *.proto
