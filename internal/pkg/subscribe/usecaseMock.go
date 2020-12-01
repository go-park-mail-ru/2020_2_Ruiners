package subscribe

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

func (m *MockUseCase) Create(authorId int, session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", authorId, session)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) Create(authorId, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), authorId,session)
}


func (m *MockUseCase) Delete(authorId int, session string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", authorId, session)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockUseCase) Delete(authorId, session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), authorId, session)
}

func (m *MockUseCase) GetAuthors(session string) (*models.PublicUsers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthors", session)
	ret0, _ := ret[0].(*models.PublicUsers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) GetAuthors(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthors", reflect.TypeOf((*MockUseCase)(nil).GetAuthors), session)
}

func (m *MockUseCase) GetFeed(session string) (*models.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeed", session)
	ret0, _ := ret[0].(*models.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) GetFeed(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeed", reflect.TypeOf((*MockUseCase)(nil).GetAuthors), session)
}

func (m *MockUseCase) Check(session string, authorId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", session, authorId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockUseCase) Check(session interface{}, authorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockUseCase)(nil).Check), session, authorId)
}