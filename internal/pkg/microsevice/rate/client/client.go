package client

type IRateClient interface {
	Rate(rating int, filmId int, session string) error
	AddReview(body string, filmId int, session string) error
}
