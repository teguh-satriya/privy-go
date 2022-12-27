// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	services "github.com/teguh-satriya/privy-go/services"
)

// GetCakesService is an autogenerated mock type for the GetCakesService type
type GetCakesService struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, params
func (_m *GetCakesService) Call(ctx context.Context, params *services.GetCakeParams) (*services.GetCakeResult, error) {
	ret := _m.Called(ctx, params)

	var r0 *services.GetCakeResult
	if rf, ok := ret.Get(0).(func(context.Context, *services.GetCakeParams) *services.GetCakeResult); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.GetCakeResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *services.GetCakeParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGetCakesService interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetCakesService creates a new instance of GetCakesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetCakesService(t mockConstructorTestingTNewGetCakesService) *GetCakesService {
	mock := &GetCakesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
