# protoc-gen-doc

[![Travis Build Status][travis-svg]][travis-ci]

This is a documentation generator plugin for the Google Protocol Buffers compiler (`protoc`). The plugin can generate
HTML, JSON, DocBook and Markdown documentation from comments in your `.proto` files.

It supports proto2 and proto3, and can handle having both in the same context (see [examples](examples/) for proof).

## Installation

`go get -u github.com/pseudomuto/protoc-gen-doc`

## Writing Documentation

Messages, Fields, Services (and their methods), Enums (and their values), Extensions, and Files can be documented.
Generally speaking, comments come in 2 forms: leading and trailing.

**Leading comments**

Leading comments can be used everywhere.

```protobuf
/**
 * This is a leading comment for a message
*/
message SomeMessage {
  // this is another leading comment
  string value = 1;
}
```

**Trailing comments**

Fields, Service Methods, Enum Values and Extensions support trailing comments.

```protobuf
enum MyEnum {
  DEFAULT = 0; // the default value
  OTHER   = 1; // the other value
}
```

Check out the [example protos](examples/proto) to see all the options.

## Invoking the Plugin

The plugin is invoked by passing the `--doc_out`, and `--doc_opt` options to the `protoc` compiler. The option has the
following format:

    --doc_opt=<FORMAT>|<TEMPLATE_FILENAME>,<OUT_FILENAME>

The format may be one of the built-in ones ( `docbook`, `html`, `markdown` or `json`)
or the name of a file containing a custom [Go template][gotemplate].

### Simple Usage

For example, to generate HTML documentation for all `.proto` files in the `proto` directory into `doc/index.html`, type:

    protoc --doc_out=./doc --doc_opt=html,index.html proto/*.proto

The plugin executable must be in `PATH` for this to work. 

### With a Custom Build

Alternatively, you can specify a pre-built/not in `PATH` binary using the `--plugin` option.

    protoc \
      --plugin=protoc-gen-doc=./protoc-gen-doc \
      --doc_out=./doc \
      --doc_opt=html,index.html \
      proto/*.proto

### With a Custom Template

If you'd like to use your own template, simply use the path to the template file rather than the type.

    protoc --doc_out=./doc --doc_opt=/path/to/template.tmpl,index.txt proto/*.proto

For information about the available template arguments and functions, see [Custom Templates][custom]. If you just want
to customize the look of the HTML output, put your CSS in `stylesheet.css` next to the output file and it will be picked
up.

## Output Example

With the input `.proto` files

* [Booking.proto](examples/proto/Booking.proto)
* [Customer.proto](examples/proto/Customer.proto)
* [Vehicle.proto](examples/proto/Vehicle.proto)

the plugin gives the output

* [Markdown](examples/doc/example.md)
* [HTML][html_preview]
* [DocBook](examples/doc/example.docbook)
* [JSON](examples/doc/example.json)

Check out the `examples` task in the [Makefile](Makefile) to see how these were generated.

[gotemplate]:
    https//golang.org/pkg/text/template/
    "Template - The Go Programming Language"
[custom]:
    https://github.com/pseudomuto/protoc-gen-doc/wiki/Custom-Templates
    "Custom templates instructions"
[html_preview]:
    https://rawgit.com/pseudomuto/protoc-gen-doc/master/examples/doc/example.html
    "HTML Example Output"
[travis-svg]:
    https://travis-ci.org/pseudomuto/protoc-gen-doc.svg?branch=master
    "Travis CI build status SVG"
[travis-ci]:
    https://travis-ci.org/pseudomuto/protoc-gen-doc
    "protoc-gen-doc at Travis CI"
