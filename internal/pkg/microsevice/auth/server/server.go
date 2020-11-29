package server

import (
	SessionRep "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/session/repository"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
	UserRep "github.com/Arkadiyche/http-rest-api/internal/pkg/user/repository"
	usecase2 "github.com/Arkadiyche/http-rest-api/internal/pkg/user/usecase"
)

type Server struct {
	port string
	auth *AuthServer
}

func NewServer(port string, db *store.Store) *Server {
	SessionsRep := SessionRep.NewSessionRepository(db.Db)
	UserRep := UserRep.NewUserRepository(db.Db)
	UserUC := usecase2.NewUserUseCase(UserRep, SessionsRep)

	return &Server{
		port: port,
		auth: NewAuthServer(UserUC),
	}
}