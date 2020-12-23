package http

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strconv"

	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

var testFilm = models.Film{
	Id:          5,
	Title:       "string",
	Rating:      7,
	SumVotes:    5,
	Description: "string",
	MainGenre:   "string",
	YoutubeLink: "string",
	BigImg:      "string",
	SmallImg:    "string",
	Year:        2007,
	Country:     "string",
}

var testFilmCard = models.FilmCard{
	Id:        testFilm.Id,
	Title:     testFilm.Title,
	MainGenre: testFilm.MainGenre,
	SmallImg:  testFilm.SmallImg,
	Year:      testFilm.Year,
}

func TestFindById(t *testing.T) {

	t.Run("FindById-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			FindById(gomock.Eq(strconv.Itoa(testFilm.Id))).
			Return(&testFilm, nil)
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/film/5", nil)

		vars := map[string]string{
			"id": "5",
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmById)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "{\"id\":5,\"title\":\"string\",\"rating\":7,\"sum_votes\":5,\"description\":\"string\",\"main_genre\":\"string\",\"youtube_link\":\"string\",\"big_img\":\"string\",\"small_img\":\"string\",\"year\":2007,\"country\":\"string\"}")
	})

	t.Run("FindById-fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			FindById(gomock.Eq(strconv.Itoa(testFilm.Id))).
			Return(&testFilm, errors.New("error"))
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/film/5", nil)

		vars := map[string]string{
			"id": "5",
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmById)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestFindByGenre(t *testing.T) {

	var testFilmCards = models.FilmCards{}
	testFilmCards = append(testFilmCards, testFilmCard)

	t.Run("FindById-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			FilmsByGenre(gomock.Eq(testFilm.MainGenre)).
			Return(&testFilmCards, nil)
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/film/string", nil)

		vars := map[string]string{
			"genre": testFilm.MainGenre,
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmsByGenre)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "[{\"id\":5,\"title\":\"string\",\"main_genre\":\"string\",\"small_img\":\"string\",\"year\":2007,\"rating\":0}]")
	})

	t.Run("FindById-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			FilmsByGenre(gomock.Eq(testFilm.MainGenre)).
			Return(&testFilmCards, errors.New("error"))
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/film/string", nil)

		vars := map[string]string{
			"genre": testFilm.MainGenre,
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmsByGenre)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestFindByPerson(t *testing.T) {

	var testFilmCards = models.FilmCards{}
	testFilmCards = append(testFilmCards, testFilmCard)

	t.Run("FindById-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			FilmsByPerson(gomock.Eq("1")).
			Return(&testFilmCards, nil)
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/person_film/1", nil)

		vars := map[string]string{
			"id": "1",
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmsByPerson)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "[{\"id\":5,\"title\":\"string\",\"main_genre\":\"string\",\"small_img\":\"string\",\"year\":2007,\"rating\":0}]")
	})

	t.Run("FindById-fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			FilmsByPerson(gomock.Eq("1")).
			Return(&testFilmCards, errors.New("error"))
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/person_film/1", nil)

		vars := map[string]string{
			"id": "1",
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmsByPerson)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestSimilarFilms(t *testing.T) {

	var testFilmCards = models.FilmCards{}
	testFilmCards = append(testFilmCards, testFilmCard)

	t.Run("SimilarFilms-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			SimilarFilms(gomock.Eq("1")).
			Return(&testFilmCards, nil)
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/person_film/1", nil)

		vars := map[string]string{
			"id": "1",
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.SimilarFilms)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "[{\"id\":5,\"title\":\"string\",\"main_genre\":\"string\",\"small_img\":\"string\",\"year\":2007,\"rating\":0}]")
	})

	t.Run("SimilarFilms-fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			SimilarFilms(gomock.Eq("1")).
			Return(&testFilmCards, errors.New("error"))
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/person_film/1", nil)

		vars := map[string]string{
			"id": "1",
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.SimilarFilms)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestSearch(t *testing.T) {
	var testFilmCards = models.FilmCards{}
	testFilmCards = append(testFilmCards, testFilmCard)

	t.Run("Search-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Search(gomock.Eq("people")).
			Return(&testFilmCards, nil)
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/search?key=people", nil)

		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.Search)
		handler.ServeHTTP(rr, req)
		res, _ := json.Marshal(&testFilmCards)
		assert.Equal(t, rr.Body.String(), string(res))
	})

	t.Run("Search-fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Search(gomock.Eq("people")).
			Return(&testFilmCards, errors.New("error"))
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/search?key=people", nil)

		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.Search)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}
