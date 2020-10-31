package http

import (
	"encoding/json"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/rating"
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
		//body, _ := ioutil.ReadAll(r.Body)
		//fmt.Println(string(body))
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
