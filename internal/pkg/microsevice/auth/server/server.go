package server

import (
	"database/sql"
	pb "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/proto"
	UserRep "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/user/repository"
	UserUC "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/user/usecase"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/client"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	port string
	auth *AuthServer
}

func NewServer(port string, db *store.Store) *Server {
	sdb, err := sql.Open("mysql", db.Config())
	if err != nil {
		return nil
	}
	if err := sdb.Ping(); err != nil {
		return nil
	}
	db.Db = sdb
	rpcSession, err := client.NewSessionClient("localhost", ":8002")
	if err != nil {
		log.Fatal(err.Error())
	}
	UserRep := UserRep.NewUserRepository(db.Db)
	UserUC := UserUC.NewUserUseCase(UserRep, rpcSession)

	return &Server{
		port: port,
		auth: NewAuthServer(UserUC),
	}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		return err
	}

	gServer := grpc.NewServer()
	pb.RegisterAuthServer(gServer, s.auth)
	err = gServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
