package rating

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	AddRating(rating int, filmId int, userId int) error
	UpdateRating(rating int, filmId int, userId int) error
	CheckRating(filmId int, userId int) (bool, error)
	AddReview(body string, filmId int, userId int) error
	GetReviewsByFilmId(filmId int) (*models.Reviews, error)
	GetUserById(id int) (string, error)
	GetRating(filmId int, userId int) (int, error)
}
