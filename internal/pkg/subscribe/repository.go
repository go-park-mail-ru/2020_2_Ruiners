package subscribe

import "github.com/Arkadiyche/http-rest-api/internal/pkg/models"

type Repository interface {
	AddSubscribe(subscriberId int, authorId int) error
	DeleteSubscribe(subscriberId int, authorId int) error
	GetAuthors(subscriberId int) (*models.PublicUsers, error)
	GetRatingFeed(subscriberId int) (*models.Feed, error)
	GetReviewFeed(subscriberId int) (*models.Feed, error)
}
