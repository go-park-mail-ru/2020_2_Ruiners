package user

import (
	"github.com/golang/mock/gomock"
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

func (m *MockUseCase) Signup(login, email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", login, email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) Signup(login, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockUseCase)(nil).Signup), login, email, password)
}

func (m *MockUseCase) Login(login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) Login(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUseCase)(nil).Login), login, password)
}

func (m *MockUseCase) Logout(s string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", s)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) Logout(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUseCase)(nil).Logout), s)
}
