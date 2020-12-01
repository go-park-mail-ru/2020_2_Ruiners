package rating

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

func (m *MockUseCase) Rate(rating int, filmId int, session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rate", rating, filmId, session)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) Rate(rating, filmId, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUseCase)(nil).Rate), rating, filmId, session)
}

func (m *MockUseCase) AddReview(body string, filmId int, session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReview", body, filmId, session)
	ret0, _ := ret[0].(error)
	return ret0
}
