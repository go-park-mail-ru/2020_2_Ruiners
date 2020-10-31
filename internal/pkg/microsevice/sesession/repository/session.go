package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (r *SessionRepository) Create(session *models.Session) (*models.Session, error)  {
	fmt.Println(session)
	_, err := r.db.Exec("INSERT INTO session (id, username) VALUES(?, ?)", session.Id, session.Username)
	if err != nil {
		return nil, err
	}
	return  nil, nil
}

func (r *SessionRepository) FindById(s string) (*models.Session, error) {
	id, err := r.db.Query("SELECT id, username FROM session WHERE id = ? ORDER BY id ASC LIMIT 1", s)
	defer id.Close()
	if err != nil {
		return nil, err
	}
	session := models.Session{}
	if id.Next() {
		id.Scan(&session.Id, &session.Username)
	} else {
		return nil, errors.New("session not found")
	}
	return &session, nil
}

func (r *SessionRepository) GetUserIdBySession(s string) (int, error) {
	userId := 0
	id, err := r.db.Query("SELECT u.id FROM users u JOIN session s ON u.username = s.username WHERE s.id = ? ORDER BY id ASC LIMIT 1", s)
	defer id.Close()
	if err != nil {
		return 0, err
	}
	if id.Next() {
		id.Scan(&userId)
		return userId, nil
	}
	return 0, errors.New("user with session not found")
}

func (r *SessionRepository) Delete(s string) error {
	_, err := r.db.Exec("DELETE FROM session WHERE id=?", s)
	if err != nil {
		return err
	}
	return nil
}
func (r *SessionRepository) UpdateLogin(oldLogin string, newLogin string) error {
	_, err := r.db.Exec("UPDATE session SET username = ? WHERE username = ?", newLogin, oldLogin)
	if err != nil {
		return err
	}
	return nil
}


