package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nickchirgin/spellingbee/cmd/auth/proto/auth"
	"google.golang.org/grpc"
)

type AuthServer struct {		
	auth.UnimplementedAuthServiceServer
}

func (a *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {



	res := &auth.LoginResponse{AccessToken: " "}
	return res, nil
}

func AuthUnaryInterceptor(
	ctx context.Context, 
	req interface{}, 
	info *grpc.UnaryServerInfo, 
	handler grpc.UnaryHandler,
	) (interface{}, error) {
	log.Println("Unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func main(){
	fmt.Println("Auth")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error while listening %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(AuthUnaryInterceptor))
	auth.RegisterAuthServiceServer(s, &AuthServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}