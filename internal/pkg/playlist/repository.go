package playlist

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	Create(title string, userId int) error
	Delete(playlistId int) error
	Add(filmId int, playlistId int) error
	Remove(filmId int, playlistId int) error
	GetList(userId int) (*models.Playlists, error)
}

