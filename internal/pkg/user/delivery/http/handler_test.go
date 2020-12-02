package http

import (
	"errors"
	"fmt"
	user2 "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/user"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	user1 "github.com/Arkadiyche/http-rest-api/internal/pkg/user"
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


		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Me(gomock.Eq(testSession.Id)).
			Return(&testUser, nil)
		userHandler := UserHandler{
			RpcAuth: m1,
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
		assert.Equal(t, rr.Body.String(), fmt.Sprint("{\"id\":", strconv.Itoa(testUser.Id), ",\"login\":\"", testUser.Username, "\",\"email\":\"", testUser.Email, "\"}"))
	})
}

func TestLogout(t *testing.T) {
	t.Run("Logout-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m1.
			EXPECT().
			Logout(gomock.Eq(testSession.Id)).
			Return(nil)
		userHandler := UserHandler{
			RpcAuth: m1,
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
		assert.Equal(t, rr.Body.String(), "{\"id\":0,\"login\":\"\",\"email\":\"\"}")
	})

	t.Run("Logout-Fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m1.
			EXPECT().
			Logout(gomock.Eq(testSession.Id)).
			Return(errors.New("fail"))

		userHandler := UserHandler{
			RpcAuth: m1,
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
