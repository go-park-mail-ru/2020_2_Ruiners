package usecase

import (
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/sesession"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/rating"
	"strconv"
)

type RatingUseCase struct {
	RatingRepository  rating.Repository
	SessionRepository sesession.Repository
}

func NewRatingUseCase(ratingRepository rating.Repository, sessionRepository sesession.Repository) *RatingUseCase {
	return &RatingUseCase{
		RatingRepository:  ratingRepository,
		SessionRepository: sessionRepository,
	}
}

func (uc *RatingUseCase) Rate(rating int, filmId int, session string) error {
	userId, err := uc.SessionRepository.GetUserIdBySession(session)
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
	userId, err := uc.SessionRepository.GetUserIdBySession(session)
	if err != nil {
		return err
	}
	err = uc.RatingRepository.AddReview(body, filmId, userId)
	if err != nil {
		return err
	}
	return nil
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

func (uc *RatingUseCase) GetCurrentRating(filmId int, session string) (int, error) {
	userId, err := uc.SessionRepository.GetUserIdBySession(session)
	if err != nil {
		return 0, err
	}
	rate, err := uc.RatingRepository.GetRating(filmId, userId)
	if err != nil {
		rate = 0
	}
	return rate, nil
}
