#!/bin/bash
set -euo pipefail

# this is required because of the wildcard expansion. Passing protos/*.proto in CMD gets escaped -_-.
exec protoc --doc_out=/out "$@" protos/*.proto
