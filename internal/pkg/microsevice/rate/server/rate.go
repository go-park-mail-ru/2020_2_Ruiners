package server

import (
	"context"
	"fmt"
	pb "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/proto"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/rating"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RateServer struct {
	UseCase rating.UseCase
}

func NewRateServer(RatingUC rating.UseCase) *RateServer {
	return &RateServer{UseCase: RatingUC}
}

func (r *RateServer) Rate(ctx context.Context, rating *pb.Rating) (*pb.RateEmpty, error) {
	err := r.UseCase.Rate(int(rating.Rating), int(rating.FilmId), rating.Session)
	if err != nil {
		fmt.Println(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.RateEmpty{}, nil
}

func (r *RateServer) AddReview(ctx context.Context, review *pb.Review) (*pb.RateEmpty, error) {
	err := r.UseCase.AddReview(review.Body, int(review.FilmId), review.Session)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.RateEmpty{}, nil
}
