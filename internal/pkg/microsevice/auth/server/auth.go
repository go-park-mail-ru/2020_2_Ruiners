package server

import (
	"context"
	"fmt"
	pb "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/proto"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	UseCase user.UseCase
}

func NewAuthServer(UserUC user.UseCase) *AuthServer {
	return &AuthServer{UseCase: UserUC}
}

func (a *AuthServer) Signup(ctx context.Context, usr *pb.AuthUserSignup) (*pb.AuthSessionId, error) {
	fmt.Println("BBBBBBBBBBBBBBBBBBBBBb")
	sessionId, err := a.UseCase.Signup(usr.Login, usr.Email, usr.Password)
	if err != nil {
		fmt.Println(err, "SUKA")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.AuthSessionId{
		SessionId: sessionId,
	}, nil
}

func (a *AuthServer) Login(ctx context.Context, usr *pb.AuthUserLogin) (*pb.AuthSessionId, error) {
	sessionId, err := a.UseCase.Login(usr.Login, usr.Password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.AuthSessionId{
		SessionId: sessionId,
	}, nil
}


func (a *AuthServer) Logout(ctx context.Context, sid *pb.AuthSessionId) (*pb.AuthEmpty, error) {
	err := a.UseCase.Logout(sid.SessionId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AuthEmpty{}, nil
}