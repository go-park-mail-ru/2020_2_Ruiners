package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
)

type UserUseCase struct {
	Repository user.Repository
}

func NewUserUseCase(repository user.Repository) *UserUseCase {
	return &UserUseCase{
		Repository: repository,
	}
}

func (u *UserUseCase) Add(input *models.User) (*models.User, error)  {
	_, err := u.Repository.FindByLogin(input.Username)
	//fmt.Println(user)
	_, err1 := u.Repository.Create(input)
	if err1 != nil {
		return nil, err
	}
	return nil, nil
}
