package models

type User struct {
	Id       int `json:"id"`
	Username string `'json:"login"'`
	Password string `'json:"password"'`
	Email    string `'json:"email"'`
	Image    string `'json:"image"'`
}

type PublicUser struct {
	Id       int `json:"id"`
	Login string `'json:"login"'`
	Email    string `'json:"email"'`
}

type Login struct {
	Login    string `'json:"login"'`
	Password string `'json:"password"'`
}

type Signup struct {
	Login    string `'json:"login"'`
	Email    string `'json:"email"'`
	Password string `'json:"password"'`
}
