package playlist

import (
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
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

func (m *MockUseCase) Create(title string, session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", title, session)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) Create(title interface{}, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), title, session)
}

func (m *MockUseCase) Delete(playlistId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", playlistId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) Delete(playlistId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), playlistId)
}

func (m *MockUseCase) Add(filmId int, playlistId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersonsByFilm", filmId, playlistId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) Add(filmId interface{}, playlistId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUseCase)(nil).Add), filmId, playlistId)
}

func (m *MockUseCase) Remove(filmId int, playlistId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", filmId, playlistId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) Remove(filmId interface{}, playlistId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockUseCase)(nil).Remove), filmId, playlistId)
}

func (m *MockUseCase) GetList(search string) (*models.Playlists, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", search)
	ret0, _ := ret[0].(*models.Playlists)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) GetList(search interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockUseCase)(nil).GetList), search)
}

func (m *MockUseCase) GetPlaylist(search string) (*models.Playlists, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaylist", search)
	ret0, _ := ret[0].(*models.Playlists)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) GetPlaylist(search interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaylist", reflect.TypeOf((*MockUseCase)(nil).GetPlaylist), search)
}
