package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/client"
)

import (
	"errors"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/bussines/crypto"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	uuid2 "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"os"
	"strconv"
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

func (u *UserUseCase) Me(s string) (*models.User, error) {
	_, login, err := u.RpcSession.FindById(s)
	if err != nil {
		fmt.Println(err)
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
	_, login, err := u.RpcSession.FindById(s)
	if err != nil {
		return err
	}
	err1 := u.UserRepository.UpdateLogin(login, newLogin)
	if err1 != nil {
		return err1
	}
	err = u.RpcSession.UpdateLogin(login, newLogin)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) ChangePassword(s string, oldPassword string, newPassword string) error {
	_, login, err := u.RpcSession.FindById(s)
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
	f, err := os.Create("/home/ubuntu/back/2020_2_Ruiners/uploads/" + str + ".png")
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, file)
	_, login, err := u.RpcSession.FindById(s)
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
	fmt.Println(user.Image)
	file, err := os.Open("/home/ubuntu/back/2020_2_Ruiners/" + user.Image + ".png")
	if err != nil {
		file, _ = os.Open("/home/ubuntu/back/2020_2_Ruiners/uploads/def.png")
	}
	return file, nil
}

func (u *UserUseCase) Search(search string) (*models.PublicUsers, error) {
	users, err := u.UserRepository.Search(search)
	if err != nil {
		return nil, err
	}
	return users, nil
}
