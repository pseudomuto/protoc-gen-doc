FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y wget unzip
RUN wget -nv http://download.opensuse.org/repositories/home:estan:protoc-gen-doc/xUbuntu_16.04/Release.key -O Release.key
RUN apt-key add - < Release.key

RUN wget -nv https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip -O protoc3.zip
RUN unzip protoc3.zip -d protoc3
RUN mv protoc3/bin/protoc /usr/bin/protoc
RUN mv protoc3/include/* /usr/include/

RUN sh -c "echo 'deb http://download.opensuse.org/repositories/home:/estan:/protoc-gen-doc/xUbuntu_16.04/ /' > /etc/apt/sources.list.d/protoc-gen-doc.list"
RUN apt-get update || true
RUN apt-get install -y protoc-gen-doc

RUN rm -rf Release.key protoc3.zip proto3
RUN apt-get purge -y wget unzip
RUN apt-get autoremove -y
RUN apt-get clean all