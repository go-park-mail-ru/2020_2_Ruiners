package http

import (
	"errors"
	//"encoding/json"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

var testUser = models.User{
	Id:       1231,
	Username: "Arkadiy",
	Password: "Arkadiy1",
	Email:    "arkadiy@mail.ru",
	Image:    "def.png",
}

var testSession = models.Session{
	Id:       "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}

func TestMe(t *testing.T) {
	t.Run("Me-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Me(gomock.Eq(testSession.Id)).
			Return(&testUser, nil)
		userHandler := UserHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/me", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Me)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), fmt.Sprint("{\"id\":", strconv.Itoa(testUser.Id), ",\"Login\":\"", testUser.Username, "\",\"Email\":\"", testUser.Email, "\"}"))
	})
}

func TestLogout(t *testing.T) {
	t.Run("Logout-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Logout(gomock.Eq(testSession.Id)).
			Return(nil)
		userHandler := UserHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("POST", "/logout", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Logout)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "{\"id\":0,\"Login\":\"\",\"Email\":\"\"}")
	})

	t.Run("Logout-Fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Logout(gomock.Eq(testSession.Id)).
			Return(errors.New("fail"))

		userHandler := UserHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("POST", "/logout", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Logout)

		handler.ServeHTTP(rr, req)
		resp := rr.Result()
		if resp.StatusCode != 400 {
			t.Errorf("expected resp status 400, got %d", resp.StatusCode)
			return
		}
	})
}
