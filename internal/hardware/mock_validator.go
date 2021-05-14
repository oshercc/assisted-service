// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go

// Package hardware is a generated GoMock package.
package hardware

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	models "github.com/openshift/assisted-service/models"
	reflect "reflect"
)

// MockValidator is a mock of Validator interface
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// GetHostValidDisks mocks base method
func (m *MockValidator) GetHostValidDisks(host *models.Host) ([]*models.Disk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostValidDisks", host)
	ret0, _ := ret[0].([]*models.Disk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostValidDisks indicates an expected call of GetHostValidDisks
func (mr *MockValidatorMockRecorder) GetHostValidDisks(host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostValidDisks", reflect.TypeOf((*MockValidator)(nil).GetHostValidDisks), host)
}

// GetHostInstallationPath mocks base method
func (m *MockValidator) GetHostInstallationPath(host *models.Host) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostInstallationPath", host)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHostInstallationPath indicates an expected call of GetHostInstallationPath
func (mr *MockValidatorMockRecorder) GetHostInstallationPath(host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostInstallationPath", reflect.TypeOf((*MockValidator)(nil).GetHostInstallationPath), host)
}

// GetClusterHostRequirements mocks base method
func (m *MockValidator) GetClusterHostRequirements(ctx context.Context, cluster *common.Cluster, host *models.Host) (*models.ClusterHostRequirements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterHostRequirements", ctx, cluster, host)
	ret0, _ := ret[0].(*models.ClusterHostRequirements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterHostRequirements indicates an expected call of GetClusterHostRequirements
func (mr *MockValidatorMockRecorder) GetClusterHostRequirements(ctx, cluster, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterHostRequirements", reflect.TypeOf((*MockValidator)(nil).GetClusterHostRequirements), ctx, cluster, host)
}

// DiskIsEligible mocks base method
func (m *MockValidator) DiskIsEligible(ctx context.Context, disk *models.Disk, cluster *common.Cluster, host *models.Host) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiskIsEligible", ctx, disk, cluster, host)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiskIsEligible indicates an expected call of DiskIsEligible
func (mr *MockValidatorMockRecorder) DiskIsEligible(ctx, disk, cluster, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiskIsEligible", reflect.TypeOf((*MockValidator)(nil).DiskIsEligible), ctx, disk, cluster, host)
}

// ListEligibleDisks mocks base method
func (m *MockValidator) ListEligibleDisks(inventory *models.Inventory) []*models.Disk {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEligibleDisks", inventory)
	ret0, _ := ret[0].([]*models.Disk)
	return ret0
}

// ListEligibleDisks indicates an expected call of ListEligibleDisks
func (mr *MockValidatorMockRecorder) ListEligibleDisks(inventory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEligibleDisks", reflect.TypeOf((*MockValidator)(nil).ListEligibleDisks), inventory)
}

// GetInstallationDiskSpeedThresholdMs mocks base method
func (m *MockValidator) GetInstallationDiskSpeedThresholdMs(ctx context.Context, cluster *common.Cluster, host *models.Host) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstallationDiskSpeedThresholdMs", ctx, cluster, host)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstallationDiskSpeedThresholdMs indicates an expected call of GetInstallationDiskSpeedThresholdMs
func (mr *MockValidatorMockRecorder) GetInstallationDiskSpeedThresholdMs(ctx, cluster, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstallationDiskSpeedThresholdMs", reflect.TypeOf((*MockValidator)(nil).GetInstallationDiskSpeedThresholdMs), ctx, cluster, host)
}

// GetPreflightHardwareRequirements mocks base method
func (m *MockValidator) GetPreflightHardwareRequirements(ctx context.Context, cluster *common.Cluster) (*models.PreflightHardwareRequirements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreflightHardwareRequirements", ctx, cluster)
	ret0, _ := ret[0].(*models.PreflightHardwareRequirements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreflightHardwareRequirements indicates an expected call of GetPreflightHardwareRequirements
func (mr *MockValidatorMockRecorder) GetPreflightHardwareRequirements(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreflightHardwareRequirements", reflect.TypeOf((*MockValidator)(nil).GetPreflightHardwareRequirements), ctx, cluster)
}

// GetDefaultVersionRequirements mocks base method
func (m *MockValidator) GetDefaultVersionRequirements(singleNode bool) (*models.VersionedHostRequirements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultVersionRequirements", singleNode)
	ret0, _ := ret[0].(*models.VersionedHostRequirements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultVersionRequirements indicates an expected call of GetDefaultVersionRequirements
func (mr *MockValidatorMockRecorder) GetDefaultVersionRequirements(singleNode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultVersionRequirements", reflect.TypeOf((*MockValidator)(nil).GetDefaultVersionRequirements), singleNode)
}
