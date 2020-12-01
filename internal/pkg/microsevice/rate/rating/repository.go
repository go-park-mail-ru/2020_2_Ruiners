package rating

type Repository interface {
	AddRating(rating int, filmId int, userId int) error
	UpdateRating(rating int, filmId int, userId int) error
	CheckRating(filmId int, userId int) (bool, error)
	AddReview(body string, filmId int, userId int) error
}