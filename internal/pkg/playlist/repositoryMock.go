package playlist

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

func (m *MockRepository) Create(title string, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", title, id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) Create(title interface{}, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), title, id)
}

func (m *MockRepository) Delete(playlistId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", playlistId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) Delete(playlistId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), playlistId)
}

func (m *MockRepository) Add(filmId int, playlistId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", filmId, playlistId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) Add(filmId interface{}, playlistId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepository)(nil).Add), filmId, playlistId)
}

func (m *MockRepository) Remove(filmId int, playlistId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", filmId, playlistId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) Remove(filmId interface{}, playlistId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockRepository)(nil).Remove), filmId, playlistId)
}

func (m *MockRepository) GetList(userId int) (*models.Playlists, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", userId)
	ret0, _ := ret[0].(*models.Playlists)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) GetList(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockRepository)(nil).GetList), userId)
}
