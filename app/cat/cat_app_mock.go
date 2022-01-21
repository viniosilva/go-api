// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/viniosilva/go-api/app/cat (interfaces: CatApp)

// Package cat is a generated GoMock package.
package cat

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCatApp is a mock of CatApp interface.
type MockCatApp struct {
	ctrl     *gomock.Controller
	recorder *MockCatAppMockRecorder
}

// MockCatAppMockRecorder is the mock recorder for MockCatApp.
type MockCatAppMockRecorder struct {
	mock *MockCatApp
}

// NewMockCatApp creates a new mock instance.
func NewMockCatApp(ctrl *gomock.Controller) *MockCatApp {
	mock := &MockCatApp{ctrl: ctrl}
	mock.recorder = &MockCatAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCatApp) EXPECT() *MockCatAppMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCatApp) Create(arg0 CreateDto) (CreateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(CreateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCatAppMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCatApp)(nil).Create), arg0)
}

// List mocks base method.
func (m *MockCatApp) List() (ListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(ListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCatAppMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCatApp)(nil).List))
}
