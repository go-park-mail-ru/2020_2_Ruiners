package film

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	FindByLId(id int) (*models.Film, error)
}