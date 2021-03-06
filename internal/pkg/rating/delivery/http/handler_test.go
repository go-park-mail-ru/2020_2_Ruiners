package http

import (
	"errors"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/rating"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testReview = models.Review{
	Id:        5,
	UserLogin: "arkadiy",
	TextBody:  "aaaa",
	UserId:    1,
	FilmId:    10,
	Rate:      1,
}

func TestGetReviews(t *testing.T) {
	var testReviews = models.Reviews{}
	testReviews = append(testReviews, testReview)
	t.Run("FindById-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := rating.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetReviews(gomock.Eq("10")).
			Return(&testReviews, nil)
		ratingHandler := RatingHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/person_film/10", nil)

		vars := map[string]string{
			"film_id": "10",
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ratingHandler.ShowReviews)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "[{\"id\":5,\"user_login\":\"arkadiy\",\"text_body\":\"aaaa\",\"user_id\":1,\"film_id\":10,\"rate\":1}]")
	})

	t.Run("FindById-fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := rating.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetReviews(gomock.Eq("10")).
			Return(&testReviews, errors.New("error"))
		ratingHandler := RatingHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/person_film/10", nil)

		vars := map[string]string{
			"film_id": "10",
		}

		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, vars)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ratingHandler.ShowReviews)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}
