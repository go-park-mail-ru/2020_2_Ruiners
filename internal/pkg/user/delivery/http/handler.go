package http

import (
	"encoding/json"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	uuid2 "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type UserHandler struct {
	UseCase user.UseCase
	logger *logrus.Logger
}


func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	u := models.Signup{}
	err := json.NewDecoder(r.Body).Decode(&u)
	fmt.Println(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	session := models.Session{Id: uuid2.NewV4().String(), Username: u.Login}
	user := models.User{Username: u.Login, Password: u.Password, Email: u.Email}
	_, err1 := uh.UseCase.Signup(&user, &session)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	CreateSession(w, session.Id)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	l := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	session := models.Session{Id: uuid2.NewV4().String(), Username: l.Login}
	_, err1 :=uh.UseCase.Login(&l, &session)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	CreateSession(w, session.Id)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session_id")
	if err != nil {
		user := models.PublicUser{Login: "", Email: ""}
		result, err := json.Marshal(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}
	user, err1 := uh.UseCase.Me(id.Value)
	if err1 != nil {
		user := models.PublicUser{Login: "", Email: ""}
		result, err := json.Marshal(&user)
		if err != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}
	public := models.PublicUser{Login: user.Username, Email: user.Email}
	result, err := json.Marshal(&public)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(public)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err != nil {
		user := models.PublicUser{Login: "", Email: ""}
		result, err := json.Marshal(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	user := models.PublicUser{Login: "", Email: ""}
	result, err := json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(result)
}

func CreateSession(w http.ResponseWriter, sessionId string) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   sessionId,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}