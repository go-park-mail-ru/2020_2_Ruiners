package repository

import (
	"database/sql"
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
)

type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{
		db: db,
	}
}

func (r *PersonRepository) FindByFilmIdAndRole(id int, role string) (*models.FilmPersons, error) {
	filmPerson := models.FilmPerson{}
	filmPersons := models.FilmPersons{}
	filmPersonQuery, err := r.db.Query("SELECT p.id, p.name  FROM person_film f join person p ON f.person_id = p.id WHERE f.film_id = ? AND f.role = ?", id, role)

	if err != nil {
		return nil, err
	}
	defer filmPersonQuery.Close()

	for filmPersonQuery.Next() {
		if filmPersonQuery.Scan(&filmPerson.Id, &filmPerson.Name) != nil {
			return nil, errors.New("db error")
		}
		filmPersons = append(filmPersons, filmPerson)
	}
	return &filmPersons, nil
}

func (r *PersonRepository) FindById(id int) (*models.Person, error) {
	person := models.Person{}
	err := r.db.QueryRow("SELECT id, name, image, born_date, born_place FROM person WHERE id = ?", id).Scan(&person.Id, &person.Name, &person.Image, &person.BornDate, &person.BornPlace)

	if err != nil {
		return nil, errors.New("person not found")
	}

	return &person, nil
}

func (r *PersonRepository) FindFilmsIdByPersonId(id int) ([]int, error) {
	var filmId int
	var ids []int
	filmIdQuery, err := r.db.Query("SELECT film_id  FROM person_film WHERE person_id = ?", id)

	if err != nil {
		return nil, err
	}
	defer filmIdQuery.Close()

	for filmIdQuery.Next() {
		if filmIdQuery.Scan(&filmId) != nil {
			return nil, errors.New("db error")
		}
		ids = append(ids, filmId)
	}
	return ids, err
}
