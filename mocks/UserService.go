// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	gopherql "github.com/shohinsan/GopherQL"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *UserService) GetByID(ctx context.Context, id string) (gopherql.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 gopherql.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (gopherql.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) gopherql.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(gopherql.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
