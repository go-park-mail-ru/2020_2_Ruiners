package film

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type UseCase interface {
	FindById(id string) (*models.Film, error)
	FilmsByGenre(genre string) (*models.FilmCards, error)
	FilmsByPerson(genre string) (*models.FilmCards, error)
	SimilarFilms(id string) (*models.FilmCards, error)
}
