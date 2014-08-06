DOCBOOK_XSL="/usr/share/xml/docbook/xsl-stylesheets-1.78.1/fo/docbook.xsl"

all: html/example.html docbook/example.pdf

html/example.html: *.proto ../protoc-gen-doc
	protoc --plugin=../protoc-gen-doc --doc_out=html,example.html:html *.proto

docbook/example.docbook: *.proto ../protoc-gen-doc
	protoc --plugin=../protoc-gen-doc --doc_out=docbook,example.docbook:docbook *.proto

docbook/example.pdf: docbook/example.docbook
	xsltproc -o docbook/example.fo \
		--stringparam use.extensions 0 \
		--stringparam fop1.extensions 1 \
		--stringparam paper.type A4 \
		--stringparam page.orientation landscape \
		$(DOCBOOK_XSL) \
		docbook/example.docbook
	fop -fo docbook/example.fo -pdf docbook/example.pdf

clean:
	-rm html/* docbook/*
