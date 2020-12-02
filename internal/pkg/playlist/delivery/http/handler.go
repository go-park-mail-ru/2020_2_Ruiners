package http

import (
	"encoding/json"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/playlist"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PlaylistHandler struct {
	UseCase   playlist.UseCase
	Logger    *logrus.Logger
	Sanitazer *bluemonday.Policy
}

func (ph *PlaylistHandler) CreatePlaylist() http.HandlerFunc {
	type CreatePlaylist struct {
		Title string `json:"title"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ph.Logger.Info("Add review")
		l := CreatePlaylist{}
		id, err := r.Cookie("session_id")
		if err != nil {
			ph.Logger.Error("No cookie delivery add review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			ph.Logger.Error("error with delivery add review json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		l.Title = ph.Sanitazer.Sanitize(l.Title)
		err = ph.UseCase.Create(l.Title, id.Value)
		if err != nil {
			ph.Logger.Error("error with usecase padd review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (ph *PlaylistHandler) DeletePlaylist() http.HandlerFunc {
	type DeletePlaylist struct {
		PlaylistId int `json:"playlist_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ph.Logger.Info("Add review")
		l := DeletePlaylist{}
		_, err := r.Cookie("session_id")
		if err != nil {
			ph.Logger.Error("No cookie delivery add review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			ph.Logger.Error("error with delivery add review json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("handler", l.PlaylistId)
		err = ph.UseCase.Delete(l.PlaylistId)
		if err != nil {
			ph.Logger.Error("error with usecase padd review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (ph *PlaylistHandler) AddPlaylist() http.HandlerFunc {
	type AddPlaylist struct {
		PlaylistId int `json:"playlist_id"`
		FilmId     int `json:"film_id"'`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ph.Logger.Info("Add review")
		l := AddPlaylist{}
		err := json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			ph.Logger.Error("error with delivery add review json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = ph.UseCase.Add(l.FilmId, l.PlaylistId)
		if err != nil {
			ph.Logger.Error("error with usecase padd review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (ph *PlaylistHandler) RemovePlaylist() http.HandlerFunc {
	type RemovePlaylist struct {
		PlaylistId int `json:"playlist_id"`
		FilmId     int `json:"film_id"'`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		ph.Logger.Info("Add review")
		l := RemovePlaylist{}
		err := json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			ph.Logger.Error("error with delivery add review json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = ph.UseCase.Remove(l.FilmId, l.PlaylistId)
		if err != nil {
			ph.Logger.Error("error with usecase padd review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (ph *PlaylistHandler) ShowList(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session_id")
	if err != nil {
		ph.Logger.Error("No cookie delivery show list")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	playlist, err := ph.UseCase.GetList(id.Value)
	if err != nil {
		ph.Logger.Error("error with usecase show playlist")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := easyjson.Marshal(playlist)
	if err != nil {
		ph.Logger.Error("error with delivery show reviews json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ph *PlaylistHandler) ShowPlaylist(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session_id")
	if err != nil {
		ph.Logger.Error("No cookie delivery show list")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	playlist, err := ph.UseCase.GetPlaylist(id.Value)
	if err != nil {
		ph.Logger.Error("error with usecase show playlist")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := easyjson.Marshal(playlist)
	if err != nil {
		ph.Logger.Error("error with delivery show reviews json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
