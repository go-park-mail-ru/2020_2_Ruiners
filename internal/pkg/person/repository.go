package person

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	FindByFilmIdAndRole(id int, role string) (*models.FilmPersons, error)
	FindById(id int) (*models.Person, error)
	FindFilmsIdByPersonId(id int) ([]int, error)
}
