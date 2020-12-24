package usecase

import (
	"errors"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/bussines/crypto"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/auth/user"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestCreate(t *testing.T) {

	t.Run("Create-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(testUser.Username)).
			Return(false, nil)
		fmt.Println(testUser)
		m0.
			EXPECT().
			Create(gomock.Any()).
			Return(nil, nil)
		m1.
			EXPECT().
			Create(gomock.Any(), gomock.Eq(testUser.Username)).
			Return(nil)

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.Signup(testUser.Username, testUser.Email, testUser.Password)
		assert.NoError(t, err)
	})

	t.Run("Create-CheckExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(testUser.Username)).
			Return(true, errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.Signup(testUser.Username, testUser.Email, testUser.Password)
		assert.Error(t, err)
	})

	t.Run("Create-UserExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(testUser.Username)).
			Return(true, nil)

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.Signup(testUser.Username, testUser.Email, testUser.Password)
		assert.Error(t, err)
	})

	t.Run("Create-Create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(testUser.Username)).
			Return(false, nil)
		fmt.Println(testUser)
		m0.
			EXPECT().
			Create(gomock.Any()).
			Return(nil, errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.Signup(testUser.Username, testUser.Email, testUser.Password)
		assert.Error(t, err)
	})

	t.Run("Create-Create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(testUser.Username)).
			Return(false, nil)
		fmt.Println(testUser)
		m0.
			EXPECT().
			Create(gomock.Any()).
			Return(nil, nil)
		m1.
			EXPECT().
			Create(gomock.Any(), gomock.Eq(testUser.Username)).
			Return(errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.Signup(testUser.Username, testUser.Email, testUser.Password)
		assert.Error(t, err)
	})
}

func TestLogin(t *testing.T) {
	var newErr = errors.New("error")
	t.Run("Login-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		login := models.Login{
			Login:    testUser.Username,
			Password: testUser.Password,
		}

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		testUser.Password, _ = crypto.HashPassword(testUser.Password)
		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testUser.Username)).
			Return(&testUser, nil)
		m1.
			EXPECT().
			Create(gomock.Any(), gomock.Eq(testUser.Username)).
			Return(nil)

		useCase := NewUserUseCase(m0, m1)
		_, err := useCase.Login(login.Login, login.Password)
		assert.NoError(t, err)
	})

	t.Run("Login-FindByLogin", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		login := models.Login{
			Login:    testUser.Username,
			Password: testUser.Password,
		}

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		testUser.Password, _ = crypto.HashPassword(testUser.Password)
		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testUser.Username)).
			Return(&testUser, errors.New("error"))

		useCase := NewUserUseCase(m0, m1)
		_, err := useCase.Login(login.Login, login.Password)
		assert.Error(t, err)
	})

	t.Run("Login-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		login := models.Login{
			Login:    testUser.Username,
			Password: testUser.Password,
		}

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		testUser.Password, _ = crypto.HashPassword(testUser.Password)
		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testUser.Username)).
			Return(&testUser, nil)
		m1.
			EXPECT().
			Create(gomock.Any(), gomock.Eq(testUser.Username)).
			Return(errors.New("error"))

		useCase := NewUserUseCase(m0, m1)
		_, err := useCase.Login(login.Login, login.Password)
		assert.Error(t, err)
	})

	t.Run("Login-UserNotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		login := models.Login{
			Login:    testUser.Username,
			Password: testUser.Password,
		}

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testUser.Username)).
			Return(nil, newErr)

		useCase := NewUserUseCase(m0, m1)
		_, err := useCase.Login(login.Login, login.Password)
		assert.EqualError(t, err, "user not found")
	})

	t.Run("Login-WrongPass", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		login := models.Login{
			Login:    testUser.Username,
			Password: "wrong",
		}

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		testUser.Password, _ = crypto.HashPassword(testUser.Password)
		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testUser.Username)).
			Return(&testUser, nil)

		useCase := NewUserUseCase(m0, m1)
		_, err := useCase.Login(login.Login, login.Password)
		assert.EqualError(t, err, "wrong password")
	})
}

func TestLogout(t *testing.T) {

	t.Run("Logout-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m1.
			EXPECT().
			Delete(gomock.Eq(testSession.Id)).
			Return(nil)

		useCase := NewUserUseCase(m0, m1)

		err := useCase.Logout(testSession.Id)
		assert.NoError(t, err)
	})

	t.Run("Logout-Delete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m1.
			EXPECT().
			Delete(gomock.Eq(testSession.Id)).
			Return(errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		err := useCase.Logout(testSession.Id)
		assert.Error(t, err)
	})
}
