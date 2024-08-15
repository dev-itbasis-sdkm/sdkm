// Code generated by MockGen. DO NOT EDIT.
// Source: sdk-versions.go
//
// Generated by this command:
//
//	mockgen -source=sdk-versions.go -package=sdkversion -destination=sdk-versions.mock.go
//

// Package sdkversion is a generated GoMock package.
package sdkversion

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSDKVersions is a mock of SDKVersions interface.
type MockSDKVersions struct {
	ctrl     *gomock.Controller
	recorder *MockSDKVersionsMockRecorder
}

// MockSDKVersionsMockRecorder is the mock recorder for MockSDKVersions.
type MockSDKVersionsMockRecorder struct {
	mock *MockSDKVersions
}

// NewMockSDKVersions creates a new mock instance.
func NewMockSDKVersions(ctrl *gomock.Controller) *MockSDKVersions {
	mock := &MockSDKVersions{ctrl: ctrl}
	mock.recorder = &MockSDKVersionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSDKVersions) EXPECT() *MockSDKVersionsMockRecorder {
	return m.recorder
}

// AllVersions mocks base method.
func (m *MockSDKVersions) AllVersions(ctx context.Context) []SDKVersion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllVersions", ctx)
	ret0, _ := ret[0].([]SDKVersion)
	return ret0
}

// AllVersions indicates an expected call of AllVersions.
func (mr *MockSDKVersionsMockRecorder) AllVersions(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllVersions", reflect.TypeOf((*MockSDKVersions)(nil).AllVersions), ctx)
}

// LatestVersion mocks base method.
func (m *MockSDKVersions) LatestVersion(ctx context.Context) SDKVersion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestVersion", ctx)
	ret0, _ := ret[0].(SDKVersion)
	return ret0
}

// LatestVersion indicates an expected call of LatestVersion.
func (mr *MockSDKVersionsMockRecorder) LatestVersion(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestVersion", reflect.TypeOf((*MockSDKVersions)(nil).LatestVersion), ctx)
}

// WithCache mocks base method.
func (m *MockSDKVersions) WithCache(cache Cache) SDKVersions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithCache", cache)
	ret0, _ := ret[0].(SDKVersions)
	return ret0
}

// WithCache indicates an expected call of WithCache.
func (mr *MockSDKVersionsMockRecorder) WithCache(cache any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithCache", reflect.TypeOf((*MockSDKVersions)(nil).WithCache), cache)
}
