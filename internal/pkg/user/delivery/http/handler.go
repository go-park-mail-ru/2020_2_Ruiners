package http

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler struct {
	UseCase user.UseCase
	logger *logrus.Logger
}


func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	u := models.User{Username: "ke2",
		Password: "hui",
	    Email: "sosi"}
	_, err := uh.UseCase.Add(&u)
	if err != nil {
		w.Write([]byte("fuck"))
		return
	}
	w.Write([]byte("yep"))
}