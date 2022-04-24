// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erda-project/erda/modules/scheduler/executor/plugins/k8s/addon (interfaces: NamespaceUtil)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNamespaceUtil is a mock of NamespaceUtil interface.
type MockNamespaceUtil struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceUtilMockRecorder
}

// MockNamespaceUtilMockRecorder is the mock recorder for MockNamespaceUtil.
type MockNamespaceUtilMockRecorder struct {
	mock *MockNamespaceUtil
}

// NewMockNamespaceUtil creates a new mock instance.
func NewMockNamespaceUtil(ctrl *gomock.Controller) *MockNamespaceUtil {
	mock := &MockNamespaceUtil{ctrl: ctrl}
	mock.recorder = &MockNamespaceUtilMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceUtil) EXPECT() *MockNamespaceUtilMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockNamespaceUtil) Create(arg0 string, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockNamespaceUtilMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNamespaceUtil)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockNamespaceUtil) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNamespaceUtilMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNamespaceUtil)(nil).Delete), arg0)
}

// Exists mocks base method.
func (m *MockNamespaceUtil) Exists(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockNamespaceUtilMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockNamespaceUtil)(nil).Exists), arg0)
}