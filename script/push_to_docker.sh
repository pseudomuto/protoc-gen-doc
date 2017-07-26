#!/bin/bash
set -euo pipefail

build_binary() {
  mkdir -p dist
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o dist/protoc-gen-doc ./cmd/...
}

build_and_tag_image() {
  docker build -t "${1}" .
  docker tag "${1}" "${2}:${3}"
  docker tag "${1}" "${2}:latest"
}

push_image() {
  # credentials are encrypted in travis.yml
  docker login -u "${DOCKER_HUB_USER}" -p "${DOCKER_HUB_PASSWORD}"
  docker push "${1}"
  docker push "${2}:${3}"
  docker push "${2}:latest"
}

main() {
  local sha="${TRAVIS_COMMIT:-}"
  if [ -z "${sha}" ]; then sha=$(git rev-parse HEAD); fi

  local repo="pseudomuto/protoc-gen-doc"
  local version="$(grep "const VERSION" "version.go" | awk '{print $NF }' | tr -d '"')"
  local git_tag="${repo}:${sha}"

  build_binary
  build_and_tag_image "${git_tag}" "${repo}" "${version}"

  if [ -n "${DOCKER_HUB_USER:-}" ]; then
    push_image "${git_tag}" "${repo}" "${version}"
  fi
}

main "$@"
