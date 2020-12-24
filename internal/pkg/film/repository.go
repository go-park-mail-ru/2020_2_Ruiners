package film

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	FindByLId(id int) (*models.Film, error)
	FindFilmsByGenre(genre string) (*models.FilmCards, error)
	FindFilmsByPerson(id int) (*models.FilmCards, error)
	FindFilmsByPlaylist(id int) (*models.FilmCards, error)
	SimilarFilms(id int) (*models.FilmCards, error)
	Search(search string) (*models.FilmCards, error)
}
