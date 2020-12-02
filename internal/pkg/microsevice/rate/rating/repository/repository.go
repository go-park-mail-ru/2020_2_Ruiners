package repository

import "database/sql"

type RatingRepository struct {
	db *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{
		db: db,
	}
}

func (r *RatingRepository) CheckRating(filmId int, userId int) (bool, error) {
	ratingQuery, err := r.db.Query("SELECT id FROM rating WHERE film_id = ? AND user_id = ?", filmId, userId)

	if err != nil {
		return false, err
	}
	defer ratingQuery.Close()

	if ratingQuery.Next() {
		return true, nil
	}
	return false, nil
}

func (r *RatingRepository) AddRating(rating int, filmId int, userId int) error {
	_, err := r.db.Exec("INSERT INTO rating(rating, film_id, user_id) VALUE(? , ?, ?)", rating, filmId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RatingRepository) UpdateRating(rating int, filmId int, userId int) error {
	_, err := r.db.Exec("UPDATE rating SET rating = ? WHERE film_id = ? AND user_id = ?", rating, filmId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RatingRepository) AddReview(body string, filmId int, userId int) error {
	_, err := r.db.Exec("INSERT INTO review(body, film_id, user_id) VALUE(?, ?, ?)", body, filmId, userId)
	if err != nil {
		return err
	}
	return nil
}
