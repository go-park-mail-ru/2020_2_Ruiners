package usecase

import (
	"errors"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/bussines/crypto"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/session"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	uuid2 "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

type UserUseCase struct {
	UserRepository    user.Repository
	SessionRepository session.Repository
}

func NewUserUseCase(userRepository user.Repository, sessionRepository session.Repository) *UserUseCase {
	return &UserUseCase{
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
	}
}

func (u *UserUseCase) Signup(username, email, password string) (string , error) {
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
	err2 := u.SessionRepository.Create(session.Id, session.Username)
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
	if !check {
		return "", errors.New("wrong password")
	}
	sessionId := uuid2.NewV4().String()
	session := models.Session{Id: sessionId, Username: login}
	err1 := u.SessionRepository.Create(session.Id, session.Username)
	if err1 != nil {
		return "", err1
	}
	return sessionId, nil
}

func (u *UserUseCase) Logout(s string) error {
	err := u.SessionRepository.Delete(s)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) Me(s string) (*models.User, error) {
	_, login, err := u.SessionRepository.FindById(s)
	if err != nil {
		return nil, err
	}
	user, err1 := u.UserRepository.FindByLogin(login)
	if err1 != nil {
		return nil, err1
	}
	return user, nil
}

func (u *UserUseCase) GetById(ids string) (*models.PublicUser, error) {
	id, err := strconv.Atoi(ids)
	if err != nil {
		return nil, err
	}
	pUser := models.PublicUser{}
	user, err1 := u.UserRepository.FindById(id)
	if err1 != nil {
		return nil, err1
	}
	pUser.Id = user.Id
	pUser.Login = user.Username
	pUser.Email = user.Email
	return &pUser, nil
}

func (u *UserUseCase) ChangeLogin(s string, newLogin string) error {
	check, err := u.UserRepository.CheckExist(newLogin)
	if err != nil {
		return err
	}
	if check {
		return errors.New("user alredy exist")
	}
	_, login, err  := u.SessionRepository.FindById(s)
	if err != nil {
		return err
	}
	err1 := u.UserRepository.UpdateLogin(login, newLogin)
	if err1 != nil {
		return err1
	}
	err = u.SessionRepository.UpdateLogin(login, newLogin)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) ChangePassword(s string, oldPassword string, newPassword string) error {
	_, login, err  := u.SessionRepository.FindById(s)
	if err != nil {
		return err
	}
	user, err1 := u.UserRepository.FindByLogin(login)
	if err1 != nil {
		return err1
	}
	check, err := crypto.CheckPassword(oldPassword, user.Password)
	if err != nil {
		return err
	}
	if !check {
		return errors.New("wrong old password")
	}
	newPassword, err = crypto.HashPassword(newPassword)
	if err != nil {
		return err
	}
	err = u.UserRepository.UpdatePassword(login, newPassword)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) ChangeAvatar(s string, file multipart.File) error {
	str := uuid2.NewV4().String()
	f, err := os.Create("uploads/" + str + ".png")
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, file)
	if err != nil {
		fmt.Println("aa")
		return err
	}
	_, login, err  := u.SessionRepository.FindById(s)
	if err != nil {
		return err
	}
	_, err1 := u.UserRepository.FindByLogin(login)
	if err1 != nil {
		return err1
	}
	err = u.UserRepository.UpdateAvatar(login, str)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) GetAvatar(ids string) (*os.File, error) {
	id, err := strconv.Atoi(ids)
	if err != nil {
		return nil, err
	}
	user, err := u.UserRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	file, err := os.Open("uploads/" + user.Image + ".png")
	if err != nil {
		file, _ = os.Open("uploads/def.png")
	}
	return file, nil
}
