package http

import (
	"encoding/json"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/rating"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RatingHandler struct {
	UseCase   rating.UseCase
	Logger    *logrus.Logger
	Sanitazer *bluemonday.Policy
}

func (rh *RatingHandler) Rate() http.HandlerFunc {
	type Rate struct {
		FilmId int `'json:"filmId"'`
		Rating int `'json:"rating"'`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		rh.Logger.Info("Rate")
		l := Rate{}
		id, err := r.Cookie("session_id")
		if err != nil {
			rh.Logger.Error("No cookie rate delivery")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			rh.Logger.Error("Error with delivery rate json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = rh.UseCase.Rate(l.Rating, l.FilmId, id.Value)
		if err != nil {
			rh.Logger.Error("error with usecase rate")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (rh *RatingHandler) AddReview() http.HandlerFunc {
	type AddReview struct {
		FilmId int    `'json:"film_id"'`
		Body   string `'json:"body"'`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		rh.Logger.Info("Add review")
		l := AddReview{}
		id, err := r.Cookie("session_id")
		if err != nil {
			rh.Logger.Error("No cookie delivery add review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			rh.Logger.Error("error with delivery add review json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		l.Body = rh.Sanitazer.Sanitize(l.Body)
		err = rh.UseCase.AddReview(l.Body, l.FilmId, id.Value)
		if err != nil {
			rh.Logger.Error("error with usecase padd review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (rh *RatingHandler) ShowReviews(w http.ResponseWriter, r *http.Request) {
	rh.Logger.Info("Show reviews")
	vars := mux.Vars(r)
	filmId := vars["film_id"]
	reviews, err := rh.UseCase.GetReviews(filmId)
	if err != nil {
		rh.Logger.Error("error with usecase show review")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(&reviews)
	if err != nil {
		rh.Logger.Error("error with delivery show reviews json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (rh *RatingHandler) GetCurrentUserRating() http.HandlerFunc {
	type Film struct {
		FilmId int    `'json:"film_id"'`
	}
	type Rate struct {
		Rate int    `json:"rate"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		rh.Logger.Info("Add review")
		l := Film{}
		rate := Rate{}
		id, err := r.Cookie("session_id")
		if err != nil {
			rh.Logger.Error("No cookie delivery add review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			rh.Logger.Error("error with delivery add review json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rate.Rate, err = rh.UseCase.GetCurrentRating(l.FilmId, id.Value)
		if err != nil {
			rh.Logger.Error("error with usecase padd review")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := json.Marshal(&rate)
		if err != nil {
			rh.Logger.Error("error with delivery show reviews json-marshal")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
