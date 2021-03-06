package usecase

import (
	"errors"
	"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/bussines/crypto"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/session/session"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
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

var testPubUser = models.PublicUser{
	Id:    1231,
	Login: "Arkadiy",
	Email: "arkadiy@mail.ru",
}
var testPubUsers = models.PublicUsers{
	testPubUser,
}

var testSession = models.Session{
	Id:       "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}

/*func TestCreate(t *testing.T) {

	t.Run("Create-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(testUser.Username)).
			Return(false, nil)
		m0.
			EXPECT().
			Create(gomock.Eq(&testUser)).
			Return(nil, nil)

		m1.
			EXPECT().
			Create(gomock.Eq(&testSession)).
			Return(nil, nil)

		useCase := NewUserUseCase(m0, m1)

		user, err := useCase.Signup(&testUser, &testSession)
		testUser.Password = user.Password
		assert.NoError(t, err)
		assert.Equal(t, *user, testUser)
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

		_, err := useCase.Signup(&testUser, &testSession)
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
			Create(gomock.Eq(&testSession)).
			Return(nil, nil)

		useCase := NewUserUseCase(m0, m1)
		user, err := useCase.Login(&login, &testSession)
		assert.NoError(t, err)
		assert.Equal(t, *user, testUser)
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
		_, err := useCase.Login(&login, &testSession)
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
		_, err := useCase.Login(&login, &testSession)
		assert.EqualError(t, err, "wrong password")
	})
}*/

func TestMe(t *testing.T) {

	t.Run("Me-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, nil)
		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testSession.Username)).
			Return(&testUser, nil)

		useCase := NewUserUseCase(m0, m1)

		user, err := useCase.Me(testSession.Id)
		assert.NoError(t, err)
		assert.Equal(t, *user, testUser)
	})

	t.Run("Me-FindById", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, errors.New("error session"))

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.Me(testSession.Id)
		assert.Error(t, err)
	})

	t.Run("Me-FindByLogin", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, nil)
		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testSession.Username)).
			Return(&testUser, errors.New("error session"))

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.Me(testSession.Id)
		assert.Error(t, err)
	})
}

/*func TestLogout(t *testing.T) {

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
}*/

func TestGetById(t *testing.T) {

	t.Run("GetById-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m0.
			EXPECT().
			FindById(gomock.Eq(testUser.Id)).
			Return(&testUser, nil)

		useCase := NewUserUseCase(m0, m1)

		user, err := useCase.GetById(fmt.Sprint(testUser.Id))
		assert.NoError(t, err)
		assert.Equal(t, *user, testPubUser)
	})

	t.Run("GetById1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m0.
			EXPECT().
			FindById(gomock.Eq(testUser.Id)).
			Return(nil, errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.GetById(fmt.Sprint(testUser.Id))
		assert.Error(t, err)
	})

	t.Run("GetById1", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.GetById("not number")
		assert.Error(t, err)
	})
}

func TestUpdateLogin(t *testing.T) {
	var newLogin = "geniy"
	t.Run("UpdateLogin-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(newLogin)).
			Return(false, nil)
		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, nil)

		m0.
			EXPECT().
			UpdateLogin(gomock.Eq(testUser.Username), gomock.Eq(newLogin)).
			Return(nil)

		m1.
			EXPECT().
			UpdateLogin(gomock.Eq(testUser.Username), gomock.Eq(newLogin)).
			Return(nil)

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangeLogin(testSession.Id, newLogin)
		assert.NoError(t, err)
	})

	t.Run("UpdateLogin-CheckExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(newLogin)).
			Return(false, errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangeLogin(testSession.Id, newLogin)
		assert.Error(t, err)
	})

	t.Run("UpdateLogin-UserExist", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(newLogin)).
			Return(true, nil)

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangeLogin(testSession.Id, newLogin)
		assert.EqualError(t, err, "user alredy exist")
	})

	t.Run("UpdateLogin-FindById", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(newLogin)).
			Return(false, nil)
		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangeLogin(testSession.Id, newLogin)
		assert.Error(t, err)
	})

	t.Run("UpdateLogin-UpdateLogin", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(newLogin)).
			Return(false, nil)
		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, nil)

		m0.
			EXPECT().
			UpdateLogin(gomock.Eq(testUser.Username), gomock.Eq(newLogin)).
			Return(errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangeLogin(testSession.Id, newLogin)
		assert.Error(t, err)
	})

	t.Run("UpdateLogin-UpdateLogin", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(newLogin)).
			Return(false, nil)
		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, nil)

		m0.
			EXPECT().
			UpdateLogin(gomock.Eq(testUser.Username), gomock.Eq(newLogin)).
			Return(nil)

		m1.
			EXPECT().
			UpdateLogin(gomock.Eq(testUser.Username), gomock.Eq(newLogin)).
			Return(errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangeLogin(testSession.Id, newLogin)
		assert.Error(t, err)
	})

}

func TestUpdatePassword(t *testing.T) {
	var oldPassword = testUser.Password
	var newPasword = "geniy"
	t.Run("UpdatePassword-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testUser.Password, _ = crypto.HashPassword(testUser.Password)

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, nil)

		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testSession.Username)).
			Return(&testUser, nil)

		m0.
			EXPECT().
			UpdatePassword(gomock.Eq(testUser.Username), gomock.Any()).
			Return(nil)

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangePassword(testSession.Id, oldPassword, newPasword)
		assert.NoError(t, err)
	})
	oldPassword = "wrong"
	t.Run("UpdatePassword-BadOld", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testUser.Password, _ = crypto.HashPassword(testUser.Password)

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, nil)
		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testSession.Username)).
			Return(&testUser, nil)

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangePassword(testSession.Id, oldPassword, newPasword)
		assert.EqualError(t, err, "wrong old password")
	})
	oldPassword = testUser.Password
	newPasword = "geniy"
	t.Run("UpdatePassword-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testUser.Password, _ = crypto.HashPassword(testUser.Password)

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, nil)

		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testSession.Username)).
			Return(&testUser, nil)

		m0.
			EXPECT().
			UpdatePassword(gomock.Eq(testUser.Username), gomock.Any()).
			Return(errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangePassword(testSession.Id, oldPassword, newPasword)
		assert.Error(t, err)
	})

	t.Run("UpdatePassword-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testUser.Password, _ = crypto.HashPassword(testUser.Password)

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, nil)

		m0.
			EXPECT().
			FindByLogin(gomock.Eq(testSession.Username)).
			Return(&testUser, errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangePassword(testSession.Id, oldPassword, newPasword)
		assert.Error(t, err)
	})

	t.Run("UpdatePassword-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		testUser.Password, _ = crypto.HashPassword(testUser.Password)

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)
		m1.
			EXPECT().
			FindById(gomock.Eq(testSession.Id)).
			Return(testSession.Id, testSession.Username, errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		err := useCase.ChangePassword(testSession.Id, oldPassword, newPasword)
		assert.Error(t, err)
	})
}

func TestSearch(t *testing.T) {
	t.Run("Search-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m0.
			EXPECT().
			Search(gomock.Eq(testPubUser.Login)).
			Return(&testPubUsers, nil)

		useCase := NewUserUseCase(m0, m1)

		user, err := useCase.Search(testPubUser.Login)
		assert.NoError(t, err)
		assert.Equal(t, *user, testPubUsers)
	})

	t.Run("Search-Search", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := session.NewMockRepository(ctrl)

		m0.
			EXPECT().
			Search(gomock.Eq(testPubUser.Login)).
			Return(&testPubUsers, errors.New("error"))

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.Search(testPubUser.Login)
		assert.Error(t, err)
	})
}
