package sesession

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

func (m *MockRepository) Create(session *models.Session) (*models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", session)
	ret0, _ := ret[0].(models.Session)
	ret1, _ := ret[1].(error)
	return &ret0, ret1
}

func (mr *MockRecorderMockRepository) Create(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), session)
}

func (m *MockRepository) FindById(s string) (*models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", s)
	ret0, _ := ret[0].(models.Session)
	ret1, _ := ret[1].(error)
	return &ret0, ret1
}

func (mr *MockRecorderMockRepository) FindById(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockRepository)(nil).FindById), s)
}

func (m *MockRepository) Delete(s string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", s)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) Delete(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), s)
}

func (m *MockRepository) UpdateLogin(oldLogin string, newLogin string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLogin", oldLogin, newLogin)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) UpdateLogin(oldLogin, newLogin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLogin", reflect.TypeOf((*MockRepository)(nil).UpdateLogin), oldLogin, newLogin)
}

func (m *MockRepository) GetUserIdBySession(s string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIdBySession",s)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) GetUserIdBySession(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIdBySession", reflect.TypeOf((*MockRepository)(nil).GetUserIdBySession), s)
}