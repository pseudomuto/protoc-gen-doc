DOCBOOK_XSL=/usr/share/xml/docbook/xsl-stylesheets-1.78.1/fo/docbook.xsl

export PATH := ..:$(PATH)

all: doc/example.html doc/example.pdf

doc/example.html: *.proto ../protoc-gen-doc
	protoc --doc_out=html,example.html:doc *.proto

doc/example.docbook: *.proto ../protoc-gen-doc
	protoc --doc_out=docbook,example.docbook:doc *.proto

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
