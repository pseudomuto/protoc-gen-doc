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

ADD dist/protoc-gen-doc /usr/local/bin/
ADD script/entrypoint.sh ./

VOLUME ["/out", "/protos"]

ENTRYPOINT ["/entrypoint.sh"]
CMD ["--doc_opt=html,index.html"]
