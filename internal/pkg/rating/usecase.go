package rating

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type UseCase interface {
	Rate(rating int, filmId int, session string) error
	AddReview(body string, filmId int, session string) error
	GetReviews(filmId string) (*models.Reviews, error)
	GetCurrentRating(filmId int, session string) (int, error)
}
