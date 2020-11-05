package repository

import (
"database/sql"
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
	defer filmPersonQuery.Close()
	if err != nil {
		return nil, err
	}
	for filmPersonQuery.Next() {
		filmPersonQuery.Scan(&filmPerson.Id, &filmPerson.Name)
		filmPersons = append(filmPersons, filmPerson)
	}
	return &filmPersons, nil
}

func (r *PersonRepository) FindById(id int) (*models.Person, error) {
	person := models.Person{}
	personQuery, err := r.db.Query("SELECT id, name, image, born_date, born_place FROM person WHERE id = ?", id)
	defer personQuery.Close()
	if err != nil {
		return nil, err
	}
	if personQuery.Next() {
		personQuery.Scan(&person.Id, &person.Name, &person.Image, &person.BornDate, &person.BornPlace)
	}
	return &person, nil
}

func (r *PersonRepository) FindFilmsIdByPersonId(id int) ([]int, error) {
	var filmId int
	var ids []int
	filmIdQuery, err := r.db.Query("SELECT film_id  FROM person_film WHERE person_id = ?", id)
	defer filmIdQuery.Close()
	if err != nil {
		return nil, err
	}
	for filmIdQuery.Next() {
		filmIdQuery.Scan(&filmId)
		ids = append(ids, id)
	}
	return ids, err
}