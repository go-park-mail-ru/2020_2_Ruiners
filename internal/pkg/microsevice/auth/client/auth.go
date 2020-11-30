package client

import (
	"context"
	"fmt"
	pb "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
)

type AuthClient struct {
	client pb.AuthClient
	gConn  *grpc.ClientConn
}

func NewAuthClient(host, port string) (*AuthClient, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(host+port, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	return &AuthClient{client: pb.NewAuthClient(conn), gConn: conn}, nil
}

func (a *AuthClient) Signup(login, email, password string) (sessionId string, err error) {
	usr := &pb.AuthUserSignup{
		Login:    login,
		Email:    email,
		Password: password,
	}
	session, err := a.client.Signup(context.Background(), usr)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return session.SessionId, nil
}

func (a *AuthClient) Login(login, password string) (sessionId string, err error) {
	usr := &pb.AuthUserLogin{
		Login:    login,
		Password: password,
	}

	session, err := a.client.Login(context.Background(), usr)
	if err != nil {
		return "", err
	}

	return session.SessionId, nil
}

func (a *AuthClient) Logout(sessionId string) error {
	sid := &pb.AuthSessionId{SessionId: sessionId}

	_, err := a.client.Logout(context.Background(), sid)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthClient) Close() {
	if err := a.gConn.Close(); err != nil {
		log.Fatal("error while closing grpc connection")
	}
}

