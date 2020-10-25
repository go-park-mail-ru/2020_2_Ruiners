package repository

import (
	"database/sql"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(u *models.User) (*models.User, error)  {
	_, err := r.db.Exec("INSERT INTO users (username, password, email) VALUES(?, ?, ?)", u.Username, u.Password, u.Email)
	if err != nil {
		return nil, err
	}
	return  nil, nil
}

func (r *UserRepository) FindByLogin(login string) (*models.User, error) {
	id, _ := r.db.Query("SELECT id FROM users WHERE username = ?", login)
	fmt.Println(id.Next())
	return nil, nil
}