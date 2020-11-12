package http

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"

	//"fmt"
	"strconv"

	"github.com/Arkadiyche/http-rest-api/internal/pkg/film"
	//"errors"
	//"encoding/json"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	//"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	//"time"
)

var testFilm = models.Film{
	Id:          5,
	Title:       "string",
	Rating:      7,
	SumVotes:    5,
	Description: "string",
	MainGenre:   "string",
	YoutubeLink: "string",
	BigImg:      "string",
	SmallImg:    "string",
	Year:        2007,
	Country:     "string",
}

func TestFindById(t *testing.T) {

	t.Run("FindById-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m := film.NewMockUseCase(ctrl)

		m.
			EXPECT().
			FindById(gomock.Eq(strconv.Itoa(testFilm.Id))).
			Return(&testFilm, nil)
		filmHandler := FilmHandler{
			UseCase: m,
			Logger:  logrus.New(),
		}
		req, err := http.NewRequest("GET", "/film/5", nil)

		vars := map[string]string{
			"id": "5",
		}

		// CHANGE THIS LINE!!!
		req = mux.SetURLVars(req, vars)
		if err != nil {
			t.Fatal(err)
		}
		//req.AddCookie(&http.Cookie{
		//	Name:    "session_id",
		//	Value:   testSession.Id,
		//	Expires: time.Now().Add(10 * time.Hour),
		//})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(filmHandler.FilmById)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Body.String(), "{\"id\":5,\"title\":\"string\",\"Rating\":7,\"SumVotes\":5,\"Description\":\"string\",\"MainGenre\":\"string\",\"YoutubeLink\":\"string\",\"BigImg\":\"string\",\"SmallImg\":\"string\",\"year\":2007,\"country\":\"string\"}")
	})
}

//func TestLogout(t *testing.T) {
//	t.Run("Logout-OK", func(t *testing.T) {
//		ctrl := gomock.NewController(t)
//		defer ctrl.Finish()
//
//		m := user.NewMockUseCase(ctrl)
//
//		m.
//			EXPECT().
//			Logout(gomock.Eq(testSession.Id)).
//			Return(nil)
//		userHandler := UserHandler{
//			UseCase: m,
//			Logger:  logrus.New(),
//		}
//		req, err := http.NewRequest("POST", "/logout", nil)
//		if err != nil {
//			t.Fatal(err)
//		}
//		req.AddCookie(&http.Cookie{
//			Name:    "session_id",
//			Value:   testSession.Id,
//			Expires: time.Now().Add(10 * time.Hour),
//		})
//		rr := httptest.NewRecorder()
//		handler := http.HandlerFunc(userHandler.Logout)
//		handler.ServeHTTP(rr, req)
//		assert.Equal(t, rr.Body.String(), "{\"id\":0,\"Login\":\"\",\"Email\":\"\"}")
//	})
//
//	t.Run("Logout-Fail", func(t *testing.T) {
//		ctrl := gomock.NewController(t)
//		defer ctrl.Finish()
//
//		m := user.NewMockUseCase(ctrl)
//
//		m.
//			EXPECT().
//			Logout(gomock.Eq(testSession.Id)).
//			Return(errors.New("fail"))
//
//		userHandler := UserHandler{
//			UseCase: m,
//			Logger:  logrus.New(),
//		}
//		req, err := http.NewRequest("POST", "/logout", nil)
//		if err != nil {
//			t.Fatal(err)
//		}
//		req.AddCookie(&http.Cookie{
//			Name:    "session_id",
//			Value:   testSession.Id,
//			Expires: time.Now().Add(10 * time.Hour),
//		})
//		rr := httptest.NewRecorder()
//		handler := http.HandlerFunc(userHandler.Logout)
//
//		handler.ServeHTTP(rr, req)
//		resp := rr.Result()
//		if resp.StatusCode != 400 {
//			t.Errorf("expected resp status 400, got %d", resp.StatusCode)
//			return
//		}
//	})
//}

