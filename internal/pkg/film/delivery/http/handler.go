package http

import (
	"encoding/json"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type FilmHandler struct {
	UseCase film.UseCase
	Logger  *logrus.Logger
}

func (fh *FilmHandler) FilmById(w http.ResponseWriter, r *http.Request) {
	fh.Logger.Info("FilmByID")
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	film, err := fh.UseCase.FindById(id)
	if err != nil {
		fh.Logger.Error("Error with Film by id usecase")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(&film)
	if err != nil {
		fh.Logger.Error("Error with film delivery film by id json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//fmt.Println(res)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (fh *FilmHandler) FilmsByGenre(w http.ResponseWriter, r *http.Request) {
	fh.Logger.Info("Film by genre")
	vars := mux.Vars(r)
	genre := vars["genre"]
	films, err := fh.UseCase.FilmsByGenre(genre)
	if err != nil {
		fh.Logger.Error("Error with Film by genre usecase")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(&films)
	if err != nil {
		fh.Logger.Error("Error with film delivery film by genre json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (fh *FilmHandler) FilmsByPerson(w http.ResponseWriter, r *http.Request) {
	fh.Logger.Info("Film by person")
	vars := mux.Vars(r)
	id := vars["id"]
	films, err := fh.UseCase.FilmsByPerson(id)
	if err != nil {
		fh.Logger.Error("Error with Film by person usecase")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(&films)
	if err != nil {
		fh.Logger.Error("Error with film delivery film by person json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
