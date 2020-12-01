package user

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	Create(u *models.User) (*models.User, error)
	FindByLogin(login string) (*models.User, error)
	FindById(id int) (*models.User, error)
	CheckExist(login string) (bool, error)
}
