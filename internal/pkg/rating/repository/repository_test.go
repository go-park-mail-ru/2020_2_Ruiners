package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCheckRating(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRatingRepository(db)

	var filmID = 1
	var userID = 10

	// good query
	rows := sqlmock.
		NewRows([]string{"id"}).
		AddRow(1)

	mock.
		ExpectQuery("SELECT").
		WithArgs(filmID, userID).
		WillReturnRows(rows)

	item, err := repo.CheckRating(filmID, userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, true) {
		t.Errorf("results not match, want %v, have %v", true, item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(filmID, userID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.CheckRating(filmID, userID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// good query empty
	rows = sqlmock.NewRows([]string{"id"})

	mock.
		ExpectQuery("SELECT").
		WithArgs(filmID, userID).
		WillReturnRows(rows)

	item, err = repo.CheckRating(filmID, userID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err != nil {
		t.Errorf("expected error, got nil")
		return
	}
	if !reflect.DeepEqual(item, false) {
		t.Errorf("results not match, want %v, have %v", false, item)
		return
	}
}

func TestGetReviewsByFilmId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRatingRepository(db)

	var filmID = 3

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "body", "film_id", "user_id"})
	expect := models.Reviews{
		{3, "", "text", 7, 8, 0},
		{3, "", "text2", 7, 10, 0},
		{3, "", "text5", 5, 2, 0},
	}

	for _, review := range expect {
		rows = rows.AddRow(review.Id, review.TextBody, review.FilmId, review.UserId)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(filmID).
		WillReturnRows(rows)

	item, err := repo.GetReviewsByFilmId(filmID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(*item, expect) {
		t.Errorf("results not match, want %v, have %v", expect, *item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(filmID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetReviewsByFilmId(filmID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "name", "count_rolls"}).
		AddRow(1, "title", 3)

	mock.
		ExpectQuery("SELECT").
		WithArgs(filmID).
		WillReturnRows(rows)

	_, err = repo.GetReviewsByFilmId(filmID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRatingRepository(db)

	var userID = 3
	var username = "Misha"

	// good query
	rows := sqlmock.
		NewRows([]string{"username"}).AddRow(username)

	mock.
		ExpectQuery("SELECT").
		WithArgs(userID).
		WillReturnRows(rows)

	item, err := repo.GetUserById(userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, username) {
		t.Errorf("results not match, want %v, have %v", username, item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(userID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetUserById(userID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "name", "count_rolls"}).
		AddRow(1, "title", 3)

	mock.
		ExpectQuery("SELECT").
		WithArgs(userID).
		WillReturnRows(rows)

	_, err = repo.GetUserById(userID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetRating(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRatingRepository(db)

	userID := 3
	filmId := 5
	rating := 8

	// good query
	rows := sqlmock.
		NewRows([]string{"rating"}).AddRow(rating)

	mock.
		ExpectQuery("SELECT").
		WithArgs(filmId, userID).
		WillReturnRows(rows)

	item, err := repo.GetRating(filmId, userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, rating) {
		t.Errorf("results not match, want %v, have %v", rating, item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(filmId, userID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetRating(filmId, userID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "name", "count_rolls"}).
		AddRow(1, "title", 3)

	mock.
		ExpectQuery("SELECT").
		WithArgs(filmId, userID).
		WillReturnRows(rows)

	_, err = repo.GetRating(filmId, userID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestAddRating(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRatingRepository(db)

	rating := 7
	filmID := 3
	userID := 4

	// good query
	mock.
		ExpectExec(`INSERT INTO rating`).
		WithArgs(rating, filmID, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddRating(rating, filmID, userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// query error
	mock.
		ExpectExec(`INSERT INTO rating`).
		WithArgs(rating, filmID, userID).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.AddRating(rating, filmID, userID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestUpdateRating(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRatingRepository(db)

	rating := 7
	filmID := 3
	userID := 4

	// good query
	mock.
		ExpectExec(`UPDATE rating`).
		WithArgs(rating, filmID, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateRating(rating, filmID, userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// query error
	mock.
		ExpectExec(`UPDATE rating`).
		WithArgs(rating, filmID, userID).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.UpdateRating(rating, filmID, userID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestAddReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRatingRepository(db)

	body := "name"
	filmID := 3
	userID := 4

	// good query
	mock.
		ExpectExec(`INSERT INTO review`).
		WithArgs(body, filmID, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddReview(body, filmID, userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// query error
	mock.
		ExpectExec(`INSERT INTO review`).
		WithArgs(body, filmID, userID).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.AddReview(body, filmID, userID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}
