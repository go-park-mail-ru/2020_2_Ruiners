package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
)

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

func (r *RatingRepository) GetReviewsByFilmId(filmId int) (*models.Reviews, error) {
	review := models.Review{}
	rev := models.Reviews{}
	reviewQuery, err := r.db.Query("SELECT id, body, film_id, user_id FROM review WHERE film_id = ?", filmId)

	if err != nil {
		return nil, err
	}
	defer reviewQuery.Close()

	for reviewQuery.Next() {
		err = reviewQuery.Scan(&review.Id, &review.TextBody, &review.FilmId, &review.UserId)
		if err != nil {
			return nil, err
		}
		rev = append(rev, review)
	}
	return &rev, nil
}

func (r *RatingRepository) GetUserById(id int) (string, error) {
	var login string
	err := r.db.QueryRow("SELECT username FROM users WHERE id = ?", id).Scan(&login)

	if err != nil {
		return "", errors.New("no user")
	}

	return login, nil
}

func (r *RatingRepository) GetRating(filmId int, userId int) (int, error) {
	var rating int
	err := r.db.QueryRow("SELECT rating FROM rating WHERE film_id = ? AND user_id = ?", filmId, userId).Scan(&rating)

	if err != nil {
		return 0, err
	}

	fmt.Println(rating, filmId, userId)
	return rating, nil
}
