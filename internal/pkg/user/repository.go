package user

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	Create(u *models.User) (*models.User, error)
	FindByLogin(login string) (*models.User, error)
	UpdadeLogin(oldLogin string, newLogin string) error
}