package user

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRecorderMockRepository
}

type MockRecorderMockRepository struct {
	mock *MockRepository
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRecorderMockRepository{mock}
	return mock
}

func (m *MockRepository) EXPECT() *MockRecorderMockRepository {
	return m.recorder
}

func (m *MockRepository) Create(user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), user)
}

func (m *MockRepository) FindByLogin(login string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLogin", login)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) FindByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLogin", reflect.TypeOf((*MockRepository)(nil).FindByLogin), login)
}

func (m *MockRepository) FindById(id int) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepository)(nil).FindById), id)
}

func (m *MockRepository) CheckExist(login string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExist", login)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) CheckExist(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExist", reflect.TypeOf((*MockRepository)(nil).CheckExist), login)
}
