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
		{Id: elemID, Title: "title", Rating: float64(6), SumVotes: 30, Description: "description", MainGenre: "Vasa Pupkin", YoutubeLink: "http://youtube", BigImg: "bigImg", SmallImg: "smallImg", Year: 2020, Country: "Russia"},
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
		{Id: 1, Title: "title1", MainGenre: genre, SmallImg: "smallImg", Year: 2005},
		{Id: 2, Title: "title2", MainGenre: genre, SmallImg: "smallImg", Year: 2020},
		{Id: 9, Title: "title9", MainGenre: genre, SmallImg: "smallImg", Year: 2000},
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
		{Id: 1, Title: "title1", MainGenre: "genre", SmallImg: "smallImg", Year: 2005, Rating: 0},
		{Id: 2, Title: "title2", MainGenre: "genre", SmallImg: "smallImg", Year: 2020, Rating: 0},
		{Id: 9, Title: "title9", MainGenre: "genre", SmallImg: "smallImg", Year: 2000, Rating: 0},
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
		{Id: 1, Title: "title1", MainGenre: "genre", SmallImg: "smallImg", Year: 2005, Rating: 0},
		{Id: 2, Title: "title2", MainGenre: "genre", SmallImg: "smallImg", Year: 2020, Rating: 0},
		{Id: 9, Title: "title9", MainGenre: "genre", SmallImg: "smallImg", Year: 2000, Rating: 0},
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

func TestSimilarFilms(t *testing.T) {
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
	rows2 := sqlmock.
		NewRows([]string{"id", "title", "mainGenre", "smallImg", "year"})
	rows3 := sqlmock.
		NewRows([]string{"id", "title", "mainGenre", "smallImg", "year"})
	expect := models.FilmCards{
		{Id: 1, Title: "title1", MainGenre: "genre", SmallImg: "smallImg", Year: 2005, Rating: 0},
		{Id: 2, Title: "title2", MainGenre: "genre", SmallImg: "smallImg", Year: 2020, Rating: 0},
		{Id: 9, Title: "title9", MainGenre: "genre", SmallImg: "smallImg", Year: 2000, Rating: 0},
		{Id: 1, Title: "title1", MainGenre: "genre", SmallImg: "smallImg", Year: 2005, Rating: 0},
		{Id: 2, Title: "title2", MainGenre: "genre", SmallImg: "smallImg", Year: 2020, Rating: 0},
		{Id: 9, Title: "title9", MainGenre: "genre", SmallImg: "smallImg", Year: 2000, Rating: 0},
		{Id: 1, Title: "title1", MainGenre: "genre", SmallImg: "smallImg", Year: 2005, Rating: 0},
		{Id: 2, Title: "title2", MainGenre: "genre", SmallImg: "smallImg", Year: 2020, Rating: 0},
		{Id: 9, Title: "title9", MainGenre: "genre", SmallImg: "smallImg", Year: 2000, Rating: 0},
	}

	for i := 0; i < 3; i++ {
		rows = rows.AddRow(expect[i].Id, expect[i].Title, expect[i].MainGenre, expect[i].SmallImg, expect[i].Year)
		rows2 = rows2.AddRow(expect[i+3].Id, expect[i+3].Title, expect[i+3].MainGenre, expect[i+3].SmallImg, expect[i+3].Year)
		rows3 = rows3.AddRow(expect[i+6].Id, expect[i+6].Title, expect[i+6].MainGenre, expect[i+6].SmallImg, expect[i+6].Year)

	}

	//for film := range (3) {
	//	rows = rows.AddRow(film.Id, film.Title, film.MainGenre, film.SmallImg, film.Year)
	//}

	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows)
	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows2)
	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows3)

	item, err := repo.SimilarFilms(playlistID)
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
		WithArgs(playlistID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.SimilarFilms(playlistID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}

	// query error 2
	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows)

	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.SimilarFilms(playlistID)

	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	// query error 3
	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows)

	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows)

	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.SimilarFilms(playlistID)

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

	_, err = repo.SimilarFilms(playlistID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	// row scan error 2
	rows = sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "title")

	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows2)
	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows)

	_, err = repo.SimilarFilms(playlistID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	// row scan  3
	rows = sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "title")

	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows2)

	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows2)

	mock.
		ExpectQuery("SELECT").
		WithArgs(playlistID).
		WillReturnRows(rows)

	_, err = repo.SimilarFilms(playlistID)
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

	repo := NewFilmRepository(db)

	search := "Mishail"
	search1 := "% " + search + "%"
	search2 := search + "%"

	// good query
	rows := sqlmock.
		NewRows([]string{"id", "title", "mainGenre", "smallImg", "year", "rating"})
	expect := models.FilmCards{
		{Id: 1, Title: "title1", MainGenre: "genre", SmallImg: "smallImg", Year: 2005, Rating: 0},
		{Id: 2, Title: "title2", MainGenre: "genre", SmallImg: "smallImg", Year: 2020, Rating: 0},
		{Id: 9, Title: "title9", MainGenre: "genre", SmallImg: "smallImg", Year: 2000, Rating: 0},
	}

	for _, film := range expect {
		rows = rows.AddRow(film.Id, film.Title, film.MainGenre, film.SmallImg, film.Year, film.Rating)
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
