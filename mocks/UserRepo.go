// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	gopherql "github.com/shohinsan/GopherQL"
	mock "github.com/stretchr/testify/mock"
)

// UserRepo is an autogenerated mock type for the UserRepo type
type UserRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *UserRepo) Create(ctx context.Context, user gopherql.User) (gopherql.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 gopherql.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, gopherql.User) (gopherql.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, gopherql.User) gopherql.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(gopherql.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, gopherql.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepo) GetByEmail(ctx context.Context, email string) (gopherql.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 gopherql.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (gopherql.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) gopherql.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(gopherql.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *UserRepo) GetByID(ctx context.Context, id string) (gopherql.User, error) {
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

// GetByIds provides a mock function with given fields: ctx, ids
func (_m *UserRepo) GetByIds(ctx context.Context, ids []string) ([]gopherql.User, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for GetByIds")
	}

	var r0 []gopherql.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]gopherql.User, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []gopherql.User); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gopherql.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *UserRepo) GetByUsername(ctx context.Context, username string) (gopherql.User, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetByUsername")
	}

	var r0 gopherql.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (gopherql.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) gopherql.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(gopherql.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepo creates a new instance of UserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepo {
	mock := &UserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
