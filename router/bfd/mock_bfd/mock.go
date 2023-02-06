// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/router/bfd (interfaces: Source,IntervalGenerator)

// Package mock_bfd is a generated GoMock package.
package mock_bfd

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSource is a mock of Source interface.
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMockRecorder
}

// MockSourceMockRecorder is the mock recorder for MockSource.
type MockSourceMockRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance.
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSource) EXPECT() *MockSourceMockRecorder {
	return m.recorder
}

// Intn mocks base method.
func (m *MockSource) Intn(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intn", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// Intn indicates an expected call of Intn.
func (mr *MockSourceMockRecorder) Intn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intn", reflect.TypeOf((*MockSource)(nil).Intn), arg0)
}

// MockIntervalGenerator is a mock of IntervalGenerator interface.
type MockIntervalGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIntervalGeneratorMockRecorder
}

// MockIntervalGeneratorMockRecorder is the mock recorder for MockIntervalGenerator.
type MockIntervalGeneratorMockRecorder struct {
	mock *MockIntervalGenerator
}

// NewMockIntervalGenerator creates a new mock instance.
func NewMockIntervalGenerator(ctrl *gomock.Controller) *MockIntervalGenerator {
	mock := &MockIntervalGenerator{ctrl: ctrl}
	mock.recorder = &MockIntervalGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntervalGenerator) EXPECT() *MockIntervalGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockIntervalGenerator) Generate(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockIntervalGeneratorMockRecorder) Generate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockIntervalGenerator)(nil).Generate), arg0, arg1)
}
