package http

import (
	"encoding/json"
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/playlist"
	"github.com/golang/mock/gomock"
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
