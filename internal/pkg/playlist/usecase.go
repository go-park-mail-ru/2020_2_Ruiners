package playlist

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type UseCase interface {
	Create(title string, session string) error
	Delete(playlistId int) error
	Add(filmId int, playlistId int) error
	Remove(filmId int, playlistId int) error
	GetList(session string) (*models.Playlists, error)
	GetPlaylist(session string) (*models.Playlists, error)
}