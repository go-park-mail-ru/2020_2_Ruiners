package usecase

import (
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/client"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/playlist"
)

type PlaylistUseCase struct {
	PlaylistRepository playlist.Repository
	FilmRepository     film.Repository
	RpcSession         client.ISessionClient
}

func NewPlaylistUseCase(playlistRepository playlist.Repository, filmRepository film.Repository, rpcSession client.ISessionClient) *PlaylistUseCase {
	return &PlaylistUseCase{
		PlaylistRepository: playlistRepository,
		FilmRepository:     filmRepository,
		RpcSession:         rpcSession,
	}
}

func (uc *PlaylistUseCase) Create(title string, session string) error {
	userId, err := uc.RpcSession.GetUserIdBySession(session)
	if err != nil {
		return err
	}
	err = uc.PlaylistRepository.Create(title, userId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *PlaylistUseCase) Add(filmId int, playlistId int) error {
	err := uc.PlaylistRepository.Add(filmId, playlistId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *PlaylistUseCase) Delete(playlistId int) error {
	fmt.Println("use", playlistId)
	err := uc.PlaylistRepository.Delete(playlistId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *PlaylistUseCase) Remove(filmId int, playlistId int) error {
	err := uc.PlaylistRepository.Remove(filmId, playlistId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *PlaylistUseCase) GetList(session string) (*models.Playlists, error) {
	userId, err := uc.RpcSession.GetUserIdBySession(session)
	if err != nil {
		return nil, err
	}
	playlist, err := uc.PlaylistRepository.GetList(userId)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (uc *PlaylistUseCase) GetPlaylist(session string) (*models.Playlists, error) {
	res := models.Playlists{}
	userId, err := uc.RpcSession.GetUserIdBySession(session)
	if err != nil {
		return nil, err
	}
	playlist, err := uc.PlaylistRepository.GetList(userId)
	if err != nil {
		return nil, err
	}
	for _, r := range *playlist {
		f, err := uc.FilmRepository.FindFilmsByPlaylist(r.Id)
		if err != nil {
			return nil, err
		}
		r.Films = f
		res = append(res, r)
	}
	return &res, nil
}
