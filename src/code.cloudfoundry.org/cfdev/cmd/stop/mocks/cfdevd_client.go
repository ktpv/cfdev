// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/stop (interfaces: CfdevdClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCfdevdClient is a mock of CfdevdClient interface
type MockCfdevdClient struct {
	ctrl     *gomock.Controller
	recorder *MockCfdevdClientMockRecorder
}

// MockCfdevdClientMockRecorder is the mock recorder for MockCfdevdClient
type MockCfdevdClientMockRecorder struct {
	mock *MockCfdevdClient
}

// NewMockCfdevdClient creates a new mock instance
func NewMockCfdevdClient(ctrl *gomock.Controller) *MockCfdevdClient {
	mock := &MockCfdevdClient{ctrl: ctrl}
	mock.recorder = &MockCfdevdClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCfdevdClient) EXPECT() *MockCfdevdClientMockRecorder {
	return m.recorder
}

// Uninstall mocks base method
func (m *MockCfdevdClient) Uninstall() (string, error) {
	ret := m.ctrl.Call(m, "Uninstall")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Uninstall indicates an expected call of Uninstall
func (mr *MockCfdevdClientMockRecorder) Uninstall() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockCfdevdClient)(nil).Uninstall))
}
