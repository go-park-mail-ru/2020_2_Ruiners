package film

import (
	//"fmt"
	"github.com/Arkadiyche/http-rest-api/internal/pkg/models"
	"github.com/golang/mock/gomock"
	//"mime/multipart"
	//"os"/
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

func (m *MockUseCase) FindById(s string) (*models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", s)
	ret0, _ := ret[0].(*models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) FindById(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUseCase)(nil).FindById), s)
}

func (m *MockUseCase) FilmsByGenre(s string) (*models.FilmCards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilmsByGenre", s)
	ret0, _ := ret[0].(*models.FilmCards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) FilmsByGenre(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilmsByGenre", reflect.TypeOf((*MockUseCase)(nil).FilmsByGenre), s)
}

func (m *MockUseCase) FilmsByPerson(s string) (*models.FilmCards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilmsByPerson", s)
	ret0, _ := ret[0].(*models.FilmCards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) FilmsByPerson(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilmsByPerson", reflect.TypeOf((*MockUseCase)(nil).FilmsByPerson), s)
}
