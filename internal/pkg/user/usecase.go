package user

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"mime/multipart"
	"os"
)

type UseCase interface {
	Signup(login, email, password string) (string, error)
	Login(login, password string) (string, error)
	Me(s string) (*models.User, error)
	GetById(id string) (*models.PublicUser, error)
	Logout(s string) error
	ChangeLogin(s string, newLogin string) error
	ChangePassword(s string, oldPassword string, newPassword string) error
	ChangeAvatar(s string, file multipart.File) error
	GetAvatar(id string) (*os.File, error)
}
