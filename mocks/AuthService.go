// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"
	gopherql "github.com/shohinsan/GopherQL"

	mock "github.com/stretchr/testify/mock"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, input
func (_m *AuthService) Login(ctx context.Context, input gopherql.LoginInput) (gopherql.AuthResponse, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 gopherql.AuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, gopherql.LoginInput) (gopherql.AuthResponse, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, gopherql.LoginInput) gopherql.AuthResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(gopherql.AuthResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, gopherql.LoginInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, input
func (_m *AuthService) Register(ctx context.Context, input gopherql.RegisterInput) (gopherql.AuthResponse, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 gopherql.AuthResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, gopherql.RegisterInput) (gopherql.AuthResponse, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, gopherql.RegisterInput) gopherql.AuthResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(gopherql.AuthResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, gopherql.RegisterInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}