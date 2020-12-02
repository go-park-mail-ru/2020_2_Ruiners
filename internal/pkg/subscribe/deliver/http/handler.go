package http

import (
	"encoding/json"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/subscribe"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type SubscribeHandler struct {
	UseCase   subscribe.UseCase
	Logger    *logrus.Logger
	Sanitazer *bluemonday.Policy
}

func (sh *SubscribeHandler) Subscribe() http.HandlerFunc {
	type Subscribe struct {
		UserId int `json:"user_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		sh.Logger.Info("Add review")
		l := Subscribe{}
		id, err := r.Cookie("session_id")
		if err != nil {
			sh.Logger.Error("No cookie delivery add review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			sh.Logger.Error("error with delivery add review json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = sh.UseCase.Create(l.UserId, id.Value)
		if err != nil {
			sh.Logger.Error("error with usecase padd review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (sh *SubscribeHandler) UnSubscribe() http.HandlerFunc {
	type UnSubscribe struct {
		UserId int `json:"user_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		sh.Logger.Info("Add review")
		l := UnSubscribe{}
		id, err := r.Cookie("session_id")
		if err != nil {
			sh.Logger.Error("No cookie delivery add review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			sh.Logger.Error("error with delivery add review json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = sh.UseCase.Delete(l.UserId, id.Value)
		if err != nil {
			sh.Logger.Error("error with usecase padd review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (sh *SubscribeHandler) ShowFeed(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session_id")
	if err != nil {
		sh.Logger.Error("No cookie delivery show list")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	feed, err := sh.UseCase.GetFeed(id.Value)
	if err != nil {
		sh.Logger.Error("error with usecase show playlist")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := easyjson.Marshal(feed)
	if err != nil {
		sh.Logger.Error("error with delivery show reviews json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (sh *SubscribeHandler) ShowAuthors(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session_id")
	if err != nil {
		sh.Logger.Error("No cookie delivery show list")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authors, err := sh.UseCase.GetAuthors(id.Value)
	if err != nil {
		sh.Logger.Error("error with usecase show playlist")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := easyjson.Marshal(authors)
	if err != nil {
		sh.Logger.Error("error with delivery show reviews json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (sh *SubscribeHandler) Check() http.HandlerFunc {
	type User struct {
		UserId int `json:"user_id"`
	}
	type Subscribe struct {
		Check bool `json:"check"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		sh.Logger.Info("Add review")
		id, err := r.Cookie("session_id")
		if err != nil {
			sh.Logger.Error("No cookie delivery add review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		vars := mux.Vars(r)
		user_id := vars["user_id"]
		userId, _ := strconv.Atoi(user_id)
		sub := Subscribe{}
		sub.Check, err = sh.UseCase.Check(id.Value, userId)
		if err != nil {
			sh.Logger.Error("error with usecase padd review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res, err := json.Marshal(&sub)
		if err != nil {
			sh.Logger.Error("error with delivery show reviews json-marshal")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
