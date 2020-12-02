package server

import (
	"context"
	pb "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/proto"
	sessionRep "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SessionServer struct {
	Rep sessionRep.SessionRepository
}

func NewSessionServer(Rep sessionRep.SessionRepository) *SessionServer {
	return &SessionServer{Rep: Rep}
}

func (s *SessionServer) Create(ctx context.Context, session *pb.Session) (*pb.Empty, error) {
	err := s.Rep.Create(session.SessionId, session.Login)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.Empty{}, nil
}

func (s *SessionServer) FindById(ctx context.Context, sessionId *pb.SessionId) (*pb.Session, error) {
	sId, sLogin, err := s.Rep.FindById(sessionId.SessionId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.Session{
		SessionId: sId,
		Login:     sLogin}, nil
}

func (s *SessionServer) Delete(ctx context.Context, sessionId *pb.SessionId) (*pb.Empty, error) {
	err := s.Rep.Delete(sessionId.SessionId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.Empty{}, nil
}

func (s *SessionServer) UpdateLogin(ctx context.Context, updateLogin *pb.UpdateLoginn) (*pb.Empty, error) {
	err := s.Rep.UpdateLogin(updateLogin.Old, updateLogin.New)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.Empty{}, nil
}

func (s *SessionServer) GetUserIdBySession(ctx context.Context, sessionId *pb.SessionId) (*pb.UserId, error) {
	userId, err := s.Rep.GetUserIdBySession(sessionId.SessionId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.UserId{
		UserId: int64(userId),
	}, nil
}
