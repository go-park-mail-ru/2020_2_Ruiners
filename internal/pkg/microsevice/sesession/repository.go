package sesession

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	Create(session *models.Session) (*models.Session, error)
	FindById(s string) (*models.Session, error)
	Delete(s string) error
	UpdateLogin(oldLogin string, newLogin string) error
}