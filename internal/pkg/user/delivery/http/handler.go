package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	"github.com/gorilla/mux"
	uuid2 "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"image/png"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	UseCase user.UseCase
	logger *logrus.Logger
}


func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	u := models.Signup{}
	err := json.NewDecoder(r.Body).Decode(&u)
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
	err1 := uh.UseCase.Logout(session.Value)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	user := models.PublicUser{Login: "", Email: ""}
	result, err := json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(result)
}

func (uh *UserHandler) ChangeLogin() http.HandlerFunc {
	type ChangeLogin struct{
		Login string `'json:"login"'`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		l := ChangeLogin{}
		err := json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		session, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	    err1 := uh.UseCase.ChangeLogin(session.Value, l.Login)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (uh *UserHandler) ChangePassword() http.HandlerFunc {
	type ChangePassword struct {
		PasswordOld string `'json:"password_old"'`
		Password    string `'json:"password"'`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		l := ChangePassword{}
		err := json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		session, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err1 := uh.UseCase.ChangePassword(session.Value, l.PasswordOld, l.Password)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (uh *UserHandler) ChangeAvatar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Avatar")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("aaa")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	session, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = uh.UseCase.ChangeAvatar(session.Value, file)
	if err != nil {
		fmt.Println("aa")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("ok")
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) AvatarById(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("111")
	vars := mux.Vars(r)
	id := vars["id"]
	file, err := uh.UseCase.GetAvatar(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *file); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
	fmt.Println("1111111111111111111111111")
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(buffer.Bytes())
}



func CreateSession(w http.ResponseWriter, sessionId string) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   sessionId,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}