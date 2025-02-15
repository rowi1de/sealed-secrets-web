// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bakito/sealed-secrets-web/pkg/seal (interfaces: Sealer)

// Package seal is a generated GoMock package.
package seal

import (
	context "context"
	io "io"
	reflect "reflect"

	seal "github.com/bakito/sealed-secrets-web/pkg/seal"
	gomock "github.com/golang/mock/gomock"
)

// MockSealer is a mock of Sealer interface.
type MockSealer struct {
	ctrl     *gomock.Controller
	recorder *MockSealerMockRecorder
}

// MockSealerMockRecorder is the mock recorder for MockSealer.
type MockSealerMockRecorder struct {
	mock *MockSealer
}

// NewMockSealer creates a new mock instance.
func NewMockSealer(ctrl *gomock.Controller) *MockSealer {
	mock := &MockSealer{ctrl: ctrl}
	mock.recorder = &MockSealerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSealer) EXPECT() *MockSealerMockRecorder {
	return m.recorder
}

// Certificate mocks base method.
func (m *MockSealer) Certificate(arg0 context.Context) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Certificate", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Certificate indicates an expected call of Certificate.
func (mr *MockSealerMockRecorder) Certificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Certificate", reflect.TypeOf((*MockSealer)(nil).Certificate), arg0)
}

// Raw mocks base method.
func (m *MockSealer) Raw(arg0 seal.Raw) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Raw indicates an expected call of Raw.
func (mr *MockSealerMockRecorder) Raw(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockSealer)(nil).Raw), arg0)
}

// Seal mocks base method.
func (m *MockSealer) Seal(arg0 string, arg1 io.Reader) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seal indicates an expected call of Seal.
func (mr *MockSealerMockRecorder) Seal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockSealer)(nil).Seal), arg0, arg1)
}

// Validate mocks base method.
func (m *MockSealer) Validate(arg0 context.Context, arg1 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockSealerMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockSealer)(nil).Validate), arg0, arg1)
}
