package repository

import (
	"database/sql"
	"errors"
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

func (r *SessionRepository) Create(session *models.Session) (*models.Session, error) {
	//fmt.Println(session)
	_, err := r.db.Exec("INSERT INTO session (id, username) VALUES(?, ?)", session.Id, session.Username)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *SessionRepository) FindById(s string) (*models.Session, error) {
	session := models.Session{}
	err := r.db.QueryRow("SELECT id, username FROM session WHERE id = ? ORDER BY id ASC LIMIT 1", s).
		Scan(&session.Id, &session.Username)

	if err != nil {
		return nil, errors.New("session not found")
	}

	return &session, nil
}

func (r *SessionRepository) GetUserIdBySession(s string) (int, error) {
	userId := 0
	err := r.db.QueryRow("SELECT u.id FROM users u JOIN session s ON u.username = s.username WHERE s.id = ? ORDER BY id ASC LIMIT 1", s).
		Scan(&userId)

	if err != nil {
		return 0, errors.New("user with session not found")
	}

	return userId, nil
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
