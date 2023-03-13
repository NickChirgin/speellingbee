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

RUN go build -o ./bin/auth ./cmd/auth/main.go
RUN go build -o ./bin/cron ./cmd/cron/main.go
RUN go build -o ./bin/game ./cmd/game/main.go

CMD ./bin/auth; ./bin/cron; ./bin/game