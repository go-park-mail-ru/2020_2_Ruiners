package repository

import (
	"database/sql"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
)

type PlaylistRepository struct {
	db *sql.DB
}

func NewRPlaylistRepository(db *sql.DB) *PlaylistRepository {
	return &PlaylistRepository{
		db: db,
	}
}

func (p *PlaylistRepository) Create(title string, userId int) error {
	_, err := p.db.Exec("INSERT INTO playlist(title, user_id) VALUE(? , ?)", title, userId)
	if err != nil {
		return err
	}
	return nil
}

func (p *PlaylistRepository) Delete(playlistId int) error {
	fmt.Println("rep", playlistId)
	_, err := p.db.Exec("DELETE FROM playlist WHERE id = ?", playlistId)
	if err != nil {
		return err
	}
	return nil
}

func (p *PlaylistRepository) Add(filmId int, playlistId int) error {
	_, err := p.db.Exec("INSERT INTO playlist_film(playlist_id, film_id) VALUE(? , ?)", playlistId, filmId)
	if err != nil {
		return err
	}
	return nil
}

func (p *PlaylistRepository) Remove(filmId int, playlistId int) error {
	_, err := p.db.Exec("DELETE FROM playlist_film WHERE playlist_id = ? AND film_id = ?", playlistId, filmId)
	if err != nil {
		return err
	}
	return nil
}

func (p *PlaylistRepository) GetList(userId int) (*models.Playlists, error) {
	playlist := models.Playlist{}
	playlists := models.Playlists{}
	playlistQuery, err := p.db.Query("SELECT id, title, user_id FROM playlist WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer playlistQuery.Close()

	for playlistQuery.Next() {
		if playlistQuery.Scan(&playlist.Id, &playlist.Title, &playlist.UserId) != nil {
			return nil, errors.New("db error")
		}
		playlists = append(playlists, playlist)
	}
	return &playlists, nil
}
