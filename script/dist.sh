#!/bin/bash
set -euo pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../" && pwd)"

package_dist() {
  local build_dir="${1}"
  local target="${2}"

  pushd "${build_dir}" >/dev/null
  tar czf "${target}.tar.gz" "${target}"
  mv "${target}.tar.gz" "${DIR}/dist"
  popd >/dev/null
}

build_dist() {
  local os="${1}"
  local version="${2}"
  local go_version="${3}"

  local build=$(mktemp -d /tmp/protoc-gen-doc.XXXXXX)
  local target="protoc-gen-doc-${version}.${os}-amd64.${go_version}"

  local ext=""
  if [ "${os}" = "windows" ]; then ext=".exe"; fi

  echo -n "Building ${target}..."
  GOOS="${os}" GOARCH=amd64 CGO_ENABLED=0 \
    go build -ldflags="-s -w" -o "${build}/${target}/protoc-gen-doc${ext}" ./cmd/... || exit 1

  package_dist "${build}" "${target}"
  rm -rf "${build}"
  echo "done."
}

main() {
  rm -rf "${DIR}/dist"
  mkdir -p "${DIR}/dist"

  local app_version=$(grep "const VERSION" "${DIR}/version.go" | awk '{print $NF }' | tr -d '"')
  local go_version=$(go version | awk '{print $3}')

  for target in windows linux darwin; do
    build_dist "${target}" "${app_version}" "${go_version}"
  done
}

main "$@"
