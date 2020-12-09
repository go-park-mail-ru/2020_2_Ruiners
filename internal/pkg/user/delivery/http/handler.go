package http

import (
	"encoding/json"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/client"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type UserHandler struct {
	RpcAuth   client.IAuthClient
	UseCase   user.UseCase
	Logger    *logrus.Logger
	Sanitazer *bluemonday.Policy
}

func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	u := new(models.Signup)
	uh.Logger.Info("signup")
	err := easyjson.UnmarshalFromReader(r.Body, u)
	u.Login, u.Email, u.Password = uh.Sanitazer.Sanitize(u.Login), uh.Sanitazer.Sanitize(u.Email), uh.Sanitazer.Sanitize(u.Password)
	if err != nil {
		uh.Logger.Warn("Error with user signup delivery json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sessionId, err1 := uh.RpcAuth.Signup(u.Login, u.Email, u.Password)
	if err1 != nil {
		uh.Logger.Error("error with usecase signup")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u.Password = ""
	CreateSession(w, sessionId)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	uh.Logger.Info("Login")
	l := new(models.Login)
	err := easyjson.UnmarshalFromReader(r.Body, l)
	l.Login, l.Password = uh.Sanitazer.Sanitize(l.Login), uh.Sanitazer.Sanitize(l.Password)
	if err != nil {
		uh.Logger.Error("Error with user login delivery json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sessionId, err1 := uh.RpcAuth.Login(l.Login, l.Password)
	if err1 != nil {
		uh.Logger.Error("error with usecase login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	l.Password = ""
	CreateSession(w, sessionId)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	uh.Logger.Info("Me")
	id, err := r.Cookie("session_id")
	if err != nil {
		uh.Logger.Warn("Error with user Me delivery get cookie")
		user := models.PublicUser{Login: "", Email: ""}
		result, err := easyjson.Marshal(user)
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
		uh.Logger.Error("error with usecase me")
		user := models.PublicUser{Login: "", Email: ""}
		result, err := easyjson.Marshal(user)
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
	result, err := easyjson.Marshal(public)
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
		result, err := easyjson.Marshal(user)
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
	err1 := uh.RpcAuth.Logout(session.Value)
	if err1 != nil {
		uh.Logger.Error("error with usecase logout")
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

func (uh *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	uh.Logger.Info("GetByID")
	vars := mux.Vars(r)
	id := vars["id"]
	user, err := uh.UseCase.GetById(id)
	if err != nil {
		uh.Logger.Error("error with usecase get by id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := easyjson.Marshal(user)
	if err != nil {
		uh.Logger.Error("Error with user get by id delivery json-marshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (uh *UserHandler) ChangeLogin() http.HandlerFunc {
	type ChangeLogin struct {
		Login string `json:"login"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		uh.Logger.Info("Change login")
		l := ChangeLogin{}
		err := json.NewDecoder(r.Body).Decode(&l)
		l.Login = uh.Sanitazer.Sanitize(l.Login)
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
			uh.Logger.Error("error with usecase change login")
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (uh *UserHandler) ChangePassword() http.HandlerFunc {
	type ChangePassword struct {
		PasswordOld string `json:"password_old"`
		Password    string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		uh.Logger.Info("Change password")
		l := ChangePassword{}
		err := json.NewDecoder(r.Body).Decode(&l)
		l.Password = uh.Sanitazer.Sanitize(l.Password)
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
			uh.Logger.Error("error with usecase change password")
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
		uh.Logger.Error("error with usecase change avatar")
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
		uh.Logger.Error("error with usecase avatar by id")
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
		HttpOnly: true,
		Domain: "kino-park.online",
	}
	http.SetCookie(w, &cookie)
}

func CreateCSRF(w http.ResponseWriter, CSRFToken string) {
	cookie := http.Cookie{
		Name:    "csrf_token",
		Value:   CSRFToken,
		Expires: time.Now().Add(15 * time.Minute),
		Domain: "kino-park.online",
	}
	http.SetCookie(w, &cookie)
}
