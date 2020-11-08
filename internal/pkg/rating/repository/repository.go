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
	defer ratingQuery.Close()
	if err != nil {
		return false, err
	}
	if ratingQuery.Next() {
		return true, nil
	}
	return false, nil
}

func (r *RatingRepository) AddRating(rating int, filmId int, userId int) error {
	_, err := r.db.Exec("insert into rating(rating, film_id, user_id) VALUE(? , ?, ?)", rating, filmId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RatingRepository) UpdateRating(rating int, filmId int, userId int) error {
	_, err := r.db.Exec("update rating set rating = ? where film_id = ? and user_id = ?", rating, filmId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RatingRepository) AddReview(body string, filmId int, userId int) error {
	_, err := r.db.Exec("insert into review(body, film_id, user_id) VALUE(?, ?, ?)", body, filmId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RatingRepository) GetReviewsByFilmId(filmId int) (*models.Reviews, error) {
	review := models.Review{}
	rev := models.Reviews{}
	reviewQuery, err := r.db.Query("SELECT id, body, film_id, user_id FROM review WHERE film_id = ?", filmId)
	defer reviewQuery.Close()
	if err != nil {
		return nil, err
	}
	for reviewQuery.Next() {
		reviewQuery.Scan(&review.Id, &review.TextBody, &review.FilmId, &review.UserId)
		if err != nil {
			return nil, err
		}
		rev = append(rev, review)
	}
	return &rev, nil
}

func (r *RatingRepository) GetUserById(id int) (string, error) {
	var login string
	queryUser, err := r.db.Query("SELECT username FROM users WHERE id = ?", id)
	defer queryUser.Close()
	if err != nil {
		return "", err
	}
	if queryUser.Next() {
		queryUser.Scan(&login)
	} else {
		return "", errors.New("no user")
	}
	return login, nil
}

func (r *RatingRepository) GetRating(filmId int, userId int) (int, error) {
	var rating int
	ratingQuery, err := r.db.Query("SELECT rating FROM rating WHERE film_id = ? AND user_id = ?", filmId, userId)
	defer ratingQuery.Close()
	if err != nil {
		return 0, err
	}
	if ratingQuery.Next() {
		ratingQuery.Scan(&rating)
	}
	fmt.Println(rating, filmId, userId)
	return rating, nil
}
