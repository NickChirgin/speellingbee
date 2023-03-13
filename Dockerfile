FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o ./bin/auth ./cmd/auth/main.go
RUN go build -o ./bin/cron ./cmd/cron/main.go
RUN go build -o ./bin/game ./cmd/game/main.go

CMD ./bin/auth; ./bin/cron; ./bin/game