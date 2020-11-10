package http

import (
	"encoding/json"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	"github.com/gorilla/mux"
	uuid2 "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type UserHandler struct {
	UseCase user.UseCase
	Logger  *logrus.Logger
}

func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	u := models.Signup{}
	uh.Logger.Info("signup")
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		uh.Logger.Warn("Error with user signup delivery json")
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
	uh.Logger.Info("Login")
	l := models.Login{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		uh.Logger.Error("Error with user login delivery json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	session := models.Session{Id: uuid2.NewV4().String(), Username: l.Login}
	_, err1 := uh.UseCase.Login(&l, &session)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	CreateSession(w, session.Id)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	uh.Logger.Info("Me")
	id, err := r.Cookie("session_id")
	if err != nil {
		uh.Logger.Warn("Error with user Me delivery get cookie")
		user := models.PublicUser{Login: "", Email: ""}
		result, err := json.Marshal(&user)
		if err != nil {
			uh.Logger.Error("Error with user me delivery json")
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
			uh.Logger.Error("Error with user Me delivery json-Marshal")
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}
	public := models.PublicUser{Id: user.Id, Login: user.Username, Email: user.Email}
	result, err := json.Marshal(&public)
	if err != nil {
		uh.Logger.Error("Error with user Me delivery json - marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	uh.Logger.Info("logout")
	session, err := r.Cookie("session_id")
	if err != nil {
		uh.Logger.Warn("No cookie")
		user := models.PublicUser{Login: "", Email: ""}
		result, err := json.Marshal(&user)
		if err != nil {
			uh.Logger.Error("Error with user logout delivery json")
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
		uh.Logger.Error("Error with user logout delivery json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(result)
}

func (uh *UserHandler) ChangeLogin() http.HandlerFunc {
	type ChangeLogin struct {
		Login string `'json:"login"'`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		uh.Logger.Info("Change login")
		l := ChangeLogin{}
		err := json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			uh.Logger.Error("Error with user changelogin delivery json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		session, err := r.Cookie("session_id")
		if err != nil {
			uh.Logger.Error("No cookie changelogin")
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
		uh.Logger.Info("Change password")
		l := ChangePassword{}
		err := json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			uh.Logger.Error("Error with user change password delivery json-decode")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		session, err := r.Cookie("session_id")
		if err != nil {
			uh.Logger.Error("No cookie change password")
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
	uh.Logger.Info("Change avatar")
	file, _, err := r.FormFile("file")
	if err != nil {
		uh.Logger.Error("Error with user Change avatar delivery file")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	session, err := r.Cookie("session_id")
	if err != nil {
		uh.Logger.Error("No cookie change avatar")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = uh.UseCase.ChangeAvatar(session.Value, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) AvatarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	file, err := uh.UseCase.GetAvatar(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}

func CreateSession(w http.ResponseWriter, sessionId string) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   sessionId,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}
