package user

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type UseCase interface {
	Add(input *models.User) (*models.User, error)
}
