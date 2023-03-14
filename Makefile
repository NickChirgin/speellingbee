build:
	docker compose build
run: 
	docker compose up

gen:
	protoc --go_out=cmd/auth proto/auth.proto --go-grpc_out=cmd/auth proto/auth.proto
