package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testFilm = models.Film{
	Id:          1231,
	Title:       "Erikdoter film",
	Rating:      10,
	SumVotes:    1000,
	Description: "Best film",
	MainGenre:   "Комедия",
	YoutubeLink: "youtube",
	BigImg:      "Big",
	SmallImg:    "Small",
	Year:        2020,
	Country:     "Russia",
}

var testFilmCard = models.FilmCard{
	Id:        1231,
	Title:     "ErikDoter film",
	MainGenre: "Комедия",
	SmallImg:  "Small",
	Year:      2020,
}

var testFilmCards = models.FilmCards{}

func TestFindById(t *testing.T) {
	t.Run("FIND", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().FindByLId(gomock.Eq(1)).Return(&testFilm, nil)

		usecase := NewFilmUseCase(m0)
		film, err := usecase.FindById("1")
		assert.NoError(t, err)
		assert.Equal(t, *film, testFilm)
	})
}

func TestFindByGenre(t *testing.T) {
	t.Run("FIND genre", func(t *testing.T) {
		testFilmCards = append(testFilmCards, testFilmCard)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().FindFilmsByGenre(gomock.Eq(testFilm.MainGenre)).Return(&testFilmCards, nil)

		usecase := NewFilmUseCase(m0)
		films, err := usecase.FilmsByGenre("comedy")
		assert.NoError(t, err)
		assert.Equal(t, *films, testFilmCards)
	})
}

func TestFindByPerson(t *testing.T) {
	t.Run("FIND person", func(t *testing.T) {
		testFilmCards = append(testFilmCards, testFilmCard)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().FindFilmsByPerson(gomock.Eq(1)).Return(&testFilmCards, nil)

		usecase := NewFilmUseCase(m0)
		films, err := usecase.FilmsByPerson("1")
		assert.NoError(t, err)
		assert.Equal(t, *films, testFilmCards)
	})
}
