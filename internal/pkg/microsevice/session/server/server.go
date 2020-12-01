package server

import (
	"database/sql"
	pb "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/proto"
	SessionRep "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session/repository"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	port string
	session *SessionServer
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
	SessionsRep := SessionRep.NewSessionRepository(db.Db)
	return &Server{
		port: port,
		session: NewSessionServer(*SessionsRep),
	}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		return err
	}

	gServer := grpc.NewServer()
	pb.RegisterSessionsServer(gServer, s.session)
	err = gServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
