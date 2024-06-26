// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	gopherql "github.com/shohinsan/GopherQL"
	mock "github.com/stretchr/testify/mock"
)

// TweetService is an autogenerated mock type for the TweetService type
type TweetService struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *TweetService) All(ctx context.Context) ([]gopherql.Tweet, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 []gopherql.Tweet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]gopherql.Tweet, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []gopherql.Tweet); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gopherql.Tweet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, input
func (_m *TweetService) Create(ctx context.Context, input gopherql.CreateTweetInput) (gopherql.Tweet, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 gopherql.Tweet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, gopherql.CreateTweetInput) (gopherql.Tweet, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, gopherql.CreateTweetInput) gopherql.Tweet); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(gopherql.Tweet)
	}

	if rf, ok := ret.Get(1).(func(context.Context, gopherql.CreateTweetInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReply provides a mock function with given fields: ctx, parentID, input
func (_m *TweetService) CreateReply(ctx context.Context, parentID string, input gopherql.CreateTweetInput) (gopherql.Tweet, error) {
	ret := _m.Called(ctx, parentID, input)

	if len(ret) == 0 {
		panic("no return value specified for CreateReply")
	}

	var r0 gopherql.Tweet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, gopherql.CreateTweetInput) (gopherql.Tweet, error)); ok {
		return rf(ctx, parentID, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, gopherql.CreateTweetInput) gopherql.Tweet); ok {
		r0 = rf(ctx, parentID, input)
	} else {
		r0 = ret.Get(0).(gopherql.Tweet)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, gopherql.CreateTweetInput) error); ok {
		r1 = rf(ctx, parentID, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TweetService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *TweetService) GetByID(ctx context.Context, id string) (gopherql.Tweet, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 gopherql.Tweet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (gopherql.Tweet, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) gopherql.Tweet); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(gopherql.Tweet)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByParentID provides a mock function with given fields: ctx, id
func (_m *TweetService) GetByParentID(ctx context.Context, id string) ([]gopherql.Tweet, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByParentID")
	}

	var r0 []gopherql.Tweet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]gopherql.Tweet, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []gopherql.Tweet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gopherql.Tweet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTweetService creates a new instance of TweetService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTweetService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TweetService {
	mock := &TweetService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
