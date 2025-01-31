// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/qdm12/gluetun/internal/provider/common (interfaces: ParallelResolver,Storage,Unzipper)

// Package common is a generated GoMock package.
package common

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	settings "github.com/qdm12/gluetun/internal/configuration/settings"
	models "github.com/qdm12/gluetun/internal/models"
	resolver "github.com/qdm12/gluetun/internal/updater/resolver"
)

// MockParallelResolver is a mock of ParallelResolver interface.
type MockParallelResolver struct {
	ctrl     *gomock.Controller
	recorder *MockParallelResolverMockRecorder
}

// MockParallelResolverMockRecorder is the mock recorder for MockParallelResolver.
type MockParallelResolverMockRecorder struct {
	mock *MockParallelResolver
}

// NewMockParallelResolver creates a new mock instance.
func NewMockParallelResolver(ctrl *gomock.Controller) *MockParallelResolver {
	mock := &MockParallelResolver{ctrl: ctrl}
	mock.recorder = &MockParallelResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParallelResolver) EXPECT() *MockParallelResolverMockRecorder {
	return m.recorder
}

// Resolve mocks base method.
func (m *MockParallelResolver) Resolve(arg0 context.Context, arg1 resolver.ParallelSettings) (map[string][]net.IP, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", arg0, arg1)
	ret0, _ := ret[0].(map[string][]net.IP)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Resolve indicates an expected call of Resolve.
func (mr *MockParallelResolverMockRecorder) Resolve(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockParallelResolver)(nil).Resolve), arg0, arg1)
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// FilterServers mocks base method.
func (m *MockStorage) FilterServers(arg0 string, arg1 settings.ServerSelection) ([]models.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterServers", arg0, arg1)
	ret0, _ := ret[0].([]models.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterServers indicates an expected call of FilterServers.
func (mr *MockStorageMockRecorder) FilterServers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterServers", reflect.TypeOf((*MockStorage)(nil).FilterServers), arg0, arg1)
}

// GetServerByName mocks base method.
func (m *MockStorage) GetServerByName(arg0, arg1 string) (models.Server, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerByName", arg0, arg1)
	ret0, _ := ret[0].(models.Server)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetServerByName indicates an expected call of GetServerByName.
func (mr *MockStorageMockRecorder) GetServerByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerByName", reflect.TypeOf((*MockStorage)(nil).GetServerByName), arg0, arg1)
}

// MockUnzipper is a mock of Unzipper interface.
type MockUnzipper struct {
	ctrl     *gomock.Controller
	recorder *MockUnzipperMockRecorder
}

// MockUnzipperMockRecorder is the mock recorder for MockUnzipper.
type MockUnzipperMockRecorder struct {
	mock *MockUnzipper
}

// NewMockUnzipper creates a new mock instance.
func NewMockUnzipper(ctrl *gomock.Controller) *MockUnzipper {
	mock := &MockUnzipper{ctrl: ctrl}
	mock.recorder = &MockUnzipperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnzipper) EXPECT() *MockUnzipperMockRecorder {
	return m.recorder
}

// FetchAndExtract mocks base method.
func (m *MockUnzipper) FetchAndExtract(arg0 context.Context, arg1 string) (map[string][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAndExtract", arg0, arg1)
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAndExtract indicates an expected call of FetchAndExtract.
func (mr *MockUnzipperMockRecorder) FetchAndExtract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAndExtract", reflect.TypeOf((*MockUnzipper)(nil).FetchAndExtract), arg0, arg1)
}
