package client

type ISessionClient interface {
	Create(sessionId, login string) error
	FindById(s string) (sessionId, login string, error error)
	Delete(s string) error
	UpdateLogin(oldLogin string, newLogin string) error
	GetUserIdBySession(s string) (int, error)
}
