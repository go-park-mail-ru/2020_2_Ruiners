package usecase

import (
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/subscribe"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSession = models.Session{
	Id:       "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}

var testUser = models.PublicUser{
	Id: 1,
	Login: "Erik",
	Email: "er@mail.ru",
}

var testUsers = models.PublicUsers{testUser}

var testSubscribe = models.Subscribe{
	Id:        1,
	IsRating:  true,
	IsReview:  false,
	Body:      "5",
	UserId:    1,
	UserLogin: "Erik",
	FilmId:    1,
	FilmTitle: "Nachalo",
	Date:      0,
}

var testSubscribe1 = models.Subscribe{
	Id:        1,
	IsRating:  false,
	IsReview:  true,
	Body:      "soska",
	UserId:    1,
	UserLogin: "Erik",
	FilmId:    1,
	FilmTitle: "Nachalo",
	Date:      0,
}

var testFeedRating = models.Feed{testSubscribe}
var testFeedReview = models.Feed{testSubscribe1}
var testFeed = models.Feed{testSubscribe, testSubscribe1}

func TestCreate(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, nil)

		m0.
			EXPECT().AddSubscribe(gomock.Eq(2), gomock.Eq(1)).Return( nil)

		usecase := NewSubscribeUseCase(m0, m1)
		err := usecase.Create(1, testSession.Id)
		assert.NoError(t, err)
	})

	t.Run("error session", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, errors.New("error session"))
		usecase := NewSubscribeUseCase(m0, m1)
		err := usecase.Create(1, testSession.Id)
		assert.EqualError(t, err, "error session")
	})

	t.Run("error subscribe", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, nil)

		m0.
			EXPECT().AddSubscribe(gomock.Eq(2), gomock.Eq(1)).Return( errors.New("error subscribe"))

		usecase := NewSubscribeUseCase(m0, m1)
		err := usecase.Create(1, testSession.Id)
		assert.EqualError(t, err, "error subscribe")
	})
}

func TestDelete(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, nil)

		m0.
			EXPECT().DeleteSubscribe(gomock.Eq(2), gomock.Eq(1)).Return( nil)

		usecase := NewSubscribeUseCase(m0, m1)
		err := usecase.Delete(1, testSession.Id)
		assert.NoError(t, err)
	})

	t.Run("error session", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, errors.New("error session"))

		usecase := NewSubscribeUseCase(m0, m1)
		err := usecase.Delete(1, testSession.Id)
		assert.EqualError(t, err, "error session")
	})

	t.Run("error delete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, nil)

		m0.
			EXPECT().DeleteSubscribe(gomock.Eq(2), gomock.Eq(1)).Return( errors.New("error delete"))

		usecase := NewSubscribeUseCase(m0, m1)
		err := usecase.Delete(1, testSession.Id)
		assert.EqualError(t, err, "error delete")
	})
}


func TestGetAuthors(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, nil)

		m0.
			EXPECT().GetAuthors(gomock.Eq(2)).Return( &testUsers, nil)

		usecase := NewSubscribeUseCase(m0, m1)
		users, err := usecase.GetAuthors(testSession.Id)
		assert.NoError(t, err)
		assert.Equal(t, *users, testUsers)
	})

	t.Run("error session", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, errors.New("error session"))

		usecase := NewSubscribeUseCase(m0, m1)
		_, err := usecase.GetAuthors(testSession.Id)
		assert.EqualError(t, err, "error session")
	})

	t.Run("error get author", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, nil)

		m0.
			EXPECT().GetAuthors(gomock.Eq(2)).Return( &testUsers, errors.New("error"))

		usecase := NewSubscribeUseCase(m0, m1)
		_, err := usecase.GetAuthors(testSession.Id)
		assert.EqualError(t, err, "error")
	})
}

func TestGetFeed(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := subscribe.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(2, nil)

		m0.
			EXPECT().GetRatingFeed(gomock.Eq(2)).Return( &testFeedRating, nil)
		m0.EXPECT().GetReviewFeed(gomock.Eq(2)).Return(&testFeedReview, nil)
		usecase := NewSubscribeUseCase(m0, m1)
		feed, err := usecase.GetFeed(testSession.Id)
		assert.NoError(t, err)
		assert.Equal(t, *feed, testFeed)
	})

}