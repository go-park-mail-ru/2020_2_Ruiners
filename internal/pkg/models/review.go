package models

type Review struct {
	Id int `'json:"id"'`
	UserLogin string `'json:"user_login"'`
	TextBody string `'json:"text_body"'`
	UserId int `'json:"user_id"'`
	FilmId int `'json:"film_id"'`
	Rate int `'json:"rate"'`
}

type Reviews []Review