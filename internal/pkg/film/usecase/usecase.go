package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
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
	id_int, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	film, err := uc.FilmRepository.FindByLId(id_int)
	if err != nil {
		return nil, err
	}
	return film, nil
}
