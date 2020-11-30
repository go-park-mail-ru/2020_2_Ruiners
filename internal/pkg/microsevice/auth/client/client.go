package client

type IAuthClient interface {
	Signup(login, email, password string) (sessionId string, err error)
	Login(login string, password string) (sessionId string, err error)
	Logout(sessionId string) error
}