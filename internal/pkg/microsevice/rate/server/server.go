package server

import (
	"database/sql"
	pb "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/proto"
	RatingRep"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/rating/repository"
	RatingUC"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/rating/usecase"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/client"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/store"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	port string
	rate *RateServer
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
	RatingRep := RatingRep.NewRatingRepository(db.Db)
	RatingUC := RatingUC.NewRatingUseCase(RatingRep, rpcSession)

	return &Server{
		port: port,
		rate: NewRateServer(RatingUC),
	}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", ":8003")
	if err != nil {
		return err
	}

	gServer := grpc.NewServer()
	pb.RegisterRateServer(gServer, s.rate)
	err = gServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}



