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
	admin := models.User{Id: 1, Username: username, Password: "fdvvvccc", Email: "gg@gmail.com", Image: "my_img"}

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

	admin := models.User{Id: user_id, Username: "username", Password: "fdvvvccc", Email: "gg@gmail.com", Image: "my_img"}

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

func TestCheckExist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	var username = "Admin"
	admin := models.User{Id: 1, Username: username, Password: "fdvvvccc", Email: "gg@gmail.com", Image: "my_img"}

	// good query
	rows := sqlmock.
		NewRows([]string{"id"}).
		AddRow(admin.Id)

	mock.
		ExpectQuery("SELECT").
		WithArgs(username).
		WillReturnRows(rows)

	item, err := repo.CheckExist(username)
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
		WithArgs(username).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.CheckExist(username)

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

	_, err = repo.CheckExist(username)
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
		WithArgs(username).
		WillReturnRows(rows)

	item, err = repo.CheckExist(username)
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

func TestSearch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	search := "Mishail"
	search1 := "% " + search + "%"
	search2 := search + "%"

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "username", "email"})
	expect := models.PublicUsers{
		{Id: 3, Login: "name", Email: "d@gmail.com"},
		{Id: 3, Login: "name3", Email: "d@gmail.com"},
		{Id: 3, Login: "name10", Email: "d@gmail.com"},
	}

	for _, person := range expect {
		rows = rows.AddRow(person.Id, person.Login, person.Email)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(search1, search2).
		WillReturnRows(rows)

	item, err := repo.Search(search)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(*item, expect) {
		t.Errorf("results not match, want %v, have %v", expect[0], *item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(search1, search2).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.Search(search)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "title")

	mock.
		ExpectQuery("SELECT").
		WithArgs(search1, search2).
		WillReturnRows(rows)

	_, err = repo.Search(search)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
