package repository

import (
	"database/sql"
	"errors"
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

func (r *UserRepository) FindByLogin(login string) (*models.User, error) {
	user := models.User{}
	err := r.db.QueryRow("SELECT id, username, password, email, image  FROM users WHERE username = ? ORDER BY id ASC LIMIT 1", login).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Image)

	if err != nil {
		return nil, errors.New("no user")
	}

	return &user, nil
}

func (r *UserRepository) FindById(id int) (*models.User, error) {
	user := models.User{}
	err := r.db.QueryRow("SELECT id, username, password, email, image  FROM users WHERE id = ? ORDER BY id ASC LIMIT 1", id).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Image)

	if err != nil {
		return nil, errors.New("no user")
	}

	return &user, nil
}

func (r *UserRepository) UpdateLogin(oldLogin string, newLogin string) error {
	_, err := r.db.Exec("UPDATE users SET username = ? WHERE username = ?", newLogin, oldLogin)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdatePassword(login string, newPassword string) error {
	_, err := r.db.Exec("UPDATE users SET password = ? WHERE username = ?", newPassword, login)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateAvatar(login string, name string) error {
	_, err := r.db.Exec("UPDATE users SET image = ? where username = ?", name, login)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CheckExist(login string) (bool, error) {
	query, err := r.db.Query("SELECT id FROM users WHERE username = ?", login)
	if err != nil {
		return false, err
	}
	defer query.Close()
	if query.Next() {
		return true, nil
	}
	return false, nil
}

func (r *UserRepository) Search(search string) (*models.PublicUsers, error) {
	search1 := "% " + search + "%"
	search2 := search + "%"
	user := models.PublicUser{}
	users := models.PublicUsers{}
	authorsQuery, err := r.db.Query("SELECT u.id, u.username, u.email FROM users u WHERE lower(u.username) LIKE ? or lower(u.username) LIKE ?", search1, search2)
	if err != nil {
		return nil, err
	}
	defer authorsQuery.Close()

	for authorsQuery.Next() {
		if authorsQuery.Scan(&user.Id, &user.Login, &user.Email) != nil {
			return nil, errors.New("db error")
		}
		users = append(users, user)
	}
	return &users, nil
}

func (r *UserRepository) Delete(login string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE username=?", login)
	if err != nil {
		return err
	}
	return nil
}
