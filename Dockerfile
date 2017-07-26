FROM debian:jessie-slim
LABEL maintainer="david.muto@gmail.com" version="1.0.0"

WORKDIR /

ADD https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip ./
RUN apt-get -q -y update && \
  apt-get -q -y install unzip && \
  unzip protoc-3.3.0-linux-x86_64.zip -d ./usr/local && \
  rm protoc-3.3.0-linux-x86_64.zip && \
  apt-get purge -y unzip && \
  apt-get autoremove

ADD script/entrypoint.sh ./

ADD dist/protoc-gen-doc-1.0.0.linux-amd64.go1.8.1.tar.gz ./
RUN mv ./protoc-gen-doc-1.0.0.linux-amd64.go1.8.1/protoc-gen-doc /usr/local/bin && \
  rm -rf ./protoc-gen-doc-*

VOLUME ["/out", "/protos"]

ENTRYPOINT ["/entrypoint.sh"]
CMD ["--doc_opt=html,index.html"]
