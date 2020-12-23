package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	user2 "github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/user"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	user1 "github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
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

var testLoginModel = models.Login{
	Login:    "Arkadiy",
	Password: "Arkadiy1",
}

var testPublicUser = models.PublicUser{
	Id:    1231,
	Login: "Arkadiy",
	Email: "arkadiy@mail.ru",
}

var testPublicUsers = models.PublicUsers{}

var testSession = models.Session{
	Id:       "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}

var testSignup = models.Signup{
	Login:    "Arkadiy",
	Email:    "arkadiy@mail.ru",
	Password: "Arkadiy1",
}

func TestSignup(t *testing.T) {
	t.Run("Signup-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m1.
			EXPECT().
			Signup(gomock.Eq(testSignup.Login), gomock.Eq(testSignup.Email), gomock.Eq(testSignup.Password)).
			Return(testSession.Id, nil)

		userHandler := UserHandler{
			RpcAuth:   m1,
			UseCase:   m,
			Logger:    logrus.New(),
			Sanitazer: bluemonday.UGCPolicy(),
		}

		bytesRepresentation, err := json.Marshal(testSignup)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Signup)
		handler.ServeHTTP(rr, req)
	})

	t.Run("Signup-Fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m1.
			EXPECT().
			Signup(gomock.Eq(testSignup.Login), gomock.Eq(testSignup.Email), gomock.Eq(testSignup.Password)).
			Return(testSession.Id, errors.New("error"))

		userHandler := UserHandler{
			RpcAuth:   m1,
			UseCase:   m,
			Logger:    logrus.New(),
			Sanitazer: bluemonday.UGCPolicy(),
		}

		bytesRepresentation, err := json.Marshal(testSignup)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Signup)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Logout-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m1.
			EXPECT().
			Login(gomock.Eq(testLoginModel.Login), gomock.Eq(testLoginModel.Password)).
			Return(testSession.Id, nil)

		userHandler := UserHandler{
			RpcAuth:   m1,
			UseCase:   m,
			Logger:    logrus.New(),
			Sanitazer: bluemonday.UGCPolicy(),
		}

		bytesRepresentation, err := json.Marshal(testLoginModel)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Login)
		handler.ServeHTTP(rr, req)
	})

	t.Run("Logout-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m1.
			EXPECT().
			Login(gomock.Eq(testLoginModel.Login), gomock.Eq(testLoginModel.Password)).
			Return(testSession.Id, errors.New("error"))

		userHandler := UserHandler{
			RpcAuth:   m1,
			UseCase:   m,
			Logger:    logrus.New(),
			Sanitazer: bluemonday.UGCPolicy(),
		}

		bytesRepresentation, err := json.Marshal(testLoginModel)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:    "session_id",
			Value:   testSession.Id,
			Expires: time.Now().Add(10 * time.Hour),
		})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Login)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
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

	t.Run("Me-fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		userHandler := UserHandler{
			RpcAuth: m1,
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/me", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Me)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})

	t.Run("Me-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Me(gomock.Eq(testSession.Id)).
			Return(&testUser, errors.New("error"))
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
		assert.Equal(t, rr.Code, 400)
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

	t.Run("Logout-No Coockie", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		userHandler := UserHandler{
			RpcAuth: m1,
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("POST", "/logout", nil)
		if err != nil {
			t.Fatal(err)
		}

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

func TestGetById(t *testing.T) {

	t.Run("GetById-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetById(gomock.Eq("1")).
			Return(&testPublicUser, nil)
		userHandler := UserHandler{
			RpcAuth: m1,
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/person/1", nil)

		vars := map[string]string{
			"id": "1",
		}

		req = mux.SetURLVars(req, vars)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.GetById)
		handler.ServeHTTP(rr, req)
		res, _ := json.Marshal(&testPublicUser)
		assert.Equal(t, rr.Body.String(), string(res))
	})

	t.Run("GetById-fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m.
			EXPECT().
			GetById(gomock.Eq("1")).
			Return(&testPublicUser, errors.New("error"))
		userHandler := UserHandler{
			RpcAuth: m1,
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/person/1", nil)

		vars := map[string]string{
			"id": "1",
		}

		req = mux.SetURLVars(req, vars)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.GetById)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}

func TestSearch(t *testing.T) {

	t.Run("Search-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Search(gomock.Eq("people")).
			Return(&testPublicUsers, nil)
		userHandler := UserHandler{
			RpcAuth: m1,
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/search?key=people", nil)

		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Search)
		handler.ServeHTTP(rr, req)
		res, _ := json.Marshal(&testPublicUsers)
		assert.Equal(t, rr.Body.String(), string(res))
	})

	t.Run("Search-fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := user1.NewMockUseCase(ctrl)
		m1 := user2.NewMockUseCase(ctrl)

		m.
			EXPECT().
			Search(gomock.Eq("people")).
			Return(&testPublicUsers, errors.New("error"))
		userHandler := UserHandler{
			RpcAuth: m1,
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/search?key=people", nil)

		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Search)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, 400)
	})
}
