// Code generated by MockGen. DO NOT EDIT.
// Source: D:/golang/src/telkomsel/ticket/user-service/config/logger.go
//
// Generated by this command:
//
//	mockgen -source=D:/golang/src/telkomsel/ticket/user-service/config/logger.go -destination=D:/golang/src/telkomsel/ticket/user-service/mock/config/logger_mock.go
//

// Package mock_config is a generated GoMock package.
package mock_config

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	zapcore "go.uber.org/zap/zapcore"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MockLogger) Error(msg string, fields ...zapcore.Field) {
	m.ctrl.T.Helper()
	varargs := []any{msg}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(msg any, fields ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{msg}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), varargs...)
}

// Info mocks base method.
func (m *MockLogger) Info(msg string, fields ...zapcore.Field) {
	m.ctrl.T.Helper()
	varargs := []any{msg}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(msg any, fields ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{msg}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), varargs...)
}