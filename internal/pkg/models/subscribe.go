package models

type Subscribe struct {
	Id int `json:"id"`
	IsRating bool `json:"is_rating"`
	IsReview bool `json:"is_review"`
	Body string `json:"body"`
	UserId int `json:"user_id"`
	UserLogin string `json:"user_login"`
	FilmId int `json:"film_id"`
	FilmTitle string `json:"film_title"`
	Date int64 `json:"date"`
}

//easyjson:json
type PublicUsers []PublicUser

//easyjson:json
type Feed []Subscribe