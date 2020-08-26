// Code generated by MockGen. DO NOT EDIT.
// Source: local_job.go

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

// Execute mocks base method
func (m *MockLocalJob) Execute(pythonCommand, pythonFilePath string, envVars []string, log logrus.FieldLogger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", pythonCommand, pythonFilePath, envVars, log)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockLocalJobMockRecorder) Execute(pythonCommand, pythonFilePath, envVars, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockLocalJob)(nil).Execute), pythonCommand, pythonFilePath, envVars, log)
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

// GenerateInstallConfig mocks base method
func (m *MockLocalJob) GenerateInstallConfig(ctx context.Context, cluster common.Cluster, cfg []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateInstallConfig", ctx, cluster, cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateInstallConfig indicates an expected call of GenerateInstallConfig
func (mr *MockLocalJobMockRecorder) GenerateInstallConfig(ctx, cluster, cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateInstallConfig", reflect.TypeOf((*MockLocalJob)(nil).GenerateInstallConfig), ctx, cluster, cfg)
}

// AbortInstallConfig mocks base method
func (m *MockLocalJob) AbortInstallConfig(ctx context.Context, cluster common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortInstallConfig", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortInstallConfig indicates an expected call of AbortInstallConfig
func (mr *MockLocalJobMockRecorder) AbortInstallConfig(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortInstallConfig", reflect.TypeOf((*MockLocalJob)(nil).AbortInstallConfig), ctx, cluster)
}
