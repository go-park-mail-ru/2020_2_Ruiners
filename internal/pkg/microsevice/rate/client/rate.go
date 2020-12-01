package client

import (
	"context"
	pb "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
)

type RateClient struct {
	client pb.RateClient
	gConn  *grpc.ClientConn
}

func NewRateClient(host, port string) (*RateClient, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(host+port, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	return &RateClient{client: pb.NewRateClient(conn), gConn: conn}, nil
}

func (r *RateClient) Rate(rating int, filmId int, session string) error {
	ra := &pb.Rating{
		Rating: int64(rating),
		FilmId: int64(filmId),
		Session: session,
	}
	_, err := r.client.Rate(context.Background(), ra)
	if err != nil {
		return err
	}

	return nil
}

func (r *RateClient) AddReview(body string, filmId int, session string) error {
	rev := &pb.Review{
		Body:    body,
		FilmId:  int64(filmId),
		Session: session,
	}

	_, err := r.client.AddReview(context.Background(), rev)
	if err != nil {
		return err
	}

	return nil
}


func (r *RateClient) Close() {
	if err := r.gConn.Close(); err != nil {
		log.Fatal("error while closing grpc connection")
	}
}


