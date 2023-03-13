package main

import (
	"context"

	"github.com/nickchirgin/spellingbee/cmd/auth/proto/auth"
)

type AuthServer struct {		
	auth.UnimplementedAuthServiceServer
}


func (a *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {



	res := &auth.LoginResponse{AccessToken: " "}
	return res, nil
}

