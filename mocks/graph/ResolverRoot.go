// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	graph "github.com/shohinsan/GopherQL/graph"
	mock "github.com/stretchr/testify/mock"
)

// ResolverRoot is an autogenerated mock type for the ResolverRoot type
type ResolverRoot struct {
	mock.Mock
}

// Mutation provides a mock function with given fields:
func (_m *ResolverRoot) Mutation() graph.MutationResolver {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Mutation")
	}

	var r0 graph.MutationResolver
	if rf, ok := ret.Get(0).(func() graph.MutationResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(graph.MutationResolver)
		}
	}

	return r0
}

// Query provides a mock function with given fields:
func (_m *ResolverRoot) Query() graph.QueryResolver {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 graph.QueryResolver
	if rf, ok := ret.Get(0).(func() graph.QueryResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(graph.QueryResolver)
		}
	}

	return r0
}

// Tweet provides a mock function with given fields:
func (_m *ResolverRoot) Tweet() graph.TweetResolver {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Tweet")
	}

	var r0 graph.TweetResolver
	if rf, ok := ret.Get(0).(func() graph.TweetResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(graph.TweetResolver)
		}
	}

	return r0
}

// NewResolverRoot creates a new instance of ResolverRoot. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResolverRoot(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResolverRoot {
	mock := &ResolverRoot{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
