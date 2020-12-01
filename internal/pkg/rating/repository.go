package rating

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	GetReviewsByFilmId(filmId int) (*models.Reviews, error)
	GetUserById(id int) (string, error)
	GetRating(filmId int, userId int) (int, error)
}
