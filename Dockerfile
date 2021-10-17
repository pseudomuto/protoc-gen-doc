FROM golang:1.17-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o protoc-gen-doc ./cmd/protoc-gen-doc

FROM debian:bookworm-slim AS final
LABEL maintainer="pseudomuto <david.muto@gmail.com>" protoc_version="3.18.1"

WORKDIR /

ADD https://github.com/google/protobuf/releases/download/v3.18.1/protoc-3.18.1-linux-x86_64.zip ./
RUN apt-get -q -y update && \
  apt-get -q -y install unzip && \
  unzip protoc-3.18.1-linux-x86_64.zip -d ./usr/local && \
  rm protoc-3.18.1-linux-x86_64.zip && \
  apt-get remove --purge -y unzip && \
  apt-get autoremove && \
  rm -rf /var/lib/apt/lists/*

COPY --from=builder /build/protoc-gen-doc /usr/local/bin
COPY script/entrypoint.sh ./

VOLUME ["/out", "/protos"]

ENTRYPOINT ["/entrypoint.sh"]
CMD ["--doc_opt=html,index.html"]
