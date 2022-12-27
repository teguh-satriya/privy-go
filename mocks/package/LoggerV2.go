// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LoggerV2 is an autogenerated mock type for the LoggerV2 type
type LoggerV2 struct {
	mock.Mock
}

// Error provides a mock function with given fields: args
func (_m *LoggerV2) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: format, args
func (_m *LoggerV2) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errorln provides a mock function with given fields: args
func (_m *LoggerV2) Errorln(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: args
func (_m *LoggerV2) Fatal(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatalf provides a mock function with given fields: format, args
func (_m *LoggerV2) Fatalf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatalln provides a mock function with given fields: args
func (_m *LoggerV2) Fatalln(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: args
func (_m *LoggerV2) Info(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Infof provides a mock function with given fields: format, args
func (_m *LoggerV2) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Infoln provides a mock function with given fields: args
func (_m *LoggerV2) Infoln(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// V provides a mock function with given fields: l
func (_m *LoggerV2) V(l int) bool {
	ret := _m.Called(l)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(l)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Warning provides a mock function with given fields: args
func (_m *LoggerV2) Warning(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warningf provides a mock function with given fields: format, args
func (_m *LoggerV2) Warningf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warningln provides a mock function with given fields: args
func (_m *LoggerV2) Warningln(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

type mockConstructorTestingTNewLoggerV2 interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoggerV2 creates a new instance of LoggerV2. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoggerV2(t mockConstructorTestingTNewLoggerV2) *LoggerV2 {
	mock := &LoggerV2{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}