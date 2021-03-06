package usecase

import (
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/rating"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSession = models.Session{
	Id:       "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}

func TestGetReviews(t *testing.T) {
	t.Run("GET OK", func(t *testing.T) {
		var testReview = models.Review{
			Id:        1,
			UserLogin: "Arkadiy",
			TextBody:  "Piece of sh*t",
			UserId:    1,
			FilmId:    1,
			Rate:      10,
		}

		var testReviews = models.Reviews{}

		testReviews = append(testReviews, testReview)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m0 := rating.NewMockRepository(ctrl)
		m0.
			EXPECT().GetReviewsByFilmId(gomock.Eq(1)).Return(&testReviews, nil)
		m0.EXPECT().GetUserById(gomock.Eq(1)).Return(testReview.UserLogin, nil)
		m0.EXPECT().GetRating(gomock.Eq(1), gomock.Eq(1)).Return(testReview.Rate, nil)
		m1 := session.NewMockRepository(ctrl)
		usecase := NewRatingUseCase(m0, m1)
		reviews, err := usecase.GetReviews("1")
		assert.NoError(t, err)
		assert.Equal(t, *reviews, testReviews)
	})

	t.Run("GET NOT OK 1", func(t *testing.T) {
		var testReview = models.Review{
			Id:        1,
			UserLogin: "Arkadiy",
			TextBody:  "Piece of sh*t",
			UserId:    1,
			FilmId:    1,
			Rate:      10,
		}

		var testReviews = models.Reviews{}

		testReviews = append(testReviews, testReview)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m0 := rating.NewMockRepository(ctrl)
		m0.
			EXPECT().GetReviewsByFilmId(gomock.Eq(1)).Return(&testReviews, errors.New("no reviews"))
		m1 := session.NewMockRepository(ctrl)
		usecase := NewRatingUseCase(m0, m1)
		_, err := usecase.GetReviews("1")
		assert.EqualError(t, err, "no reviews")
	})

	t.Run("GET NOT OK 2", func(t *testing.T) {
		var testReview = models.Review{
			Id:        1,
			UserLogin: "Arkadiy",
			TextBody:  "Piece of sh*t",
			UserId:    1,
			FilmId:    1,
			Rate:      10,
		}

		var testReviewNotOk = models.Review{
			Id:        1,
			UserLogin: "Deleted",
			TextBody:  "Piece of sh*t",
			UserId:    1,
			FilmId:    1,
			Rate:      10,
		}
		var testReviews = models.Reviews{}
		var testReviewsNotOk = models.Reviews{}

		testReviews = append(testReviews, testReview)
		testReviewsNotOk = append(testReviewsNotOk, testReviewNotOk)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m0 := rating.NewMockRepository(ctrl)
		m0.
			EXPECT().GetReviewsByFilmId(gomock.Eq(1)).Return(&testReviews, nil)
		m0.EXPECT().GetUserById(gomock.Eq(1)).Return(testReview.UserLogin, errors.New("deleted"))
		m0.EXPECT().GetRating(gomock.Eq(1), gomock.Eq(1)).Return(testReview.Rate, nil)
		m1 := session.NewMockRepository(ctrl)
		usecase := NewRatingUseCase(m0, m1)
		reviews, err := usecase.GetReviews("1")
		assert.NoError(t, err)
		assert.Equal(t, *reviews, testReviewsNotOk)
	})

	t.Run("GET NOT OK 3", func(t *testing.T) {
		var testReview = models.Review{
			Id:        1,
			UserLogin: "Arkadiy",
			TextBody:  "Piece of sh*t",
			UserId:    1,
			FilmId:    1,
			Rate:      10,
		}

		//var testReviewNotOk = models.Review{
		//	Id:        1,
		//	UserLogin: "Deleted",
		//	TextBody:  "Piece of sh*t",
		//	UserId:    1,
		//	FilmId:    1,
		//	Rate:      10,
		//}

		var testReviewNotOk2 = models.Review{
			Id:        1,
			UserLogin: "Arkadiy",
			TextBody:  "Piece of sh*t",
			UserId:    1,
			FilmId:    1,
			Rate:      0,
		}

		testReviews := models.Reviews{}
		//testReviewsNotOk1 := models.Reviews{}
		testReviewsNotOk2 := models.Reviews{}

		testReviews = append(testReviews, testReview)
		//testReviewsNotOk1 = append(testReviewsNotOk1, testReviewNotOk)
		testReviewsNotOk2 = append(testReviewsNotOk2, testReviewNotOk2)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m0 := rating.NewMockRepository(ctrl)
		m0.
			EXPECT().GetReviewsByFilmId(gomock.Eq(1)).Return(&testReviews, nil)
		m0.EXPECT().GetUserById(gomock.Eq(1)).Return(testReview.UserLogin, nil)
		m0.EXPECT().GetRating(gomock.Eq(1), gomock.Eq(1)).Return(testReview.Rate, errors.New("0"))
		m1 := session.NewMockRepository(ctrl)
		usecase := NewRatingUseCase(m0, m1)
		reviews, err := usecase.GetReviews("1")
		assert.NoError(t, err)
		assert.Equal(t, *reviews, testReviewsNotOk2)
	})

	t.Run("GET NOT OK 4", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := rating.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		usecase := NewRatingUseCase(m0, m1)
		_, err := usecase.GetReviews("tuk,1rjmrym")
		assert.Error(t, err)
	})
}

func TestGetCurrentRating(t *testing.T) {
	t.Run("GetCurrentRating-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := rating.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(5, nil)

		m0.
			EXPECT().
			GetRating(gomock.Eq(1), gomock.Eq(5)).
			Return(9, nil)

		usecase := NewRatingUseCase(m0, m1)
		rate, err := usecase.GetCurrentRating("1", testSession.Id)
		assert.NoError(t, err)
		assert.Equal(t, rate, 9)
	})

	t.Run("GetCurrentRating-OK-Non rate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := rating.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(5, nil)

		m0.
			EXPECT().
			GetRating(gomock.Eq(1), gomock.Eq(5)).
			Return(0, errors.New("error"))

		usecase := NewRatingUseCase(m0, m1)
		rate, err := usecase.GetCurrentRating("1", testSession.Id)
		assert.NoError(t, err)
		assert.Equal(t, rate, 0)
	})

	t.Run("GetCurrentRating-OK-Non rate", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := rating.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m1.
			EXPECT().
			GetUserIdBySession(gomock.Eq(testSession.Id)).
			Return(5, errors.New("error"))

		usecase := NewRatingUseCase(m0, m1)
		_, err := usecase.GetCurrentRating("1", testSession.Id)
		assert.Error(t, err)
	})

	t.Run("GetCurrentRating-Bad in", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := rating.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		usecase := NewRatingUseCase(m0, m1)
		_, err := usecase.GetCurrentRating("tuk,1rjmrym", testSession.Id)
		assert.Error(t, err)
	})
}
