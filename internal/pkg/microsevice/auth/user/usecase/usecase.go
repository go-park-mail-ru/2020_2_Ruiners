package usecase

import (
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/bussines/crypto"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/user"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/client"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	uuid2 "github.com/satori/go.uuid"
)

type UserUseCase struct {
	UserRepository user.Repository
	RpcSession     client.ISessionClient
}

func NewUserUseCase(userRepository user.Repository, rpcSession client.ISessionClient) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		RpcSession:     rpcSession,
	}
}

func (u *UserUseCase) Signup(username, email, password string) (string, error) {
	check, err := u.UserRepository.CheckExist(username)
	if err != nil {
		return "", err
	}
	if check {
		return "", errors.New("user alredy exist")
	}
	password, err = crypto.HashPassword(password)
	if err != nil {
		return "", errors.New("bad hash")
	}
	user := models.User{Username: username, Password: password, Email: email}
	_, err1 := u.UserRepository.Create(&user)
	if err1 != nil {
		return "", err1
	}
	sessionId := uuid2.NewV4().String()
	session := models.Session{Id: sessionId, Username: username}
	err2 := u.RpcSession.Create(session.Id, session.Username)
	if err2 != nil {
		return "", err2
	}

	return sessionId, nil
}

func (u *UserUseCase) Login(login, password string) (string, error) {
	user, err := u.UserRepository.FindByLogin(login)
	if user == nil {
		return "", errors.New("user not found")
	}
	if err != nil {
		return "", err
	}
	check, err := crypto.CheckPassword(password, user.Password)
	if !check || err != nil{
		return "", errors.New("wrong password")
	}
	sessionId := uuid2.NewV4().String()
	session := models.Session{Id: sessionId, Username: login}
	err1 := u.RpcSession.Create(session.Id, session.Username)
	if err1 != nil {
		return "", err1
	}
	return sessionId, nil
}

func (u *UserUseCase) Logout(s string) error {
	err := u.RpcSession.Delete(s)
	if err != nil {
		return err
	}
	return nil
}
