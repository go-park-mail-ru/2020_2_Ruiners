package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFindByLId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFilmRepository(db)

	var elemID = 1

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "title", "rating", "sumVotes", "description", "mainGenre", "youtubeLink", "bigImg", "smallImg", "year", "country"})
	expect := []models.Film{
		{elemID, "title", float64(6), 30, "description", "Vasa Pupkin", "http://youtube", "bigImg", "smallImg", 2020, "Russia"},
	}

	for _, film := range expect {
		rows = rows.AddRow(film.Id, film.Title, film.Rating, film.SumVotes, film.Description, film.MainGenre, film.YoutubeLink, film.BigImg, film.SmallImg, film.Year, film.Country)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := repo.FindByLId(elemID)
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

	_, err = repo.FindByLId(elemID)

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
		WithArgs(elemID).
		WillReturnRows(rows)

	_, err = repo.FindByLId(elemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestFindFilmsByGenre(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFilmRepository(db)

	var genre = "Mishail"

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "title", "mainGenre", "smallImg", "year"})
	expect := models.FilmCards{
		{1, "title1", genre, "smallImg", 2005},
		{2, "title2", genre, "smallImg", 2020},
		{9, "title9", genre, "smallImg", 2000},
	}

	for _, film := range expect {
		rows = rows.AddRow(film.Id, film.Title, film.MainGenre, film.SmallImg, film.Year)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(genre).
		WillReturnRows(rows)

	item, err := repo.FindFilmsByGenre(genre)
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
		WithArgs(genre).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.FindFilmsByGenre(genre)

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
		WithArgs(genre).
		WillReturnRows(rows)

	_, err = repo.FindFilmsByGenre(genre)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestFindFilmsByPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFilmRepository(db)

	var persID = 1

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "title", "mainGenre", "smallImg", "year"})
	expect := models.FilmCards{
		{1, "title1", "genre", "smallImg", 2005},
		{2, "title2", "genre", "smallImg", 2020},
		{9, "title9", "genre", "smallImg", 2000},
	}

	for _, film := range expect {
		rows = rows.AddRow(film.Id, film.Title, film.MainGenre, film.SmallImg, film.Year)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(persID).
		WillReturnRows(rows)

	item, err := repo.FindFilmsByPerson(persID)
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
		WithArgs(persID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.FindFilmsByPerson(persID)

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
		WithArgs(persID).
		WillReturnRows(rows)

	_, err = repo.FindFilmsByPerson(persID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestFindFilmsByPlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := NewFilmRepository(db)

	var playlistID = 1

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "title", "mainGenre", "smallImg", "year"})
	expect := models.FilmCards{
		{1, "title1", "genre", "smallImg", 2005},
		{2, "title2", "genre", "smallImg", 2020},
		{9, "title9", "genre", "smallImg", 2000},
	}

	for _, film := range expect {
		rows = rows.AddRow(film.Id, film.Title, film.MainGenre, film.SmallImg, film.Year)
	}

	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows)

	item, err := repo.FindFilmsByPlaylist(playlistID)
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
		WithArgs(playlistID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.FindFilmsByPlaylist(playlistID)

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
		WithArgs(playlistID).
		WillReturnRows(rows)

	_, err = repo.FindFilmsByPlaylist(playlistID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}
