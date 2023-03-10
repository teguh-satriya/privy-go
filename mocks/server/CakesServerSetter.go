// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	server "github.com/teguh-satriya/privy-go/server"
)

// CakesServerSetter is an autogenerated mock type for the CakesServerSetter type
type CakesServerSetter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *CakesServerSetter) Execute(_a0 *server.CakesServer) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewCakesServerSetter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCakesServerSetter creates a new instance of CakesServerSetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCakesServerSetter(t mockConstructorTestingTNewCakesServerSetter) *CakesServerSetter {
	mock := &CakesServerSetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
