package http

import (
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/person"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PersonHandler struct {
	UseCase person.UseCase
	Logger  *logrus.Logger
}

func (ph *PersonHandler) PersonById(w http.ResponseWriter, r *http.Request) {
	ph.Logger.Info("Person by id")
	vars := mux.Vars(r)
	id := vars["id"]
	person, err := ph.UseCase.GetPerson(id)
	if err != nil {
		ph.Logger.Error("error with usecase person by id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := easyjson.Marshal(person)
	if err != nil {
		ph.Logger.Error("error with delivery person by id json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ph *PersonHandler) PersonsByFilm(w http.ResponseWriter, r *http.Request) {
	ph.Logger.Info("Person by film")
	vars := mux.Vars(r)
	id := vars["film_id"]
	role := vars["role"]
	persons, err := ph.UseCase.GetPersonsByFilm(id, role)
	if err != nil {
		ph.Logger.Error("error with usecase person by film")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := easyjson.Marshal(persons)
	fmt.Println(string(res))
	if err != nil {
		ph.Logger.Error("error with delivery person by film json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ph *PersonHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("key")
	persons, err := ph.UseCase.Search(query)
	if err != nil {
		ph.Logger.Error("error with usecase person by film")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := easyjson.Marshal(persons)
	//fmt.Println(string(res))
	if err != nil {
		ph.Logger.Error("error with delivery person by film json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
