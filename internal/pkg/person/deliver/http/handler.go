package http

import (
	"encoding/json"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/person"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PersonHandler struct {
	UseCase person.UseCase
	logger  *logrus.Logger
}

func (ph *PersonHandler) PersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	person, err := ph.UseCase.GetPerson(id)
	res, err := json.Marshal(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(string(res))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ph *PersonHandler) PersonsByFilm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["film_id"]
	role := vars["role"]
	persons, err := ph.UseCase.GetPersonsByFilm(id, role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(&persons)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(string(res))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
