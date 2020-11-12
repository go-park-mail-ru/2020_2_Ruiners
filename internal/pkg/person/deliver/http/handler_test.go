package http

import (
	"encoding/json"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/person"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)




func TestPersonById(t *testing.T) {

	t.Run("FindById-OK", func(t *testing.T) {
		var testPerson = models.Person{
			Id: 1,
			Name: "Erik",
			Image: "image",
			BornDate: "12.12.2000",
			BornPlace: "Moscow",
			FilmsId: make([]int, 1),
		}
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := person.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetPerson(gomock.Eq("1")).
			Return(&testPerson, nil)
		personHandler := PersonHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/person/1", nil)

		vars := map[string]string{
			"id": "1",
		}

		req = mux.SetURLVars(req, vars)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(personHandler.PersonById)
		handler.ServeHTTP(rr, req)
		res, _ := json.Marshal(&testPerson)
		assert.Equal(t, rr.Body.String(), string(res))
	})
}

func TestPersonsByFilm(t *testing.T) {

	t.Run("FindByFilm-OK", func(t *testing.T) {
		var testFilmPerson = models.FilmPerson{
			Id: 1,
			Name: "Erik",
		}
		var testFilmPersons = models.FilmPersons{}
		testFilmPersons = append(testFilmPersons, testFilmPerson)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := person.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetPersonsByFilm(gomock.Eq("1"), gomock.Eq("actor")).
			Return(&testFilmPersons, nil)
		personHandler := PersonHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/actor/1", nil)

		vars := map[string]string{
			"film_id": "1",
			"role": "actor",
		}

		req = mux.SetURLVars(req, vars)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(personHandler.PersonsByFilm)
		handler.ServeHTTP(rr, req)
		res, _ := json.Marshal(&testFilmPersons)
		assert.Equal(t, rr.Body.String(), string(res))
	})
}