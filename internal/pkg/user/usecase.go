package user

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"mime/multipart"
	"os"
)

type UseCase interface {
	Me(s string) (*models.User, error)
	GetById(id string) (*models.PublicUser, error)
	ChangeLogin(s string, newLogin string) error
	ChangePassword(s string, oldPassword string, newPassword string) error
	ChangeAvatar(s string, file multipart.File) error
	GetAvatar(id string) (*os.File, error)
	Search(search string) (*models.PublicUsers, error)
}
