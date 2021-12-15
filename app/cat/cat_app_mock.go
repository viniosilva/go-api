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

// CreateCat mocks base method.
func (m *MockCatApp) CreateCat(arg0 CatCmd) (CatDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCat", arg0)
	ret0, _ := ret[0].(CatDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCat indicates an expected call of CreateCat.
func (mr *MockCatAppMockRecorder) CreateCat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCat", reflect.TypeOf((*MockCatApp)(nil).CreateCat), arg0)
}

// FindCats mocks base method.
func (m *MockCatApp) FindCats() ListCatsDto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCats")
	ret0, _ := ret[0].(ListCatsDto)
	return ret0
}

// FindCats indicates an expected call of FindCats.
func (mr *MockCatAppMockRecorder) FindCats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCats", reflect.TypeOf((*MockCatApp)(nil).FindCats))
}
