// Code generated by MockGen. DO NOT EDIT.
// Source: go-bff/bff/adapter/microservices/microservice_email (interfaces: Service)

// Package mock_microservice_email is a generated GoMock package.
package mock_microservice_email

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	pb "go-bff/email/pb"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Build mocks base method
func (m *MockService) Build(arg0 context.Context, arg1 *pb.BuildRequest) (*pb.BuildResponse, error) {
	ret := m.ctrl.Call(m, "Build", arg0, arg1)
	ret0, _ := ret[0].(*pb.BuildResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build
func (mr *MockServiceMockRecorder) Build(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockService)(nil).Build), arg0, arg1)
}

// Reverse mocks base method
func (m *MockService) Reverse(arg0 context.Context, arg1 *pb.ReverseRequest) (*pb.ReverseResponse, error) {
	ret := m.ctrl.Call(m, "Reverse", arg0, arg1)
	ret0, _ := ret[0].(*pb.ReverseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reverse indicates an expected call of Reverse
func (mr *MockServiceMockRecorder) Reverse(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reverse", reflect.TypeOf((*MockService)(nil).Reverse), arg0, arg1)
}
