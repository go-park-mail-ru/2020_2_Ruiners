package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"math"
	"sort"
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

func (uc *FilmUseCase) SimilarFilms(id string) (*models.FilmCards, error) {
	type FilmCardCount struct {
		FilmCard models.FilmCard
		Count    int
	}
	var fcCount []FilmCardCount
	cards := models.FilmCards{}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	films, err := uc.FilmRepository.SimilarFilms(idInt)
	if err != nil {
		return nil, err
	}
	for _, x := range *films {
		fcCount = append(fcCount, FilmCardCount{FilmCard: x, Count: 1})
	}
	for i := 0; i < len(fcCount)-1; i++ {
		counter := 0
		var deleteSlice []int
		for j := i + 1; j < len(fcCount); j++ {
			if fcCount[i].FilmCard == fcCount[j].FilmCard {
				fcCount[i].Count++
				deleteSlice = append(deleteSlice, j)
			}
		}
		for _, x := range deleteSlice {
			fcCount = append(fcCount[:x-counter], fcCount[x+1-counter:]...)
			counter++
		}
	}
	sort.SliceStable(fcCount, func(i, j int) bool {
		return fcCount[i].Count > fcCount[j].Count
	})
	for i, x := range fcCount {
		cards = append(cards, x.FilmCard)
		if i == 4 {
			break
		}
	}
	return &cards, nil
}

func (uc *FilmUseCase) Search(search string) (*models.FilmCards, error) {
	films, err := uc.FilmRepository.Search(search)
	if err != nil {
		return nil, err
	}
	return films, nil
}
