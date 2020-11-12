package person

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type UseCase interface {
	GetPerson(id string) (*models.Person, error)
	GetPersonsByFilm(id string, role string) (*models.FilmPersons, error)
}
