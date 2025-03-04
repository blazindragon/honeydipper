// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/dipper/rpc.go

// Package mock_dipper is a generated GoMock package.
package mock_dipper

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockRPCCallerStub is a mock of RPCCallerStub interface
type MockRPCCallerStub struct {
	ctrl     *gomock.Controller
	recorder *MockRPCCallerStubMockRecorder
}

// MockRPCCallerStubMockRecorder is the mock recorder for MockRPCCallerStub
type MockRPCCallerStubMockRecorder struct {
	mock *MockRPCCallerStub
}

// NewMockRPCCallerStub creates a new mock instance
func NewMockRPCCallerStub(ctrl *gomock.Controller) *MockRPCCallerStub {
	mock := &MockRPCCallerStub{ctrl: ctrl}
	mock.recorder = &MockRPCCallerStubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRPCCallerStub) EXPECT() *MockRPCCallerStubMockRecorder {
	return m.recorder
}

// GetName mocks base method
func (m *MockRPCCallerStub) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (mr *MockRPCCallerStubMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockRPCCallerStub)(nil).GetName))
}

// GetStream mocks base method
func (m *MockRPCCallerStub) GetStream(feature string) io.Writer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStream", feature)
	ret0, _ := ret[0].(io.Writer)
	return ret0
}

// GetStream indicates an expected call of GetStream
func (mr *MockRPCCallerStubMockRecorder) GetStream(feature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStream", reflect.TypeOf((*MockRPCCallerStub)(nil).GetStream), feature)
}
