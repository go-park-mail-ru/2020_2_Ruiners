package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestAddSubscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSubscribeRepository(db)

	subscriberId := 30
	authorId := 5

	// good query
	mock.
		ExpectExec(`INSERT INTO subscribe`).
		WithArgs(subscriberId, authorId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddSubscribe(subscriberId, authorId)
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
		ExpectExec(`INSERT INTO subscribe`).
		WithArgs(subscriberId, authorId).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.AddSubscribe(subscriberId, authorId)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestDeleteSubscribe(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSubscribeRepository(db)

	subscriberId := 30
	authorId := 5

	// good query
	mock.
		ExpectExec(`DELETE FROM subscribe`).
		WithArgs(subscriberId, authorId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteSubscribe(subscriberId, authorId)
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
		ExpectExec(`DELETE FROM subscribe`).
		WithArgs(subscriberId, authorId).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.DeleteSubscribe(subscriberId, authorId)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestGetAuthors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSubscribeRepository(db)

	subscriberId := 30

	rows := sqlmock.
		NewRows([]string{"id", "username", "email"})
	expect := models.PublicUsers{
		{Id: 1, Login: "title1", Email: "gen@r.e"},
		{Id: 2, Login: "title2", Email: "ge@nr.e"},
		{Id: 9, Login: "title9", Email: "ge@n.re"},
	}

	for _, people := range expect {
		rows = rows.AddRow(people.Id, people.Login, people.Email)
	}

	// good query

	mock.
		ExpectQuery("SELECT").
		WithArgs(subscriberId).
		WillReturnRows(rows)

	item, err := repo.GetAuthors(subscriberId)
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
		WithArgs(subscriberId).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetAuthors(subscriberId)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "title")

	mock.
		ExpectQuery("SELECT").
		WithArgs(subscriberId).
		WillReturnRows(rows)

	_, err = repo.GetAuthors(subscriberId)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetRatingFeed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSubscribeRepository(db)

	subscriberId := 30

	rows := sqlmock.
		NewRows([]string{"rating", "user_id", "username", "film_id", "title", "date"})
	expect := models.Feed{
		{Id: 0, IsRating: true, IsReview: false, Body: "j", UserId: 3, UserLogin: "title1", FilmId: 6, FilmTitle: "gen@r.e", Date: 78},
		{Id: 0, IsRating: true, IsReview: false, Body: "k", UserId: 3, UserLogin: "title2", FilmId: 7, FilmTitle: "gen@hr.e", Date: 782},
		{Id: 0, IsRating: true, IsReview: false, Body: "A", UserId: 3, UserLogin: "title3", FilmId: 8, FilmTitle: "gen@r.e", Date: 78},
	}

	for _, elem := range expect {
		rows = rows.AddRow(elem.Body, elem.UserId, elem.UserLogin, elem.FilmId, elem.FilmTitle, elem.Date)
	}

	// good query

	mock.
		ExpectQuery("SELECT").
		WithArgs(subscriberId).
		WillReturnRows(rows)

	item, err := repo.GetRatingFeed(subscriberId)
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
		WithArgs(subscriberId).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetRatingFeed(subscriberId)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "title")

	mock.
		ExpectQuery("SELECT").
		WithArgs(subscriberId).
		WillReturnRows(rows)

	_, err = repo.GetRatingFeed(subscriberId)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetReviewFeed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSubscribeRepository(db)

	subscriberId := 30

	rows := sqlmock.
		NewRows([]string{"rating", "user_id", "username", "film_id", "title", "date"})
	expect := models.Feed{
		{Id: 0, IsRating: false, IsReview: true, Body: "j", UserId: 3, UserLogin: "title1", FilmId: 6, FilmTitle: "gen@r.e", Date: 78},
		{Id: 0, IsRating: false, IsReview: true, Body: "k", UserId: 3, UserLogin: "title2", FilmId: 7, FilmTitle: "gen@hr.e", Date: 782},
		{Id: 0, IsRating: false, IsReview: true, Body: "A", UserId: 3, UserLogin: "title3", FilmId: 8, FilmTitle: "gen@r.e", Date: 78},
	}

	for _, elem := range expect {
		rows = rows.AddRow(elem.Body, elem.UserId, elem.UserLogin, elem.FilmId, elem.FilmTitle, elem.Date)
	}

	// good query

	mock.
		ExpectQuery("SELECT").
		WithArgs(subscriberId).
		WillReturnRows(rows)

	item, err := repo.GetReviewFeed(subscriberId)
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
		WithArgs(subscriberId).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetReviewFeed(subscriberId)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "title")

	mock.
		ExpectQuery("SELECT").
		WithArgs(subscriberId).
		WillReturnRows(rows)

	_, err = repo.GetReviewFeed(subscriberId)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCheck(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSubscribeRepository(db)

	subscriberId := 30
	authorId := 5

	// good query
	rows := sqlmock.
		NewRows([]string{"id"}).
		AddRow(6)

	mock.
		ExpectQuery("SELECT").
		WithArgs(subscriberId, authorId).
		WillReturnRows(rows)

	item, err := repo.Check(subscriberId, authorId)
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
		WithArgs(subscriberId, authorId).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.Check(subscriberId, authorId)

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
		WithArgs(subscriberId, authorId).
		WillReturnRows(rows)

	_, err = repo.Check(subscriberId, authorId)
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	rows = sqlmock.NewRows([]string{"id"})

	mock.
		ExpectQuery("SELECT").
		WithArgs(subscriberId, authorId).
		WillReturnRows(rows)

	item, err = repo.Check(subscriberId, authorId)
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
