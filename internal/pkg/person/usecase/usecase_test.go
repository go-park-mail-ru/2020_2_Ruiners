package usecase

import (
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/person"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPerson = models.Person{
	Id:        1,
	Name:      "Erik Nabiev",
	Image:     "Erik",
	BornDate:  "12.12.2000",
	BornPlace: "Moscow",
	FilmsId:   make([]int, 1),
}

var testPersonFilm = models.FilmPerson{
	Id:   1,
	Name: "Erik Nabiev",
}

var testPersonFilms = models.FilmPersons{}

var array = make([]int, 1)

func TestGetPerson(t *testing.T) {
	t.Run("GetPerson-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)
		m0.
			EXPECT().
			FindById(gomock.Eq(1)).
			Return(&testPerson, nil)

		m0.
			EXPECT().
			FindFilmsIdByPersonId(gomock.Eq(1)).
			Return(array, nil)

		usecase := NewPersonUseCase(m0)
		person, err := usecase.GetPerson("1")
		assert.NoError(t, err)
		assert.Equal(t, *person, testPerson)
	})

	t.Run("GetPerson-FindFilmsIdByPersonId", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)
		m0.
			EXPECT().
			FindById(gomock.Eq(1)).
			Return(&testPerson, nil)

		m0.
			EXPECT().
			FindFilmsIdByPersonId(gomock.Eq(1)).
			Return(array, errors.New("error"))

		usecase := NewPersonUseCase(m0)
		_, err := usecase.GetPerson("1")
		assert.Error(t, err)
	})

	t.Run("GetPerson-FindFilmsIdByPersonId", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)
		m0.
			EXPECT().
			FindById(gomock.Eq(1)).
			Return(&testPerson, errors.New("error"))

		usecase := NewPersonUseCase(m0)
		_, err := usecase.GetPerson("1")
		assert.Error(t, err)
	})

	t.Run("GetPerson-Bad in", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)

		usecase := NewPersonUseCase(m0)
		_, err := usecase.GetPerson("not num")
		assert.Error(t, err)
	})
}

func TestGetPersonsByFilm(t *testing.T) {
	t.Run("GetPersonsByFilm-OK", func(t *testing.T) {
		testPersonFilms = append(testPersonFilms, testPersonFilm)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)
		m0.
			EXPECT().
			FindByFilmIdAndRole(gomock.Eq(1), gomock.Eq("actor")).
			Return(&testPersonFilms, nil)

		usecase := NewPersonUseCase(m0)
		persons, err := usecase.GetPersonsByFilm("1", "actor")
		assert.NoError(t, err)
		assert.Equal(t, *persons, testPersonFilms)
	})

	t.Run("GetPersonsByFilm-FindByFilmIdAndRole", func(t *testing.T) {
		testPersonFilms = append(testPersonFilms, testPersonFilm)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)
		m0.
			EXPECT().
			FindByFilmIdAndRole(gomock.Eq(1), gomock.Eq("actor")).
			Return(&testPersonFilms, errors.New("error"))

		usecase := NewPersonUseCase(m0)
		_, err := usecase.GetPersonsByFilm("1", "actor")
		assert.Error(t, err)
	})

	t.Run("GetPersonsByFilm-Bad in", func(t *testing.T) {
		testPersonFilms = append(testPersonFilms, testPersonFilm)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)

		usecase := NewPersonUseCase(m0)
		_, err := usecase.GetPersonsByFilm("not num", "actor")
		assert.Error(t, err)
	})
}

func TestSearch(t *testing.T) {
	search := "film"
	t.Run("Search-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)

		m0.
			EXPECT().
			Search(gomock.Eq(search)).
			Return(&testPersonFilms, nil)

		useCase := NewPersonUseCase(m0)

		user, err := useCase.Search(search)
		assert.NoError(t, err)
		assert.Equal(t, *user, testPersonFilms)
	})

	t.Run("Search-Search", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)

		m0.
			EXPECT().
			Search(gomock.Eq(search)).
			Return(&testPersonFilms, errors.New("error"))

		useCase := NewPersonUseCase(m0)

		_, err := useCase.Search(search)
		assert.Error(t, err)
	})
}
