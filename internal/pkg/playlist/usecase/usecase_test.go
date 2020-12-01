package usecase

import (
	"errors"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/playlist"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testFilms = models.FilmCards{models.FilmCard{
	Id: 1,
	Title: "titanic",
	MainGenre: "fantasy",
	SmallImg: "img",
	Year: 2010,
}}

var testPlaylist = models.Playlist{
	Id: 1,
	Title: "playlist",
	Films: &testFilms,
	UserId: 1,
}

var testSession = models.Session{
	Id:       "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}

var testPlaylists = models.Playlists{testPlaylist}

func TestGetPlayList(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := playlist.NewMockRepository(ctrl)
		m2 := film.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(1, nil)

		m0.
			EXPECT().GetList(gomock.Eq(1)).Return(&testPlaylists, nil)

		m2.EXPECT().FindFilmsByPlaylist(gomock.Eq(1)).Return(&testFilms, nil)

		usecase := NewPlaylistUseCase(m0, m2, m1)
		playlists, err := usecase.GetPlaylist(testSession.Id)
		assert.NoError(t, err)
		assert.Equal(t, *playlists, testPlaylists)
	})

	t.Run("playlist error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := playlist.NewMockRepository(ctrl)
		m2 := film.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(1, nil)

		m0.
			EXPECT().GetList(gomock.Eq(1)).Return(nil, errors.New("playlist error"))


		usecase := NewPlaylistUseCase(m0, m2, m1)
		playlists, err := usecase.GetPlaylist(testSession.Id)
		fmt.Println(playlists)
		assert.EqualError(t, err, "playlist error")
	})

	t.Run("film error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := playlist.NewMockRepository(ctrl)
		m2 := film.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(1, nil)

		m0.
			EXPECT().GetList(gomock.Eq(1)).Return(&testPlaylists, nil)

		m2.EXPECT().FindFilmsByPlaylist(gomock.Eq(1)).Return(nil, errors.New("film error"))

		usecase := NewPlaylistUseCase(m0, m2, m1)
		_, err := usecase.GetPlaylist(testSession.Id)
		assert.EqualError(t, err, "film error")
	})
}

func TestGetList(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := playlist.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m2 := film.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(1, nil)

		m0.
			EXPECT().GetList(gomock.Eq(1)).Return(&testPlaylists, nil)


		usecase := NewPlaylistUseCase(m0, m2, m1)
		playlists, err := usecase.GetList(testSession.Id)
		assert.NoError(t, err)
		assert.Equal(t, *playlists, testPlaylists)
	})

	t.Run("Error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := playlist.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m2 := film.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(1, nil)

		m0.
			EXPECT().GetList(gomock.Eq(1)).Return(nil, errors.New("playlist error"))


		usecase := NewPlaylistUseCase(m0, m2, m1)
		_, err := usecase.GetList(testSession.Id)
		assert.EqualError(t, err, "playlist error")
	})
}