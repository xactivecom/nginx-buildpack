// Code generated by MockGen. DO NOT EDIT.
// Source: dynatrace.go

// Package hooks_test is a generated GoMock package.
package hooks_test

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockCommand is a mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockCommand) Execute(arg0 string, arg1, arg2 io.Writer, arg3 string, arg4 ...string) error {
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockCommandMockRecorder) Execute(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCommand)(nil).Execute), varargs...)
}
