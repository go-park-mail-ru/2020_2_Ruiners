package subscribe

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type UseCase interface {
	Create(authorId int, session string) error
	Delete(authorId int, session string) error
	GetAuthors(session string) (*models.PublicUsers, error)
	GetFeed(session string) (*models.Feed, error)
	Check(session string, authorId int) (bool, error)
}
