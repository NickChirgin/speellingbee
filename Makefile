build:	
	go build -o ./bin/auth ./cmd/auth/main.go
	go build -o ./bin/cron ./cmd/cron/main.go
	go build -o ./bin/game ./cmd/game/main.go

run: 
	./bin/auth
	./bin/cron
	./bin/game

