package http

import (
	"encoding/json"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)
type FilmHandler struct {
	UseCase film.UseCase
	logger *logrus.Logger
}


func (fh *FilmHandler) FilmById(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	film, err := fh.UseCase.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(&film)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println(res)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (fh *FilmHandler) FilmsByGenre(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	genre := vars["genre"]
	films, err := fh.UseCase.FilmsByGenre(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(&films)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
