// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	services "github.com/teguh-satriya/privy-go/services"
)

// CreateCakesService is an autogenerated mock type for the CreateCakesService type
type CreateCakesService struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, params
func (_m *CreateCakesService) Call(ctx context.Context, params *services.CreateCakeParams) (*services.CreateCakeResult, error) {
	ret := _m.Called(ctx, params)

	var r0 *services.CreateCakeResult
	if rf, ok := ret.Get(0).(func(context.Context, *services.CreateCakeParams) *services.CreateCakeResult); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.CreateCakeResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *services.CreateCakeParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCreateCakesService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateCakesService creates a new instance of CreateCakesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateCakesService(t mockConstructorTestingTNewCreateCakesService) *CreateCakesService {
	mock := &CreateCakesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
