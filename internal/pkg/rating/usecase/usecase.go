package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/sesession"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/rating"
)

type RatingUseCase struct {
	RatingRepository rating.Repository
	SessionRepository sesession.Repository
}

func NewRatingUseCase(ratingRepository rating.Repository, sessionRepository sesession.Repository) *RatingUseCase {
	return &RatingUseCase{
		RatingRepository: ratingRepository,
		SessionRepository: sessionRepository,
	}
}

func (uc *RatingUseCase) Rate(rating int, filmId int, session string) error{
	userId, err := uc.SessionRepository.GetUserIdBySession(session)
	if err != nil {
		return err
	}
	check, err := uc.RatingRepository.Check(filmId, userId)
	if err != nil {
		return err
	}
	if !check {
		uc.RatingRepository.Add(rating, filmId, userId)
	} else {
		uc.RatingRepository.Update(rating, filmId, userId)
	}
	return nil
}

