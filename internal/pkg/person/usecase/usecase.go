package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/person"
	"strconv"
)

type PersonUseCase struct {
	PersonRepository person.Repository
}

func NewPersonUseCase(personRepository person.Repository) *PersonUseCase {
	return &PersonUseCase{
		PersonRepository: personRepository,
	}
}

func (uc *PersonUseCase) GetPerson(id string) (*models.Person, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	p, err := uc.PersonRepository.FindById(idInt)
	if err != nil {
		return nil, err
	}
	ids, err := uc.PersonRepository.FindFilmsIdByPersonId(p.Id)
	if err != nil {
		return nil, err
	}
	p.FilmsId = ids
	return p, nil
}

func (uc *PersonUseCase) GetPersonsByFilm(id string, role string) (*models.FilmPersons, error)  {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	persons, err := uc.PersonRepository.FindByFilmIdAndRole(idInt, role)
	if err != nil {
		return nil, err
	}
	return  persons, nil
}

