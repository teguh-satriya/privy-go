// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeCakesServiceServer is an autogenerated mock type for the UnsafeCakesServiceServer type
type UnsafeCakesServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedCakesServiceServer provides a mock function with given fields:
func (_m *UnsafeCakesServiceServer) mustEmbedUnimplementedCakesServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafeCakesServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeCakesServiceServer creates a new instance of UnsafeCakesServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeCakesServiceServer(t mockConstructorTestingTNewUnsafeCakesServiceServer) *UnsafeCakesServiceServer {
	mock := &UnsafeCakesServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
