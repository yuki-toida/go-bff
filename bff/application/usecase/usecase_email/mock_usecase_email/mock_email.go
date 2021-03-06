// Code generated by MockGen. DO NOT EDIT.
// Source: go-bff/bff/application/usecase/usecase_email (interfaces: UseCase)

// Package mock_usecase_email is a generated GoMock package.
package mock_usecase_email

import (
	gomock "github.com/golang/mock/gomock"
	entity_email "go-bff/bff/domain/entities/entity_email"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// First mocks base method
func (m *MockUseCase) First(arg0 uint64) (*entity_email.Email, error) {
	ret := m.ctrl.Call(m, "First", arg0)
	ret0, _ := ret[0].(*entity_email.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// First indicates an expected call of First
func (mr *MockUseCaseMockRecorder) First(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockUseCase)(nil).First), arg0)
}

// Update mocks base method
func (m *MockUseCase) Update(arg0 uint64, arg1 string) (*entity_email.Email, error) {
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*entity_email.Email)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUseCaseMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), arg0, arg1)
}
