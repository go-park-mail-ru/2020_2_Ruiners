package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/client"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/subscribe"
	"sort"
)

type SubscribeUseCase struct {
	SubscribeRepository subscribe.Repository
	RpcSession          client.ISessionClient
}

func NewSubscribeUseCase(subscribeRepository subscribe.Repository, rpcSession client.ISessionClient) *SubscribeUseCase {
	return &SubscribeUseCase{
		SubscribeRepository: subscribeRepository,
		RpcSession:          rpcSession,
	}
}

func (uc *SubscribeUseCase) Create(authorId int, session string) error {
	userId, err := uc.RpcSession.GetUserIdBySession(session)
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
	userId, err := uc.RpcSession.GetUserIdBySession(session)
	if err != nil {
		return err
	}
	err = uc.SubscribeRepository.DeleteSubscribe(userId, authorId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *SubscribeUseCase) GetAuthors(session string) (*models.PublicUsers, error) {
	userId, err := uc.RpcSession.GetUserIdBySession(session)
	if err != nil {
		return nil, err
	}
	authors, err := uc.SubscribeRepository.GetAuthors(userId)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (uc *SubscribeUseCase) GetFeed(session string) (*models.Feed, error) {
	feed := models.Feed{}
	userId, err := uc.RpcSession.GetUserIdBySession(session)
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
		return feed[i].Date > feed[j].Date
	})
	return &feed, nil
}

func (uc *SubscribeUseCase) Check(session string, authorId int) (bool, error) {
	subscriberId, err := uc.RpcSession.GetUserIdBySession(session)
	if err != nil {
		return false, err
	}
	check, err := uc.SubscribeRepository.Check(subscriberId, authorId)
	if err != nil {
		return false, err
	}
	return check, nil
}
