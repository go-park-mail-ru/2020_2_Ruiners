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
	err := r.db.QueryRow("SELECT id, title, rating, sumVotes, description, mainGenre, youtubeLink, bigImg, smallImg, year, country  FROM films WHERE id = ?", id).Scan(&film.Id, &film.Title, &film.Rating, &film.SumVotes, &film.Description, &film.MainGenre, &film.YoutubeLink, &film.BigImg, &film.SmallImg, &film.Year, &film.Country)

	if err != nil {
		return nil, errors.New("film not found")
	}

	return &film, nil
}

func (r *FilmRepository) FindFilmsByGenre(genre string) (*models.FilmCards, error) {
	filmCard := models.FilmCard{}
	filmCards := models.FilmCards{}
	filmQuery, err := r.db.Query("SELECT id, title, mainGenre, smallImg, year FROM films WHERE mainGenre = ? LIMIT 10", genre)

	if err != nil {
		return nil, err
	}
	defer filmQuery.Close()

	for filmQuery.Next() {
		if filmQuery.Scan(&filmCard.Id, &filmCard.Title, &filmCard.MainGenre, &filmCard.SmallImg, &filmCard.Year) != nil {
			return nil, errors.New("db error")
		}
		filmCards = append(filmCards, filmCard)
	}
	return &filmCards, nil
}

func (r *FilmRepository) FindFilmsByPerson(id int) (*models.FilmCards, error) {
	filmCard := models.FilmCard{}
	filmCards := models.FilmCards{}
	filmQuery, err := r.db.Query("SELECT f.id, f.title, f.mainGenre, f.smallImg, f.year FROM films f JOIN person_film p ON f.id = p.film_id WHERE p.person_id = ? LIMIT 10", id)

	if err != nil {
		return nil, err
	}
	defer filmQuery.Close()

	for filmQuery.Next() {
		if filmQuery.Scan(&filmCard.Id, &filmCard.Title, &filmCard.MainGenre, &filmCard.SmallImg, &filmCard.Year) != nil {
			return nil, errors.New("db error")
		}
		filmCards = append(filmCards, filmCard)
	}
	return &filmCards, nil
}

func (r *FilmRepository) FindFilmsByPlaylist(id int) (*models.FilmCards, error) {
	filmCard := models.FilmCard{}
	filmCards := models.FilmCards{}
	filmQuery, err := r.db.Query("SELECT  f.id, f.title, f.mainGenre, f.smallImg, f.year FROM playlist p JOIN playlist_film pf ON(p.id = pf.playlist_id) JOIN films f ON(pf.film_id = f.id) WHERE p.id=?", id)

	if err != nil {
		return nil, err
	}
	defer filmQuery.Close()

	for filmQuery.Next() {
		if filmQuery.Scan(&filmCard.Id, &filmCard.Title, &filmCard.MainGenre, &filmCard.SmallImg, &filmCard.Year) != nil {
			return nil, errors.New("db error")
		}
		filmCards = append(filmCards, filmCard)
	}
	return &filmCards, nil
}
