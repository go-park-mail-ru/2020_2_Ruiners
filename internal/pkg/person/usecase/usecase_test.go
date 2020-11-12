package usecase

import (
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

func TestFindById(t *testing.T) {
	t.Run("FIND", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)
		m0.
			EXPECT().FindById(gomock.Eq(1)).Return(&testPerson, nil)
		m0.EXPECT().FindFilmsIdByPersonId(gomock.Eq(1)).Return(array, nil)

		usecase := NewPersonUseCase(m0)
		person, err := usecase.GetPerson("1")
		assert.NoError(t, err)
		assert.Equal(t, *person, testPerson)
	})
}

func TestFindByPerson(t *testing.T) {
	t.Run("FIND", func(t *testing.T) {
		testPersonFilms = append(testPersonFilms, testPersonFilm)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := person.NewMockRepository(ctrl)
		m0.
			EXPECT().FindByFilmIdAndRole(gomock.Eq(1), gomock.Eq("actor")).Return(&testPersonFilms, nil)

		usecase := NewPersonUseCase(m0)
		persons, err := usecase.GetPersonsByFilm("1", "actor")
		assert.NoError(t, err)
		assert.Equal(t, *persons, testPersonFilms)
	})
}
