package person

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

func (m *MockUseCase) GetPerson(id string) (*models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPerson", id)
	ret0, _ := ret[0].(*models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) GetPerson(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPerson", reflect.TypeOf((*MockUseCase)(nil).GetPerson), id)
}

func (m *MockUseCase) GetPersonsByFilm(id string, role string) (*models.FilmPersons, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersonsByFilm", id, role)
	ret0, _ := ret[0].(*models.FilmPersons)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) GetPersonsByFilm(id interface{}, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersonsByFilm", reflect.TypeOf((*MockUseCase)(nil).GetPersonsByFilm), id, role)
}

func (m *MockUseCase) Search(search string) (*models.FilmPersons, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", search)
	ret0, _ := ret[0].(*models.FilmPersons)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) Search(search interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUseCase)(nil).Search), search)
}
