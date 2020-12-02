package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/rating"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/client"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session"
)

type RatingUseCase struct {
	RatingRepository rating.Repository
	RpcSession       client.ISessionClient
}

func NewRatingUseCase(ratingRepository rating.Repository, rpcSession session.Repository) *RatingUseCase {
	return &RatingUseCase{
		RatingRepository: ratingRepository,
		RpcSession:       rpcSession,
	}
}

func (uc *RatingUseCase) Rate(rating int, filmId int, session string) error {
	userId, err := uc.RpcSession.GetUserIdBySession(session)
	if err != nil {
		return err
	}
	check, err := uc.RatingRepository.CheckRating(filmId, userId)
	if err != nil {
		return err
	}
	if !check {
		uc.RatingRepository.AddRating(rating, filmId, userId)
	} else {
		uc.RatingRepository.UpdateRating(rating, filmId, userId)
	}
	return nil
}

func (uc *RatingUseCase) AddReview(body string, filmId int, session string) error {
	userId, err := uc.RpcSession.GetUserIdBySession(session)
	if err != nil {
		return err
	}
	err = uc.RatingRepository.AddReview(body, filmId, userId)
	if err != nil {
		return err
	}
	return nil
}
