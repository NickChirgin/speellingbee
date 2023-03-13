package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nickchirgin/spellingbee/cmd/auth/proto/auth"
	"google.golang.org/grpc"
)

func main(){
	fmt.Println("Auth")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error while listening %v", err)
	}
	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, &AuthServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}