package rating

type UseCase interface {
	Rate(rating int, filmId int, session string) error
}
