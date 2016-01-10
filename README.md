# protoc-gen-doc

[![Build Status][travis-svg]][travis-ci]

This is a documentation generator plugin for the Google Protocol Buffers compiler
(`protoc`). The plugin can generate HTML, DocBook or Markdown documentation from
comments in your `.proto` files.

## Installation

* [Debian][obs]
* [Ubuntu][obs]
* [openSUSE][obs]
* [Fedora][obs]
* [Arch Linux][obs]
* [CentOS 7][centos] (x86_64 only)
* [Windows][releases]
* Or [Build From Source](BUILDING.md)

## Writing Documentation

Use `/** */` or `///` comments to document your files. Comments for enumerations and
messages go before the message/enumeration definition. Comments for fields, enum
values and extensions can go either before or after the field/value/extension
definition. If a documentation comment begins with `@exclude`, the message, enum,
enum value, field or extension will be excluded from the
generated documentation.

## Invoking the Plugin

The plugin is invoked by passing the `--doc_out` option to the `protoc` compiler. The
option has the following format:

    --doc_out=docbook|html|markdown|<TEMPLATE_FILENAME>,<OUT_FILENAME>:<OUT_DIR>

The format may be one of the built-in ones ( `docbook`, `html` or `markdown`) or the
name of a file containing a custom [Mustache][mustache] template. For example, to
generate HTML documentation for all `.proto` files in the `proto` directory into
`doc/index.html`, type:

    protoc --doc_out=html,index.html:doc proto/*.proto

The plugin executable must be in `PATH` or specified explicitly using the `--plugin`
option in order for `protoc` to find it. If you need support for a custom output
format, see [Custom Templates][custom]. If you just want to customize the look of the
HTML output, put your CSS in `stylesheet.css` next to the output file and it will be
picked up.

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

Look in [examples/Makefile](examples/Makefile) to see how these outputs were built.


[epel]:
    https://fedoraproject.org/wiki/EPEL
    "EPEL repository"
[mustache]:
    http://mustache.github.io/
    "Mustache - Logic-less templates"
[custom]:
    https://github.com/estan/protoc-gen-doc/wiki/Custom-Templates
    "Custom templates instructions"
[fop]:
    http://xmlgraphics.apache.org/fop/
    "Apache™ FOP (Formatting Objects Processor)"
[html_preview]:
    https://rawgit.com/estan/protoc-gen-doc/master/examples/doc/example.html
    "HTML Example Output"
[obs]:
    http://tinyurl.com/protoc-gen-doc-packages
    "Packages at Open Build Service"
[releases]:
    https://github.com/estan/protoc-gen-doc/releases
    "Releases for download"
[centos]:
    http://estan.github.io/protoc-gen-doc/
    "CentOS 7 repository"
[travis-svg]:
    https://travis-ci.org/estan/protoc-gen-doc.svg?branch=master
    "Travis CI build status SVG"
[travis-ci]:
    https://travis-ci.org/estan/protoc-gen-doc
    "protoc-gen-doc at Travis CI"
