// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/bryansandoval/projects/apple-coding-challenge/services/names/interface.go

// Package services_mocks is a generated GoMock package.
package services_mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mock_names is a mock of Service interface
type Mock_names struct {
	ctrl     *gomock.Controller
	recorder *Mock_namesMockRecorder
}

// Mock_namesMockRecorder is the mock recorder for Mock_names
type Mock_namesMockRecorder struct {
	mock *Mock_names
}

// NewMock_names creates a new mock instance
func NewMock_names(ctrl *gomock.Controller) *Mock_names {
	mock := &Mock_names{ctrl: ctrl}
	mock.recorder = &Mock_namesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mock_names) EXPECT() *Mock_namesMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *Mock_names) Get() (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *Mock_namesMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Mock_names)(nil).Get))
}
