package user

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type UseCase interface {
	Signup(input *models.User, session *models.Session) (*models.User, error)
	Login(input *models.Login, session *models.Session) (*models.User, error)
	Me(s string) (*models.User, error)
	Logout(s string) error
	ChangeLogin(s string, newLogin string) error
	ChangePassword(s string, oldPassword string, newPassword string) error
}
