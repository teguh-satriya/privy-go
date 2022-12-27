// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	services "github.com/teguh-satriya/privy-go/services"
)

// DeleteCakesService is an autogenerated mock type for the DeleteCakesService type
type DeleteCakesService struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, params
func (_m *DeleteCakesService) Call(ctx context.Context, params *services.DeleteCakeParams) (*services.DeleteCakeResult, error) {
	ret := _m.Called(ctx, params)

	var r0 *services.DeleteCakeResult
	if rf, ok := ret.Get(0).(func(context.Context, *services.DeleteCakeParams) *services.DeleteCakeResult); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.DeleteCakeResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *services.DeleteCakeParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDeleteCakesService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDeleteCakesService creates a new instance of DeleteCakesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDeleteCakesService(t mockConstructorTestingTNewDeleteCakesService) *DeleteCakesService {
	mock := &DeleteCakesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
