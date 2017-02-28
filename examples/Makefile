DOCBOOK_XSL?=/usr/share/xml/docbook/xsl-stylesheets-1.79.1/fo/docbook.xsl
export PATH := ..:$(PATH)

all: doc/example.md doc/example.html doc/example.docbook doc/example.pdf

doc/example.md: proto/*.proto ../protoc-gen-doc
	protoc --doc_out=markdown,example.md:doc proto/*.proto

doc/example.html: proto/*.proto ../protoc-gen-doc
	protoc --doc_out=html,example.html:doc proto/*.proto

doc/example.docbook: proto/*.proto ../protoc-gen-doc
	protoc --doc_out=docbook,example.docbook:doc proto/*.proto

doc/example.pdf: doc/example.docbook
	fop -xml doc/example.docbook \
		-xsl $(DOCBOOK_XSL) \
		-param use.extensions 0 \
		-param fop1.extensions 1 \
		-param paper.type A4 \
		-param page.orientation landscape \
		-pdf doc/example.pdf

clean:
	-rm doc/*
