package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFindByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	var username = "Admin"
	admin := models.User{1, username, "fdvvvccc", "gg@gmail.com", "my_img"}

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "email", "image"}).
		AddRow(admin.Id, admin.Username, admin.Password, admin.Email, admin.Image)

	mock.
		ExpectQuery("SELECT").
		WithArgs(username).
		WillReturnRows(rows)

	item, err := repo.FindByLogin(username)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(*item, admin) {
		t.Errorf("results not match, want %v, have %v", admin, *item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(username).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.FindByLogin(username)

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
		WithArgs(username).
		WillReturnRows(rows)

	_, err = repo.FindByLogin(username)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestFindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	var user_id = 5

	admin := models.User{user_id, "username", "fdvvvccc", "gg@gmail.com", "my_img"}

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "email", "image"}).
		AddRow(admin.Id, admin.Username, admin.Password, admin.Email, admin.Image)

	mock.
		ExpectQuery("SELECT").
		WithArgs(user_id).
		WillReturnRows(rows)

	item, err := repo.FindById(user_id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(*item, admin) {
		t.Errorf("results not match, want %v, have %v", admin, *item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(user_id).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.FindById(user_id)

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
		WithArgs(user_id).
		WillReturnRows(rows)

	_, err = repo.FindById(user_id)
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
		ExpectExec(`UPDATE users`).
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
		ExpectExec(`UPDATE users`).
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

	repo := NewUserRepository(db)
	forUpdate(t, mock, repo.UpdateLogin)
	forUpdate(t, mock, repo.UpdatePassword)
	forUpdate(t, mock, repo.UpdateAvatar)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	admin := models.User{0, "username", "fdvvvccc", "gg@gmail.com", ""}

	// good query
	mock.
		ExpectExec(`INSERT INTO users`).
		WithArgs(admin.Username, admin.Password, admin.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(&admin)
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
		ExpectExec(`INSERT INTO users`).
		WithArgs(admin.Username, admin.Password, admin.Email).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.Create(&admin)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}
