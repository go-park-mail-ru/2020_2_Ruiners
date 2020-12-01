package rating

import (
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

func (m *MockRepository) AddRating(rating int, filmId int, userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRating", rating, filmId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) AddRating(rating interface{}, filmId interface{}, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRating", reflect.TypeOf((*MockRepository)(nil).AddRating), rating, filmId, userId)
}

func (m *MockRepository) UpdateRating(rating int, filmId int, userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRating", rating, filmId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) UpdateRating(rating interface{}, filmId interface{}, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRating", reflect.TypeOf((*MockRepository)(nil).UpdateRating), rating, filmId, userId)
}

func (m *MockRepository) CheckRating(filmId int, userId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRating", filmId, userId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) CheckRating(filmId interface{}, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRating", reflect.TypeOf((*MockRepository)(nil).CheckRating), filmId, userId)
}

func (m *MockRepository) AddReview(body string, filmId int, userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReview", body, filmId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) AddReview(body interface{}, filmId interface{}, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReview", reflect.TypeOf((*MockRepository)(nil).AddReview), body, filmId, userId)
}

