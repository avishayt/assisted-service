// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/pkg/job (interfaces: LocalJob)

// Package job is a generated GoMock package.
package job

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockLocalJob is a mock of LocalJob interface
type MockLocalJob struct {
	ctrl     *gomock.Controller
	recorder *MockLocalJobMockRecorder
}

// MockLocalJobMockRecorder is the mock recorder for MockLocalJob
type MockLocalJobMockRecorder struct {
	mock *MockLocalJob
}

// NewMockLocalJob creates a new mock instance
func NewMockLocalJob(ctrl *gomock.Controller) *MockLocalJob {
	mock := &MockLocalJob{ctrl: ctrl}
	mock.recorder = &MockLocalJobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocalJob) EXPECT() *MockLocalJobMockRecorder {
	return m.recorder
}

// AbortInstallConfig mocks base method
func (m *MockLocalJob) AbortInstallConfig(arg0 context.Context, arg1 common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortInstallConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortInstallConfig indicates an expected call of AbortInstallConfig
func (mr *MockLocalJobMockRecorder) AbortInstallConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortInstallConfig", reflect.TypeOf((*MockLocalJob)(nil).AbortInstallConfig), arg0, arg1)
}

// Execute mocks base method
func (m *MockLocalJob) Execute(arg0, arg1 string, arg2 []string, arg3 logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockLocalJobMockRecorder) Execute(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockLocalJob)(nil).Execute), arg0, arg1, arg2, arg3)
}

// GenerateInstallConfig mocks base method
func (m *MockLocalJob) GenerateInstallConfig(arg0 context.Context, arg1 common.Cluster, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateInstallConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateInstallConfig indicates an expected call of GenerateInstallConfig
func (mr *MockLocalJobMockRecorder) GenerateInstallConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateInstallConfig", reflect.TypeOf((*MockLocalJob)(nil).GenerateInstallConfig), arg0, arg1, arg2)
}

// UploadBaseISO mocks base method
func (m *MockLocalJob) UploadBaseISO() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadBaseISO")
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadBaseISO indicates an expected call of UploadBaseISO
func (mr *MockLocalJobMockRecorder) UploadBaseISO() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadBaseISO", reflect.TypeOf((*MockLocalJob)(nil).UploadBaseISO))
}
