package usecase

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/microsevice/sesession"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testUser = models.User{
	Id: 1231,
	Username: "Arkadiy",
	Password: "Arkadiy1",
	Email: "arkadiy@mail.ru",
	Image: "def.png",
}

var testSession = models.Session{
	Id: "wefwuifbwiuhegfdjvsoafjh",
	Username: "Arkadiy",
}


func TestCreate(t *testing.T) {

	t.Run("Create-OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		m0 := user.NewMockRepository(ctrl)
		m1 := sesession.NewMockRepository(ctrl)
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
		m1 := sesession.NewMockRepository(ctrl)
		m0.
			EXPECT().
			CheckExist(gomock.Eq(testUser.Username)).
			Return(true, nil)

		useCase := NewUserUseCase(m0, m1)

		_, err := useCase.Signup(&testUser, &testSession)
		assert.Error(t, err)
	})
}