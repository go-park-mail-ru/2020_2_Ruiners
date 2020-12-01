package usecase

import (
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/client"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/rating"
	"strconv"
)

type RatingUseCase struct {
	RatingRepository  rating.Repository
	RpcSession client.ISessionClient
}

func NewRatingUseCase(ratingRepository rating.Repository, rpcSession session.Repository) *RatingUseCase {
	return &RatingUseCase{
		RatingRepository:  ratingRepository,
		RpcSession: rpcSession,
	}
}

func (uc *RatingUseCase) GetReviews(filmId string) (*models.Reviews, error) {
	rs := models.Reviews{}
	id, err := strconv.Atoi(filmId)
	if err != nil {
		return nil, err
	}
	reviews, err := uc.RatingRepository.GetReviewsByFilmId(id)
	if err != nil {
		return nil, err
	}
	for _, r := range *reviews {
		login, err := uc.RatingRepository.GetUserById(r.UserId)
		if err != nil {
			login = "Deleted"
		}
		rate, err := uc.RatingRepository.GetRating(r.FilmId, r.UserId)
		if err != nil {
			rate = 0
		}
		r.UserLogin = login
		r.Rate = rate
		rs = append(rs, r)
	}
	fmt.Println(rs)
	return &rs, nil
}

func (uc *RatingUseCase) GetCurrentRating(filmId string, session string) (int, error) {
	id, err := strconv.Atoi(filmId)
	if err != nil {
		return 0, err
	}
	userId, err := uc.RpcSession.GetUserIdBySession(session)
	if err != nil {
		return 0, err
	}
	rate, err := uc.RatingRepository.GetRating(id, userId)
	if err != nil {
		rate = 0
	}
	return rate, nil
}
