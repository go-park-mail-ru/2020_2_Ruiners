package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"math"
	"strconv"
)

type FilmUseCase struct {
	FilmRepository film.Repository
}

func NewFilmUseCase(filmRepository film.Repository) *FilmUseCase {
	return &FilmUseCase{
		FilmRepository: filmRepository,
	}
}

func (uc *FilmUseCase) FindById(id string) (*models.Film, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	film, err := uc.FilmRepository.FindByLId(idInt)
	if err != nil {
		return nil, err
	}
	film.Rating = math.Round(film.Rating*100) / 100
	return film, nil
}

func (uc *FilmUseCase) FilmsByGenre(genre string) (*models.FilmCards, error) {
	films, err := uc.FilmRepository.FindFilmsByGenre(models.TranslateGenre[genre])
	if err != nil {
		return nil, err
	}
	return films, nil
}

func (uc *FilmUseCase) FilmsByPerson(id string) (*models.FilmCards, error) {
	idInt, err := strconv.Atoi(id)
	films, err := uc.FilmRepository.FindFilmsByPerson(idInt)
	if err != nil {
		return nil, err
	}
	return films, nil
}
