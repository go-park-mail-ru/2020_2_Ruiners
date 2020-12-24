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

	repo := NewPersonRepository(db)

	var elemID = 1
	var emptyArray []int

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "name", "image", "born_date", "born_place"})
	expect := []models.Person{
		{
			Id:       elemID,
			Name:      "name",
			Image:     "image",
			BornDate:  "30.12.2000",
			BornPlace: "Moskow",
			FilmsId:   emptyArray,
		},
	}

	for _, person := range expect {
		rows = rows.AddRow(person.Id, person.Name, person.Image, person.BornDate, person.BornPlace)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repo.FindById(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(*item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], *item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(elemID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.FindById(elemID)

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
		AddRow(1, "name")

	mock.
		ExpectQuery("SELECT").
		WithArgs(elemID).
		WillReturnRows(rows)

	_, err = repo.FindById(elemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestFindByFilmIdAndRole(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewPersonRepository(db)

	var filmID = 3
	var role = "Mishail"

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "name"})
	expect := models.FilmPersons{
		{Id: 3, Name: "name",  Image: ""},
		{Id: 3, Name: "name3",  Image: ""},
		{Id: 3, Name: "name10",  Image: ""},
	}

	for _, person := range expect {
		rows = rows.AddRow(person.Id, person.Name)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(filmID, role).
		WillReturnRows(rows)

	item, err := repo.FindByFilmIdAndRole(filmID, role)
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
		WithArgs(filmID, role).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.FindByFilmIdAndRole(filmID, role)

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
		WithArgs(filmID, role).
		WillReturnRows(rows)

	_, err = repo.FindByFilmIdAndRole(filmID, role)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestFindFilmsIdByPersonId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewPersonRepository(db)

	var persID = 3

	// good query
	rows := sqlmock.
		NewRows([]string{"film_id"})
	expect := []int{
		1,
		3,
		10,
	}

	for _, film := range expect {
		rows = rows.AddRow(film)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(persID).
		WillReturnRows(rows)

	item, err := repo.FindFilmsIdByPersonId(persID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT").
		WithArgs(persID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.FindFilmsIdByPersonId(persID)

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
		WithArgs(persID).
		WillReturnRows(rows)

	_, err = repo.FindFilmsIdByPersonId(persID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestSearch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewPersonRepository(db)

	search := "Mishail"
	search1 := "% " + search + "%"
	search2 := search + "%"

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "name", "img"})
	expect := models.FilmPersons{
		{Id: 3, Name: "name",  Image: ""},
		{Id: 3, Name: "name3",  Image: ""},
		{Id: 3, Name: "name10",  Image: ""},
	}

	for _, person := range expect {
		rows = rows.AddRow(person.Id, person.Name, person.Image)
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
