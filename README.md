[![Build Status](https://travis-ci.org/estan/protoc-gen-doc.svg?branch=master)](https://travis-ci.org/estan/protoc-gen-doc)

# Google Protocol Buffers<br>Documentation Generator

This is a documentation generator plugin for the Google Protocol
Buffers compiler (`protoc`). The plugin can generate HTML, DocBook
or Markdown documentation from comments in your `.proto` files.

## Building the Plugin

### Prerequisites
* Protocol Buffers library from Google
* QtCore from Qt 5

### Linux
At a terminal command prompt, run

    $ qmake
    $ make

in the top-level directory to build the plugin. This will produce
the plugin executable (`protoc-gen-doc`). There's no install step,
just copy the executable to where you want it.

### Windows
On windows, it's easiest to just use the [ZIP distribution][release_zip].
If you want to build yourself, start a Qt/MSVC command prompt,
load `vcvarsall.bat` and then run

    > set PROTOBUF_PREFIX=/path/to/protobuf-2.6.1
    > qmake
    > nmake

in the top-level directory to build the plugin. This will produce
the plugin executable (`release\protoc-gen-doc.exe`). You can
create a standalone ZIP distribution with `nmake zip`. MSVC is
currently the only supported compiler on Windows. Building with
MinGW should work, but the `zip` target is not available. I'll try
to fix this in the future.

### Mac OS X
At a Terminal prompt, run

    $ export PROTOBUF_PREFIX=/path/to/protobuf-2.6.1
    $ qmake
    $ make

in the top-level directory to build the plugin. This will produce
the plugin executable (`protoc-gen-doc`). There's no install step,
just copy the executable to where you want it.

Note that on Mac OS X, the protobuf library should be build with
with clang (`CC=clang` and `CXX=clang++`), or you'll get linker
errors.

## Invoking the Plugin

The plugin is invoked by passing the `--doc_out` option to the
`protoc` compiler. The option has the following format:

    --doc_out=docbook|html|markdown|<TEMPLATE_FILENAME>,<OUT_FILENAME>:<OUT_DIR>

The format may be one of the built-in ones ( `docbook`, `html` or
`markdown`) or the name of a file containing a custom
[Mustache][mustache] template. For example, to generate HTML
documentation for all `.proto` files in the `proto` directory into
`doc/index.html`, type:

    protoc --doc_out=html,index.html:doc proto/*.proto

The plugin executable must be in `PATH` or specified explicitly
using the `--plugin` option in order for `protoc` to find it. If
you need support for a custom output format, see the built-in
templates in [templates](templates) for how to write your
own. If you just want to customize the look of the HTML output,
put your CSS in `stylesheet.css` next to the output file and it
will be picked up.

## Writing Documentation

Use `/** */` or `///` comments to document your files. Comments
for enumerations and messages go before the message/enumeration
definition. Comments for fields or enum values can go either
before or after the field/value definition. If a documentation
comment begins with `@exclude`, the message, enum or field will
be excluded from the generated documentation.

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

Look in [examples/Makefile](examples/Makefile) to see how these
outputs were built.


[mustache]: http://mustache.github.io/ "Mustache - Logic-less templates"
[fop]: http://xmlgraphics.apache.org/fop/ "Apache™ FOP (Formatting Objects Processor)"
[html_preview]: https://rawgit.com/estan/protoc-gen-doc/master/examples/doc/example.html "HTML Example Output"
[release_zip]: https://github.com/estan/protoc-gen-doc/releases/download/v0.4/protoc-gen-doc-v0.4-win32.zip "Release 0.4 for Windows"
