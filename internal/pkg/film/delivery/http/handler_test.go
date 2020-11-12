package http

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"

	//"fmt"
	"strconv"

	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	//"errors"
	//"encoding/json"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	//"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	//"time"
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

		// CHANGE THIS LINE!!!
		req = mux.SetURLVars(req, vars)
		if err != nil {
			t.Fatal(err)
		}
		//req.AddCookie(&http.Cookie{
		//	Name:    "session_id",
		//	Value:   testSession.Id,
		//	Expires: time.Now().Add(10 * time.Hour),
		//})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmById)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "{\"id\":5,\"title\":\"string\",\"Rating\":7,\"SumVotes\":5,\"Description\":\"string\",\"MainGenre\":\"string\",\"YoutubeLink\":\"string\",\"BigImg\":\"string\",\"SmallImg\":\"string\",\"year\":2007,\"country\":\"string\"}")
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

		// CHANGE THIS LINE!!!
		req = mux.SetURLVars(req, vars)
		if err != nil {
			t.Fatal(err)
		}
		//req.AddCookie(&http.Cookie{
		//	Name:    "session_id",
		//	Value:   testSession.Id,
		//	Expires: time.Now().Add(10 * time.Hour),
		//})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmsByGenre)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "[{\"id\":5,\"title\":\"string\",\"MainGenre\":\"string\",\"SmallImg\":\"string\",\"year\":2007}]")
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

		// CHANGE THIS LINE!!!
		req = mux.SetURLVars(req, vars)
		if err != nil {
			t.Fatal(err)
		}
		//req.AddCookie(&http.Cookie{
		//	Name:    "session_id",
		//	Value:   testSession.Id,
		//	Expires: time.Now().Add(10 * time.Hour),
		//})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmsByPerson)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "[{\"id\":5,\"title\":\"string\",\"MainGenre\":\"string\",\"SmallImg\":\"string\",\"year\":2007}]")
	})
}
