# Google Protocol Buffers<br/>Documentation Generator

`protoc-gen-doc` is a documentation generator plugin for the
Google Protocol Buffers compiler (`protoc`). It can generate
HTML, DocBook or Markdown documentation from comments in your
`.proto` files.

The plugin depends on QtCore from Qt 5.

## Building the Plugin

Run `qmake` followed by `make` to build the plugin.

## Invoking the Plugin

The plugin is invoked by passing the `--doc_out` option to the
`protoc` compiler. The option has the following format:

    --doc_out=docbook|html|markdown|<TEMPLATE_FILENAME>,<OUT_FILENAME>:<OUT_DIR>

The `protoc-gen-doc` executable must be placed in a directory
that is in `PATH` for the `protoc` compiler to find it, or else
your must specify the path to the executable with the `--plugin`
option. For example, to generate HTML documentation for all
`.proto` files in the `proto` directory into `doc/index.html`,
type:

    protoc --doc_out=html,index.html:doc proto/*.proto

The format may be either `docbook`, `html` or `markdown` or the
name of a file containing a custom [Mustache][mustache] template.

## Documenting your Messages

Use `/** */` or `///` comments to document your files. Comments
for enumerations and messages go before the message/enumeration
definition. Comments for fields or enum values can go either
before or after the field/value definition.

## Example Output

The [examples][examples] directory contains a full example with
multiple `.proto` files, including the [Markdown][markdown],
[HTML][html] and [DocBook][docbook] output, as well as a [PDF][pdf]
obtained from running [Apache FOP][fop] on the DocBook XML.


[mustache]: http://mustache.github.io/ "Mustache - Logic-less templates"
[examples]: https://github.com/estan/protoc-gen-doc/tree/master/examples "Full Example"
[markdown]: https://github.com/estan/protoc-gen-doc/tree/master/examples/doc/example.md "Markdown Example Output"
[html]: https://rawgit.com/estan/protoc-gen-doc/master/examples/doc/example.html "HTML Example Output"
[docbook]: https://github.com/estan/protoc-gen-doc/blob/master/examples/doc/example.docbook "DocBook Example Output"
[pdf]: https://github.com/estan/protoc-gen-doc/raw/master/examples/doc/example.pdf "Formatted PDF"
[fop]: http://xmlgraphics.apache.org/fop/ "Apache FOP Website"
