package server

import (
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	"google.golang.org/grpc"
	"net"
)

type AuthServer struct {
	UseCase user.UseCase
}

func NewAuthServer(UserUC user.UseCase) *AuthServer {
	return &AuthServer{UseCase: UserUC}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	gServer := grpc.NewServer()
	//api.RegisterAuthServer(gServer, s.auth)
	fmt.Println(s)
	err = gServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}