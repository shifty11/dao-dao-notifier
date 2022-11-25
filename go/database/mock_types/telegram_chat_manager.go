// Code generated by MockGen. DO NOT EDIT.
// Source: go/database/telegram_chat_manager.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/shifty11/cosmos-notifier/ent"
	types "github.com/shifty11/cosmos-notifier/types"
)

// MockITelegramChatManager is a mock of ITelegramChatManager interface.
type MockITelegramChatManager struct {
	ctrl     *gomock.Controller
	recorder *MockITelegramChatManagerMockRecorder
}

// MockITelegramChatManagerMockRecorder is the mock recorder for MockITelegramChatManager.
type MockITelegramChatManagerMockRecorder struct {
	mock *MockITelegramChatManager
}

// NewMockITelegramChatManager creates a new mock instance.
func NewMockITelegramChatManager(ctrl *gomock.Controller) *MockITelegramChatManager {
	mock := &MockITelegramChatManager{ctrl: ctrl}
	mock.recorder = &MockITelegramChatManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITelegramChatManager) EXPECT() *MockITelegramChatManagerMockRecorder {
	return m.recorder
}

// AddOrRemoveChain mocks base method.
func (m *MockITelegramChatManager) AddOrRemoveChain(tgChatId int64, chainId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrRemoveChain", tgChatId, chainId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrRemoveChain indicates an expected call of AddOrRemoveChain.
func (mr *MockITelegramChatManagerMockRecorder) AddOrRemoveChain(tgChatId, chainId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrRemoveChain", reflect.TypeOf((*MockITelegramChatManager)(nil).AddOrRemoveChain), tgChatId, chainId)
}

// AddOrRemoveContract mocks base method.
func (m *MockITelegramChatManager) AddOrRemoveContract(tgChatId int64, contractId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrRemoveContract", tgChatId, contractId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrRemoveContract indicates an expected call of AddOrRemoveContract.
func (mr *MockITelegramChatManagerMockRecorder) AddOrRemoveContract(tgChatId, contractId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrRemoveContract", reflect.TypeOf((*MockITelegramChatManager)(nil).AddOrRemoveContract), tgChatId, contractId)
}

// CountSubscriptions mocks base method.
func (m *MockITelegramChatManager) CountSubscriptions(chatId int64) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountSubscriptions", chatId)
	ret0, _ := ret[0].(int)
	return ret0
}

// CountSubscriptions indicates an expected call of CountSubscriptions.
func (mr *MockITelegramChatManagerMockRecorder) CountSubscriptions(chatId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountSubscriptions", reflect.TypeOf((*MockITelegramChatManager)(nil).CountSubscriptions), chatId)
}

// CreateOrUpdateChat mocks base method.
func (m *MockITelegramChatManager) CreateOrUpdateChat(userId int64, userName string, tgChatId int64, name string, isGroup bool) (*ent.TelegramChat, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateChat", userId, userName, tgChatId, name, isGroup)
	ret0, _ := ret[0].(*ent.TelegramChat)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CreateOrUpdateChat indicates an expected call of CreateOrUpdateChat.
func (mr *MockITelegramChatManagerMockRecorder) CreateOrUpdateChat(userId, userName, tgChatId, name, isGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateChat", reflect.TypeOf((*MockITelegramChatManager)(nil).CreateOrUpdateChat), userId, userName, tgChatId, name, isGroup)
}

// Delete mocks base method.
func (m *MockITelegramChatManager) Delete(userId, chatId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, chatId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockITelegramChatManagerMockRecorder) Delete(userId, chatId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITelegramChatManager)(nil).Delete), userId, chatId)
}

// DeleteMultiple mocks base method.
func (m *MockITelegramChatManager) DeleteMultiple(chatIds []int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteMultiple", chatIds)
}

// DeleteMultiple indicates an expected call of DeleteMultiple.
func (mr *MockITelegramChatManagerMockRecorder) DeleteMultiple(chatIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiple", reflect.TypeOf((*MockITelegramChatManager)(nil).DeleteMultiple), chatIds)
}

// GetAllIds mocks base method.
func (m *MockITelegramChatManager) GetAllIds() []types.TgChatQueryResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllIds")
	ret0, _ := ret[0].([]types.TgChatQueryResult)
	return ret0
}

// GetAllIds indicates an expected call of GetAllIds.
func (mr *MockITelegramChatManagerMockRecorder) GetAllIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllIds", reflect.TypeOf((*MockITelegramChatManager)(nil).GetAllIds))
}

// GetChatUsers mocks base method.
func (m *MockITelegramChatManager) GetChatUsers(chatId int64) []*ent.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatUsers", chatId)
	ret0, _ := ret[0].([]*ent.User)
	return ret0
}

// GetChatUsers indicates an expected call of GetChatUsers.
func (mr *MockITelegramChatManagerMockRecorder) GetChatUsers(chatId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatUsers", reflect.TypeOf((*MockITelegramChatManager)(nil).GetChatUsers), chatId)
}

// GetSubscribedIds mocks base method.
func (m *MockITelegramChatManager) GetSubscribedIds(query *ent.TelegramChatQuery) []types.TgChatQueryResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribedIds", query)
	ret0, _ := ret[0].([]types.TgChatQueryResult)
	return ret0
}

// GetSubscribedIds indicates an expected call of GetSubscribedIds.
func (mr *MockITelegramChatManagerMockRecorder) GetSubscribedIds(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribedIds", reflect.TypeOf((*MockITelegramChatManager)(nil).GetSubscribedIds), query)
}
