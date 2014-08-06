DOCBOOK_XSL=/usr/share/xml/docbook/xsl-stylesheets-1.78.1/fo/docbook.xsl

all: doc/example.html doc/example.pdf

doc/example.html: *.proto ../protoc-gen-doc
	protoc --plugin=../protoc-gen-doc \
		--doc_out=html,example.html:doc *.proto

doc/example.docbook: *.proto ../protoc-gen-doc
	protoc --plugin=../protoc-gen-doc \
		--doc_out=docbook,example.docbook:doc *.proto

doc/example.pdf: doc/example.docbook
	xsltproc -o doc/example.fo \
		--stringparam use.extensions 0 \
		--stringparam fop1.extensions 1 \
		--stringparam paper.type A4 \
		--stringparam page.orientation landscape \
		$(DOCBOOK_XSL) \
		doc/example.docbook
	fop -fo doc/example.fo -pdf doc/example.pdf
	rm doc/example.fo

clean:
	-rm doc/*
