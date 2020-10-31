package repository

import (
	"database/sql"
)

type RatingRepository struct {
	db *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{
		db: db,
	}
}

func (r *RatingRepository) Check(filmId int, userId int) (bool, error) {
	ratingQuery, err := r.db.Query("SELECT id FROM rating WHERE film_id = ? AND user_id = ?", filmId, userId)
	defer ratingQuery.Close()
	if err != nil {
		return false, err
	}
	if ratingQuery.Next() {
		return true, nil
	}
	return false, nil
}

func (r *RatingRepository) Add(rating int, filmId int, userId int) error {
	_, err := r.db.Exec("insert into rating(rating, film_id, user_id) VALUE(? , ?, ?)", rating, filmId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RatingRepository) Update(rating int, filmId int, userId int) error {
	_, err := r.db.Exec("update rating set rating = ? where film_id = ? and user_id = ?", rating, filmId, userId)
	if err != nil {
		return err
	}
	return nil
}