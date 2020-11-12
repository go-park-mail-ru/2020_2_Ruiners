package film

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

func (m *MockRepository) FindByLId(id int) (*models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByLId", id)
	ret0, _ := ret[0].(*models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) FindByLId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByLId", reflect.TypeOf((*MockRepository)(nil).FindByLId), id)
}

func (m *MockRepository) FindFilmsByGenre(genre string) (*models.FilmCards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilmsByGenre", genre)
	ret0, _ := ret[0].(*models.FilmCards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) FindFilmsByGenre(genre interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilmsByGenre", reflect.TypeOf((*MockRepository)(nil).FindFilmsByGenre), genre)
}

func (m *MockRepository) FindFilmsByPerson(id int) (*models.FilmCards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFilmsByPerson", id)
	ret0, _ := ret[0].(*models.FilmCards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) FindFilmsByPerson(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFilmsByPerson", reflect.TypeOf((*MockRepository)(nil).FindFilmsByPerson), id)
}
