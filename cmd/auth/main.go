package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/nickchirgin/spellingbee/cmd/auth/authpb"
	jwtmanager "github.com/nickchirgin/spellingbee/internal/jwtManager"
	"google.golang.org/grpc"
)

type AuthServer struct {		
	authpb.UnimplementedAuthServiceServer
	jwtManager *jwtmanager.JWTManager
}

const (
	SecretKey = "JustASecretKey"
	tokenDuration = time.Hour
)
func NewAuthServer(jwt *jwtmanager.JWTManager) *AuthServer {
	return &AuthServer{jwtManager: jwt}
}

func (a *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {



	res := &authpb.LoginResponse{AccessToken: " "}
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
	jwt := jwtmanager.NewJWTManager(SecretKey, tokenDuration) 
	a := NewAuthServer(jwt)
	s := grpc.NewServer(grpc.UnaryInterceptor(AuthUnaryInterceptor))
	authpb.RegisterAuthServiceServer(s, a)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}