package film

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type UseCase interface {
	FindById(id string) (*models.Film, error)
}