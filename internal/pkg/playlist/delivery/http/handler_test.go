package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/playlist"
	"github.com/golang/mock/gomock"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var testSession = models.Session{
	Id:       "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}

var testPlaylists = models.Playlists{}

func TestShowList(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetList(gomock.Eq(testSession.Id)).
			Return(&testPlaylists, nil)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		req, err := http.NewRequest("GET", "/authors", nil)
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(playlistHandler.ShowList)
		handler.ServeHTTP(rr, req)
		res, _ := json.Marshal(testPlaylists)
		assert.Equal(t, rr.Body.String(), string(res))
	})

	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetList(gomock.Eq(testSession.Id)).
			Return(&testPlaylists, errors.New("error"))

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		req, err := http.NewRequest("GET", "/authors", nil)
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(playlistHandler.ShowList)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})

	t.Run("No Cookie", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := playlist.NewMockUseCase(ctrl)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		req, err := http.NewRequest("GET", "/authors", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(playlistHandler.ShowList)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestShowPlaylist(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetPlaylist(gomock.Eq(testSession.Id)).
			Return(&testPlaylists, nil)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		req, err := http.NewRequest("GET", "/authors", nil)
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(playlistHandler.ShowPlaylist)
		handler.ServeHTTP(rr, req)
		//res, _ := json.Marshal(testPlaylists)
		//assert.Equal(t, rr.Body.String(), string(res))
	})

	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetPlaylist(gomock.Eq(testSession.Id)).
			Return(&testPlaylists, errors.New("error"))

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		req, err := http.NewRequest("GET", "/authors", nil)
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(playlistHandler.ShowPlaylist)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})

	t.Run("No Cookie", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := playlist.NewMockUseCase(ctrl)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		req, err := http.NewRequest("GET", "/authors", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(playlistHandler.ShowPlaylist)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestRemovePlaylist(t *testing.T) {
	type RemovePlaylist struct {
		PlaylistId int `json:"playlist_id"`
		FilmId     int `json:"film_id"'`
	}
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			PlaylistId: 6,
			FilmId:     7,
		}

		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Remove(gomock.Eq(elem.FilmId), gomock.Eq(elem.PlaylistId)).
			Return(nil)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.RemovePlaylist()
		handler.ServeHTTP(rr, req)
	})

	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			PlaylistId: 6,
			FilmId:     7,
		}

		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Remove(gomock.Eq(elem.FilmId), gomock.Eq(elem.PlaylistId)).
			Return(errors.New("error"))

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.RemovePlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
	type RemovePlaylistBad struct {
		PlaylistId string `json:"playlist_id"`
		FilmId     string `json:"film_id"'`
	}
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylistBad{
			PlaylistId: "gjx",
			FilmId:     "ic",
		}

		m := playlist.NewMockUseCase(ctrl)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.RemovePlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestAddPlaylist(t *testing.T) {
	type RemovePlaylist struct {
		PlaylistId int `json:"playlist_id"`
		FilmId     int `json:"film_id"'`
	}
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			PlaylistId: 6,
			FilmId:     7,
		}

		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Add(gomock.Eq(elem.FilmId), gomock.Eq(elem.PlaylistId)).
			Return(nil)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.AddPlaylist()
		handler.ServeHTTP(rr, req)
	})

	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			PlaylistId: 6,
			FilmId:     7,
		}

		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Add(gomock.Eq(elem.FilmId), gomock.Eq(elem.PlaylistId)).
			Return(errors.New("error"))

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.AddPlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
	type RemovePlaylistBad struct {
		PlaylistId string `json:"playlist_id"`
		FilmId     string `json:"film_id"'`
	}
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylistBad{
			PlaylistId: "gjx",
			FilmId:     "ic",
		}

		m := playlist.NewMockUseCase(ctrl)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.AddPlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestDeletePlaylist(t *testing.T) {
	type RemovePlaylist struct {
		PlaylistId int `json:"playlist_id"`
	}
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			PlaylistId: 6,
		}

		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Delete(gomock.Eq(elem.PlaylistId)).
			Return(nil)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.DeletePlaylist()
		handler.ServeHTTP(rr, req)
	})

	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			PlaylistId: 6,
		}

		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Delete(gomock.Eq(elem.PlaylistId)).
			Return(errors.New("error"))

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.DeletePlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})

	t.Run("No Cookie", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			PlaylistId: 7,
		}

		m := playlist.NewMockUseCase(ctrl)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.DeletePlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})

	type RemovePlaylistBad struct {
		PlaylistId string `json:"playlist_id"`
	}
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylistBad{
			PlaylistId: "gjx",
		}

		m := playlist.NewMockUseCase(ctrl)

		playlistHandler := PlaylistHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.DeletePlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestCreatePlaylist(t *testing.T) {
	type RemovePlaylist struct {
		Title string `json:"title"`
	}
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			Title: "6",
		}

		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Create(gomock.Eq(elem.Title), gomock.Eq(testSession.Id)).
			Return(nil)

		playlistHandler := PlaylistHandler{
			UseCase:   m,
			Logger:    logrus.New(),
			Sanitazer: bluemonday.UGCPolicy(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.CreatePlaylist()
		handler.ServeHTTP(rr, req)
	})

	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			Title: "6",
		}

		m := playlist.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Create(gomock.Eq(elem.Title), gomock.Eq(testSession.Id)).
			Return(errors.New("error"))

		playlistHandler := PlaylistHandler{
			UseCase:   m,
			Logger:    logrus.New(),
			Sanitazer: bluemonday.UGCPolicy(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.CreatePlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})

	t.Run("No Cookie", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylist{
			Title: "j",
		}

		m := playlist.NewMockUseCase(ctrl)

		playlistHandler := PlaylistHandler{
			UseCase:   m,
			Logger:    logrus.New(),
			Sanitazer: bluemonday.UGCPolicy(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.CreatePlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})

	type RemovePlaylistBad struct {
		Title int `json:"title"`
	}
	t.Run("fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		var elem = RemovePlaylistBad{
			Title: 8,
		}

		m := playlist.NewMockUseCase(ctrl)

		playlistHandler := PlaylistHandler{
			UseCase:   m,
			Logger:    logrus.New(),
			Sanitazer: bluemonday.UGCPolicy(),
		}

		bytesRepresentation, err := json.Marshal(elem)
		req, err := http.NewRequest("GET", "/authors", bytes.NewBuffer(bytesRepresentation))
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := playlistHandler.CreatePlaylist()
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}
