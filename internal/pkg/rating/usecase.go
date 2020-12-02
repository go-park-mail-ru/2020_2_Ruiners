package rating

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type UseCase interface {
	GetReviews(filmId string) (*models.Reviews, error)
	GetCurrentRating(filmId string, session string) (int, error)
}
