package repository

import (
	"fmt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"reflect"
	"testing"
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

