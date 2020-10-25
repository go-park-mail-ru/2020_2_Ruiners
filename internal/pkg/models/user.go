package models

type User struct {
	Id       int `json:"id"`
	Username string `'json:"username"'`
	Password string `'json:"password"'`
	Email    string `'json:"email"'`
	image    string `'json:"image"'`
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
