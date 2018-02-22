// protoc-gen-doc is used to generate documentation from comments in your proto files.
//
// It is a protoc plugin, and can be invoked by passing `--doc_out` and `--doc_opt` arguments to protoc.
//
// Example: generate HTML documentation
//
//     protoc --doc_out=. --doc_opt=html,index.html protos/*.proto
//
// Example: use a custom template
//
//     protoc --doc_out=. --doc_opt=custom.tmpl,docs.txt protos/*.proto
//
// For more details, check out the README at https://github.com/pseudomuto/protoc-gen-doc
package main

import (
	"github.com/pseudomuto/protokit"

	"log"
	"os"

	"github.com/pseudomuto/protoc-gen-doc"
)

func main() {
	flags := gendoc.ParseFlags(os.Stdout, os.Args)

	if flags.HasMatch() {
		if flags.ShowHelp() {
			flags.PrintHelp()
		}

		if flags.ShowVersion() {
			flags.PrintVersion()
		}

		os.Exit(flags.Code())
	}

	if err := protokit.RunPlugin(new(gendoc.Plugin)); err != nil {
		log.Fatal(err)
	}
}
