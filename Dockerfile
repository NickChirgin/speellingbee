FROM golang:1.20

RUN apt-get update
RUN apt-get install unzip

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip && \
  unzip -o protoc-3.9.1-linux-x86_64.zip -d /usr/local bin/protoc && \
  unzip -o protoc-3.9.1-linux-x86_64.zip -d /usr/local include/* && \
  rm -rf protoc-3.9.1-linux-x86_64.zip

RUN export PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /app

COPY . .



RUN go mod tidy

RUN protoc --go_out=cmd/auth/proto proto/auth.proto --go-grpc_out=cmd/auth/proto proto/auth.proto

RUN go build -o ./bin/auth ./cmd/auth/main.go

CMD ./bin/auth