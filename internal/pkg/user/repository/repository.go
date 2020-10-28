package repository

import (
	"database/sql"
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
	user := models.User{}
	id, err := r.db.Query("SELECT id, username, password, email, image  FROM users WHERE username = ? ORDER BY id ASC LIMIT 1", login)
	if err != nil {
		return nil, err
	}
	if id.Next() {
		id.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Image)
	} else {
		return nil, nil
	}
	return &user, nil
}