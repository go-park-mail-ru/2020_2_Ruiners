package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRPlaylistRepository(db)

	subscriberId := "title"
	authorId := 5

	// good query
	mock.
		ExpectExec(`INSERT INTO playlist`).
		WithArgs(subscriberId, authorId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(subscriberId, authorId)
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
		ExpectExec(`INSERT INTO playlist`).
		WithArgs(subscriberId, authorId).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.Create(subscriberId, authorId)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRPlaylistRepository(db)

	playlistId := 30

	// good query
	mock.
		ExpectExec(`DELETE FROM playlist`).
		WithArgs(playlistId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(playlistId)
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
		ExpectExec(`DELETE FROM playlist`).
		WithArgs(playlistId).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.Delete(playlistId)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRPlaylistRepository(db)

	subscriberId := 30
	authorId := 5

	// good query
	mock.
		ExpectExec(`INSERT INTO playlist_film`).
		WithArgs(subscriberId, authorId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Add(authorId, subscriberId)
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
		ExpectExec(`INSERT INTO playlist_film`).
		WithArgs(subscriberId, authorId).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.Add(authorId, subscriberId)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestRemove(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRPlaylistRepository(db)

	filmID := 6
	playlistId := 30

	// good query
	mock.
		ExpectExec(`DELETE FROM playlist_film`).
		WithArgs(filmID, playlistId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Remove(playlistId, filmID)
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
		ExpectExec(`DELETE FROM playlist_film`).
		WithArgs(filmID, playlistId).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.Remove(playlistId, filmID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestGetList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewRPlaylistRepository(db)

	userID := 30

	rows := sqlmock.
		NewRows([]string{"id", "title", "user_id"})
	expect := models.Playlists{
		{Id: 1, Title: "title1", Films: nil, UserId: 5},
		{Id: 2, Title: "title2", Films: nil, UserId: 6},
		{Id: 9, Title: "title9", Films: nil, UserId: 5},
	}

	for _, people := range expect {
		rows = rows.AddRow(people.Id, people.Title, people.UserId)
	}

	// good query

	mock.
		ExpectQuery("SELECT").
		WithArgs(userID).
		WillReturnRows(rows)

	item, err := repo.GetList(userID)
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
		WithArgs(userID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetList(userID)

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
		WithArgs(userID).
		WillReturnRows(rows)

	_, err = repo.GetList(userID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
