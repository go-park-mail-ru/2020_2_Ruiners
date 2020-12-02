package http

import (
	"encoding/json"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/subscribe"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var testUser = models.PublicUser{
	Id: 1,
	Login: "Erik",
	Email: "er@mail.ru",
}

var testUsers = models.PublicUsers{testUser}

var testSession = models.Session{
	Id:       "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}

func TestShowAuthors(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := subscribe.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetAuthors(gomock.Eq(testSession.Id)).
			Return(&testUsers, nil)
		subscribeHandler := SubscribeHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}

		req, err := http.NewRequest("GET", "/authors", nil)
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(subscribeHandler.ShowAuthors)
		handler.ServeHTTP(rr, req)
		res, _ := json.Marshal(testUsers)
		assert.Equal(t, rr.Body.String(), string(res))
	})
}