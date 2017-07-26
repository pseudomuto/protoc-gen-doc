#!/bin/bash
set -euo pipefail

exec protoc --doc_out=/out "$@" protos/*.proto
