package subscribe

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

func (m *MockRepository) AddSubscribe(subscriberId int, authorId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubscribe", subscriberId, authorId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) AddSubscribe(subscriberId interface{}, authorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubscribe", reflect.TypeOf((*MockRepository)(nil).AddSubscribe), subscriberId, authorId)
}


func (m *MockRepository) DeleteSubscribe(subscriberId int, authorId int) error  {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubscribe", subscriberId, authorId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockRecorderMockRepository) DeleteSubscribe(subscriberId interface{}, authorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubscribe", reflect.TypeOf((*MockRepository)(nil).DeleteSubscribe), subscriberId, authorId)
}

func (m *MockRepository) GetAuthors(subscriberId int) (*models.PublicUsers, error)  {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthors", subscriberId)
	ret0, _ := ret[0].(*models.PublicUsers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) GetAuthors(subscriberId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthors", reflect.TypeOf((*MockRepository)(nil).GetAuthors), subscriberId)
}

func (m *MockRepository) GetRatingFeed(subscriberId int) (*models.Feed, error)  {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRatingFeed", subscriberId)
	ret0, _ := ret[0].(*models.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) GetRatingFeed(subscriberId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRatingFeed", reflect.TypeOf((*MockRepository)(nil).GetRatingFeed), subscriberId)
}

func (m *MockRepository) GetReviewFeed(subscriberId int) (*models.Feed, error)  {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewFeed", subscriberId)
	ret0, _ := ret[0].(*models.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) GetReviewFeed(subscriberId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewFeed", reflect.TypeOf((*MockRepository)(nil).GetReviewFeed), subscriberId)
}

func (m *MockRepository) Check(subscriberId int, authorId int) (bool, error)  {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", subscriberId, authorId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRecorderMockRepository) Check(subscriberId interface{}, authorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockRepository)(nil).Check), subscriberId, authorId)
}


