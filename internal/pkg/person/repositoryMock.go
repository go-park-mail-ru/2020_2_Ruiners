package person

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

func (m *MockRepository) FindByFilmIdAndRole(id int, role string) (*models.FilmPersons, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByFilmIdAndRole", id, role)
	ret0, _ := ret[0].(*models.FilmPersons)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) FindByFilmIdAndRole(id interface{}, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByFilmIdAndRole", reflect.TypeOf((*MockRepository)(nil).FindByFilmIdAndRole), id, role)
}

func (m *MockRepository) FindById(id int) (*models.Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(*models.Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepository)(nil).FindById), id)
}

func (m *MockRepository) FindFilmsIdByPersonId(id int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilmsIdByPersonId", id)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) FindFilmsIdByPersonId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilmsIdByPersonId", reflect.TypeOf((*MockRepository)(nil).FindFilmsIdByPersonId), id)
}

func (m *MockRepository) Search(search string) (*models.FilmPersons, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", search)
	ret0, _ := ret[0].(*models.FilmPersons)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) Search(search interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), search)
}
