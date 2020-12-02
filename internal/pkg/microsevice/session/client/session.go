package client

import (
	"context"
	"fmt"
	pb "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
)

type SessionClient struct {
	client pb.SessionsClient
	gConn  *grpc.ClientConn
}

func NewSessionClient(host, port string) (*SessionClient, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(host+port, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	return &SessionClient{client: pb.NewSessionsClient(conn), gConn: conn}, nil
}

func (s *SessionClient) Create(sessionId, login string) (err error) {
	usr := &pb.Session{
		SessionId: sessionId,
		Login:     login,
	}
	_, err = s.client.Create(context.Background(), usr)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionClient) FindById(session string) (sessionId, login string, err error) {
	fmt.Println("find")
	se := &pb.SessionId{
		SessionId: session,
	}

	ses, err := s.client.FindById(context.Background(), se)
	if err != nil {
		return "", "", err
	}

	return ses.SessionId, ses.Login, nil
}

func (s *SessionClient) Delete(session string) (err error) {
	se := &pb.SessionId{
		SessionId: session,
	}

	_, err = s.client.Delete(context.Background(), se)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionClient) UpdateLogin(old, new string) (err error) {
	se := &pb.UpdateLoginn{
		Old: old,
		New: new,
	}

	_, err = s.client.UpdateLogin(context.Background(), se)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionClient) GetUserIdBySession(session string) (int, error) {
	se := &pb.SessionId{
		SessionId: session,
	}

	UserId, err := s.client.GetUserIdBySession(context.Background(), se)
	if err != nil {
		return 0, err
	}

	return int(UserId.UserId), err
}

func (s *SessionClient) Close() {
	if err := s.gConn.Close(); err != nil {
		log.Fatal("error while closing grpc connection")
	}
}
