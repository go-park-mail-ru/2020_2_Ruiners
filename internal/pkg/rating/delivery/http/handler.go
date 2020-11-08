package http

import (
	"encoding/json"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/rating"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RatingHandler struct {
	UseCase rating.UseCase
	logger *logrus.Logger
}

func (rh *RatingHandler) Rate() http.HandlerFunc {
	type Rate struct {
		FilmId int `'json:"filmId"'`
		Rating int `'json:"rating"'`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		l := Rate{}
		id, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(l)
		err = rh.UseCase.Rate(l.Rating, l.FilmId, id.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (rh *RatingHandler) AddReview() http.HandlerFunc {
	type AddReview struct {
		FilmId int `'json:"film_id"'`
		Body string `'json:"body"'`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		l := AddReview{}
		id, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = rh.UseCase.AddReview(l.Body, l.FilmId, id.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (rh *RatingHandler) ShowReviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmId := vars["film_id"]
	reviews, err := rh.UseCase.GetReviews(filmId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(&reviews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}