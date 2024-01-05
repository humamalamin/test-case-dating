// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/auth/jwt/jwt.go

// Package jwtAuth is a generated GoMock package.
package jwtAuth

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockJwt is a mock of Jwt interface.
type MockJwt struct {
	ctrl     *gomock.Controller
	recorder *MockJwtMockRecorder
}

// MockJwtMockRecorder is the mock recorder for MockJwt.
type MockJwtMockRecorder struct {
	mock *MockJwt
}

// NewMockJwt creates a new mock instance.
func NewMockJwt(ctrl *gomock.Controller) *MockJwt {
	mock := &MockJwt{ctrl: ctrl}
	mock.recorder = &MockJwtMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwt) EXPECT() *MockJwtMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockJwt) GenerateToken(data *JwtData) (string, string, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockJwtMockRecorder) GenerateToken(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockJwt)(nil).GenerateToken), data)
}

// VerifyAccessToken mocks base method.
func (m *MockJwt) VerifyAccessToken(token string) (*JwtData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAccessToken", token)
	ret0, _ := ret[0].(*JwtData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAccessToken indicates an expected call of VerifyAccessToken.
func (mr *MockJwtMockRecorder) VerifyAccessToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAccessToken", reflect.TypeOf((*MockJwt)(nil).VerifyAccessToken), token)
}

// VerifyRefreshToken mocks base method.
func (m *MockJwt) VerifyRefreshToken(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyRefreshToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyRefreshToken indicates an expected call of VerifyRefreshToken.
func (mr *MockJwtMockRecorder) VerifyRefreshToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyRefreshToken", reflect.TypeOf((*MockJwt)(nil).VerifyRefreshToken), token)
}
