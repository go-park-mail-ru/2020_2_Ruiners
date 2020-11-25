package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/sesession"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/subscribe"
	"sort"
)

type SubscribeUseCase struct {
	SubscribeRepository subscribe.Repository
	SessionRepository sesession.Repository
}

func NewSubscribeUseCase(subscribeRepository subscribe.Repository, sessionRepository sesession.Repository) *SubscribeUseCase {
	return &SubscribeUseCase{
		SubscribeRepository: subscribeRepository,
		SessionRepository: sessionRepository,
	}
}

func (uc *SubscribeUseCase) Create(authorId int, session string) error {
	userId, err := uc.SessionRepository.GetUserIdBySession(session)
	if err != nil {
		return err
	}
	err = uc.SubscribeRepository.AddSubscribe(userId, authorId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *SubscribeUseCase) Delete(authorId int, session string) error {
	userId, err := uc.SessionRepository.GetUserIdBySession(session)
	if err != nil {
		return err
	}
	err = uc.SubscribeRepository.DeleteSubscribe(userId, authorId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *SubscribeUseCase) GetAuthors(session string) (*models.PublicUsers, error)  {
	userId, err := uc.SessionRepository.GetUserIdBySession(session)
	if err != nil {
		return nil, err
	}
	authors, err := uc.SubscribeRepository.GetAuthors(userId)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (uc *SubscribeUseCase) GetFeed(session string) (*models.Feed, error)  {
	feed := models.Feed{}
	userId, err := uc.SessionRepository.GetUserIdBySession(session)
	if err != nil {
		return nil, err
	}
	ratingFeed, err := uc.SubscribeRepository.GetRatingFeed(userId)
	if err != nil {
		return nil, err
	}
	reviewFeed, err := uc.SubscribeRepository.GetReviewFeed(userId)
	if err != nil {
		return nil, err
	}
	feed = append(feed, *ratingFeed...)
	feed = append(feed, *reviewFeed...)
	sort.SliceStable(feed, func(i, j int) bool {
		return feed[i].Date < feed[j].Date
	})
	return &feed, nil
}

