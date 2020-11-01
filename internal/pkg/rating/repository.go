package rating

type Repository interface {
	Add(rating int, filmId int, userId int) error
	Update(rating int, filmId int, userId int) error
	Check(filmId int, userId int) (bool, error)
}