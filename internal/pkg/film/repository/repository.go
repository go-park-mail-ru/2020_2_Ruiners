package repository

import (
	"database/sql"
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
)

type FilmRepository struct {
	db *sql.DB
}

func NewFilmRepository(db *sql.DB) *FilmRepository {
	return &FilmRepository{
		db: db,
	}
}

func (r *FilmRepository) FindByLId(id int) (*models.Film, error) {
	film := models.Film{}
	filmQuery, err := r.db.Query("SELECT id, title, description, mainGenre, youtubeLink, bigImg, smallImg, year, country  FROM films WHERE id = ?", id)
	defer filmQuery.Close()
	if err != nil {
		return nil, err
	}
	if filmQuery.Next() {
		filmQuery.Scan(&film.Id, &film.Title, &film.Description, &film.MainGenre, &film.YoutubeLink, &film.BigImg, &film.SmallImg, &film.Year, &film.Country)
	} else {
		return nil, errors.New("film not found")
	}
	return &film, nil
}