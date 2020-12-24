package usecase

import (
	"errors"
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

var testFilmCard2 = models.FilmCard{
	Id:        2,
	Title:     "ErikDoter film3123",
	MainGenre: "Комедия213",
	SmallImg:  "Small",
	Year:      2020,
}

var testFilmCard3 = models.FilmCard{
	Id:        3,
	Title:     "ErikDoter film12",
	MainGenre: "Комедия3213",
	SmallImg:  "Small",
	Year:      2020,
}

var testFilmCard4 = models.FilmCard{
	Id:        4,
	Title:     "ErikDoter filmdcsdsc",
	MainGenre: "Комедияxsaxsax",
	SmallImg:  "Small",
	Year:      2020,
}

var testFilmCard5 = models.FilmCard{
	Id:        5,
	Title:     "ErikDoter filmdcsdscdsx",
	MainGenre: "Комедияxcdcdsccdsxsax",
	SmallImg:  "Small",
	Year:      2020,
}

var testFilmCard6 = models.FilmCard{
	Id:        6,
	Title:     "ErikDoter filmdcsdscdswdwqdx",
	MainGenre: "Комедияxcdcdsccdsxsadwqswqswx",
	SmallImg:  "Small",
	Year:      2020,
}

var testFilmCards = models.FilmCards{}

var testSimilar = models.FilmCards{testFilmCard}
var testSimilar2 = models.FilmCards{testFilmCard, testFilmCard}
var testSimilar3 = models.FilmCards{testFilmCard, testFilmCard2, testFilmCard3, testFilmCard4, testFilmCard5, testFilmCard6}
var testSimilar4 = models.FilmCards{testFilmCard, testFilmCard2, testFilmCard3, testFilmCard4, testFilmCard5}

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

	t.Run("Bad in", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)

		usecase := NewFilmUseCase(m0)
		_, err := usecase.FindById("lgv")
		assert.Error(t, err)
	})

	t.Run("error rep", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().FindByLId(gomock.Eq(1)).Return(&testFilm, errors.New("error rep"))

		usecase := NewFilmUseCase(m0)
		_, err := usecase.FindById("1")
		assert.EqualError(t, err, "error rep")
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

	t.Run("Repo error", func(t *testing.T) {
		testFilmCards = append(testFilmCards, testFilmCard)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().FindFilmsByGenre(gomock.Eq(testFilm.MainGenre)).Return(&testFilmCards, errors.New("error"))

		usecase := NewFilmUseCase(m0)
		_, err := usecase.FilmsByGenre("comedy")
		assert.EqualError(t, err, "error")
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

	t.Run("Repo error", func(t *testing.T) {
		testFilmCards = append(testFilmCards, testFilmCard)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().FindFilmsByPerson(gomock.Eq(1)).Return(&testFilmCards, errors.New("error"))

		usecase := NewFilmUseCase(m0)
		_, err := usecase.FilmsByPerson("1")
		assert.EqualError(t, err, "error")
	})
}

func TestSearch(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		testFilmCards = append(testFilmCards, testFilmCard)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().Search(gomock.Eq("Erik")).Return(&testFilmCards, nil)

		usecase := NewFilmUseCase(m0)
		films, err := usecase.Search("Erik")
		assert.NoError(t, err)
		assert.Equal(t, *films, testFilmCards)
	})

	t.Run("error", func(t *testing.T) {
		testFilmCards = append(testFilmCards, testFilmCard)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().Search(gomock.Eq("Erik")).Return(&testFilmCards, errors.New("error"))

		usecase := NewFilmUseCase(m0)
		_, err := usecase.Search("Erik")
		assert.EqualError(t, err, "error")
	})
}

func TestSimilar(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		testFilmCards = append(testFilmCards, testFilmCard)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().SimilarFilms(gomock.Eq(1)).Return(&testSimilar, nil)

		usecase := NewFilmUseCase(m0)
		films, err := usecase.SimilarFilms("1")
		assert.NoError(t, err)
		assert.Equal(t, *films, testSimilar)
	})

	t.Run("Bad in", func(t *testing.T) {
		testFilmCards = append(testFilmCards, testFilmCard)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)

		usecase := NewFilmUseCase(m0)
		_, err := usecase.SimilarFilms("khg")
		assert.Error(t, err)
	})

	t.Run("Same films", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().SimilarFilms(gomock.Eq(1)).Return(&testSimilar2, nil)

		usecase := NewFilmUseCase(m0)
		films, err := usecase.SimilarFilms("1")
		assert.NoError(t, err)
		assert.Equal(t, *films, testSimilar)
	})

	t.Run("higher 4", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().SimilarFilms(gomock.Eq(1)).Return(&testSimilar3, nil)

		usecase := NewFilmUseCase(m0)
		films, err := usecase.SimilarFilms("1")
		assert.NoError(t, err)
		assert.Equal(t, *films, testSimilar4)
	})

	t.Run("rep error", func(t *testing.T) {
		testFilmCards = append(testFilmCards, testFilmCard)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := film.NewMockRepository(ctrl)
		m0.
			EXPECT().SimilarFilms(gomock.Eq(1)).Return(&testSimilar, errors.New("error"))

		usecase := NewFilmUseCase(m0)
		_, err := usecase.SimilarFilms("1")
		assert.EqualError(t, err, "error")
	})
}
