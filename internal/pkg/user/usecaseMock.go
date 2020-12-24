package user

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"mime/multipart"
	"os"
	"reflect"
)

type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockRecorderMockUseCase
}

type MockRecorderMockUseCase struct {
	mock *MockUseCase
}

func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockRecorderMockUseCase{mock}
	return mock
}

func (m *MockUseCase) EXPECT() *MockRecorderMockUseCase {
	return m.recorder
}

/*func (m *MockUseCase) Signup(input *models.User, session *models.Session) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", input, session)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) Signup(input, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockUseCase)(nil).Signup), input, session)
}*/

/*func (m *MockUseCase) Login(input *models.Login, session *models.Session) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", input, session)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) Login(input, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUseCase)(nil).Login), input, session)
}*/

func (m *MockUseCase) GetById(s string) (*models.PublicUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", s)
	ret0, _ := ret[0].(*models.PublicUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) GetById(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUseCase)(nil).GetById), s)
}

func (m *MockUseCase) Me(s string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Me", s)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) Me(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Me", reflect.TypeOf((*MockUseCase)(nil).Me), s)
}

/*func (m *MockUseCase) Logout(s string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", s)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) Logout(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUseCase)(nil).Logout), s)
}*/

func (m *MockUseCase) ChangeLogin(s string, newLogin string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeLogin", s, newLogin)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) ChangeLogin(s, newLogin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeLogin", reflect.TypeOf((*MockUseCase)(nil).ChangeLogin), s, newLogin)
}

func (m *MockUseCase) ChangePassword(s string, oldPassword string, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", s, oldPassword, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) ChangePassword(s, oldPassword, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUseCase)(nil).ChangePassword), s, oldPassword, newPassword)
}

func (m *MockUseCase) ChangeAvatar(s string, file multipart.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeAvatar", s, file)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) ChangeAvatar(s, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAvatar", reflect.TypeOf((*MockUseCase)(nil).ChangeAvatar), s, file)
}

func (m *MockUseCase) GetAvatar(id string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatar", id)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) GetAvatar(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeAvatar", reflect.TypeOf((*MockUseCase)(nil).GetAvatar), id)
}

func (m *MockUseCase) Search(s string) (*models.PublicUsers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", s)
	ret0, _ := ret[0].(*models.PublicUsers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) Search(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUseCase)(nil).Search), s)
}
