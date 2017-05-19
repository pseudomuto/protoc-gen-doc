# protoc-gen-doc

[![Travis Build Status][travis-svg]][travis-ci]

This is a documentation generator plugin for the Google Protocol Buffers compiler
(`protoc`). The plugin can generate HTML, DocBook and Markdown documentation from
comments in your `.proto` files, as well as a raw JSON representation.

## Installation

`go get -u github.com/pseudomuto/protoc-gen-doc`

## Writing Documentation

Use `/** */` or `///` comments to document your files. Comments for files go at the
very top of the the file. Comments for enumerations, messages and services go before
the message, enumeration or service definition. Comments for fields, enum values,
extensions and service methods can go either before or after the definition. If a
documentation comment begins with `@exclude`, the corresponding item will be excluded
from the generated documentation.

## Invoking the Plugin

The plugin is invoked by passing the `--doc_out` option to the `protoc` compiler. The
option has the following format:

    --doc_out=docbook|html|markdown|json|<TEMPLATE_FILENAME>,<OUT_FILENAME>[,no-exclude]:<OUT_DIR>

The format may be one of the built-in ones ( `docbook`, `html`, `markdown` or `json`)
or the name of a file containing a custom [Mustache][mustache] template. For example,
to generate HTML documentation for all `.proto` files in the `proto` directory into
`doc/index.html`, type:

    protoc --doc_out=html,index.html:doc proto/*.proto

The plugin executable must be in `PATH` or specified explicitly using the `--plugin`
option in order for `protoc` to find it. If you need support for a custom output
format, see [Custom Templates][custom]. If you just want to customize the look of the
HTML output, put your CSS in `stylesheet.css` next to the output file and it will be
picked up. If the optional `no-exclude` flag is given, all `@exclude` directives are
ignored.

## Output Example

With the input `.proto` files

* [Booking.proto](examples/proto/Booking.proto)
* [Customer.proto](examples/proto/Customer.proto)
* [Vehicle.proto](examples/proto/Vehicle.proto)

the plugin gives the output

* [Markdown](examples/doc/example.md)
* [HTML][html_preview]
* [DocBook](examples/doc/example.docbook)
* [PDF](examples/doc/example.pdf?raw=true) (Using [Apache FOP][fop])
* [JSON](examples/doc/example.json)

Look in [examples/Makefile](examples/Makefile) to see how these outputs were built.

[epel]:
    https://fedoraproject.org/wiki/EPEL
    "EPEL repository"
[mustache]:
    http://mustache.github.io/
    "Mustache - Logic-less templates"
[custom]:
    https://github.com/pseudomuto/protoc-gen-doc/wiki/Custom-Templates
    "Custom templates instructions"
[fop]:
    http://xmlgraphics.apache.org/fop/
    "Apache™ FOP (Formatting Objects Processor)"
[html_preview]:
    https://rawgit.com/pseudomuto/protoc-gen-doc/master/examples/doc/example.html
    "HTML Example Output"
[obs]:
    http://tinyurl.com/protoc-gen-doc-packages
    "Packages at Open Build Service"
[releases]:
    https://github.com/pseudomuto/protoc-gen-doc/releases
    "Releases for download"
[centos]:
    http://estan.github.io/protoc-gen-doc/
    "CentOS 7 repository"
[travis-svg]:
    https://travis-ci.org/pseudomuto/protoc-gen-doc.svg?branch=master
    "Travis CI build status SVG"
[travis-ci]:
    https://travis-ci.org/pseudomuto/protoc-gen-doc
    "protoc-gen-doc at Travis CI"
