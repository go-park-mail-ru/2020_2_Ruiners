package usecase

import (
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/sesession"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
)

type UserUseCase struct {
	UserRepository user.Repository
	SessionRepository sesession.Repository
}

func NewUserUseCase(userRepository user.Repository, sessionRepository sesession.Repository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		SessionRepository: sessionRepository,
	}
}

func (u *UserUseCase) Signup(input *models.User, session *models.Session) (*models.User, error)  {
	user, err := u.UserRepository.FindByLogin(input.Username)
	if user != nil {
		return nil, errors.New("user alredy exist")
	}
	if err != nil {
		return nil, err
	}
	_, err1 := u.UserRepository.Create(input)
	if err1 != nil {
		return nil, err1
	}
	_, err2 := u.SessionRepository.Create(session)
	if err2 != nil {
		return nil, err2
	}
	return nil, nil
}


func (u *UserUseCase) Login(input *models.Login, session *models.Session) (*models.User, error)  {
	user, err := u.UserRepository.FindByLogin(input.Login)
	if user == nil {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	if user.Password != input.Password {
		return nil, errors.New("wrong password")
	}
	_, err1 := u.SessionRepository.Create(session)
	if err1 != nil {
		return nil, err1
	}
	return user, nil
}

func (u *UserUseCase) Me(s string) (*models.User, error) {
	session, err := u.SessionRepository.FindById(s)
	if err != nil {
		return nil, err
	}
	user, err1 := u.UserRepository.FindByLogin(session.Username)
	if err1 != nil {
		return nil, err1
	}
	return user, nil
}

func (u *UserUseCase) Logout(s string) error {
	err := u.SessionRepository.Delete(s)
	if err != nil {
		return err
	}
	return nil
}

