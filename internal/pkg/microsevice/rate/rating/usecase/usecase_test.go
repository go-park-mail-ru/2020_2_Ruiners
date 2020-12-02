package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/rate/rating"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSession = models.Session{
	Id:       "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}

func TestRate(t *testing.T) {
	t.Run("RATE UPDATE", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m1 := session.NewMockRepository(ctrl)
		m1.EXPECT().GetUserIdBySession(gomock.Eq(testSession.Id)).Return(1, nil)
		m0 := rating.NewMockRepository(ctrl)
		m0.
			EXPECT().CheckRating(gomock.Eq(1), gomock.Eq(1)).Return(true, nil)
		m0.EXPECT().UpdateRating(gomock.Eq(1), gomock.Eq(1), gomock.Eq(1)).Return(nil)
		usecase := NewRatingUseCase(m0, m1)
		err := usecase.Rate(1, 1, testSession.Id)
		assert.NoError(t, err)
	})

	t.Run("RATE ADD", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m1 := session.NewMockRepository(ctrl)
		m1.EXPECT().GetUserIdBySession(gomock.Eq(testSession.Id)).Return(1, nil)
		m0 := rating.NewMockRepository(ctrl)
		m0.
			EXPECT().CheckRating(gomock.Eq(1), gomock.Eq(1)).Return(false, nil)
		m0.EXPECT().AddRating(gomock.Eq(1), gomock.Eq(1), gomock.Eq(1)).Return(nil)
		usecase := NewRatingUseCase(m0, m1)
		err := usecase.Rate(1, 1, testSession.Id)
		assert.NoError(t, err)
	})
}

func TestAddReview(t *testing.T) {
	t.Run("ADD REVIEW", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m1 := session.NewMockRepository(ctrl)
		m1.EXPECT().GetUserIdBySession(gomock.Eq(testSession.Id)).Return(1, nil)
		m0 := rating.NewMockRepository(ctrl)
		m0.
			EXPECT().AddReview(gomock.Eq("ERIK"), gomock.Eq(1), gomock.Eq(1)).Return(nil)
		usecase := NewRatingUseCase(m0, m1)
		err := usecase.AddReview("ERIK", 1, testSession.Id)
		assert.NoError(t, err)
	})
}
