package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSessionRepository(db)

	var session_id = "hhh5"

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "username"}).
		AddRow(session_id, "Admin")

	mock.
		ExpectQuery("SELECT").
		WithArgs(session_id).
		WillReturnRows(rows)

	item, err := repo.FindById(session_id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item.Id, session_id) {
		t.Errorf("results not match, want %v, have %v", session_id, item.Id)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(session_id).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.FindById(session_id)

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
		WithArgs(session_id).
		WillReturnRows(rows)

	_, err = repo.FindById(session_id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestGetUserIdBySession(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSessionRepository(db)

	session_id := "hhh5"
	userId := 1

	// good query
	rows := sqlmock.
		NewRows([]string{"id"}).
		AddRow(userId)

	mock.
		ExpectQuery("SELECT").
		WithArgs(session_id).
		WillReturnRows(rows)

	item, err := repo.GetUserIdBySession(session_id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, userId) {
		t.Errorf("results not match, want %v, have %v", userId, item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(session_id).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetUserIdBySession(session_id)

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
		WithArgs(session_id).
		WillReturnRows(rows)

	_, err = repo.FindById(session_id)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func forUpdate(t *testing.T, mock sqlmock.Sqlmock, f func(string, string) error) {
	srt2 := "newLogPasOrName"
	str1 := "login"

	// good query
	mock.
		ExpectExec(`UPDATE session`).
		WithArgs(srt2, str1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := f(str1, srt2)
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
		ExpectExec(`UPDATE session`).
		WithArgs(srt2, str1).
		WillReturnError(fmt.Errorf("db_error"))

	err = f(str1, srt2)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSessionRepository(db)
	forUpdate(t, mock, repo.UpdateLogin)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewSessionRepository(db)

	session := models.Session{"hhh5", "username"}

	// good query
	mock.
		ExpectExec(`INSERT INTO session`).
		WithArgs(session.Id, session.Username).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(&session)
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
		ExpectExec(`INSERT INTO session`).
		WithArgs(session.Id, session.Username).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.Create(&session)

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

	repo := NewSessionRepository(db)

	id := "hhhh"

	// good query
	mock.
		ExpectExec(`DELETE FROM session`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Delete(id)
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
		ExpectExec(`DELETE FROM session`).
		WithArgs(id).
		WillReturnError(fmt.Errorf("db_error"))

	err = repo.Delete(id)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}
