// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	comment "github.com/garindradeksa/socialmedia-mini/features/comment"
	mock "github.com/stretchr/testify/mock"
)

// CommentService is an autogenerated mock type for the CommentService type
type CommentService struct {
	mock.Mock
}

// Add provides a mock function with given fields: token, newComment, contentID
func (_m *CommentService) Add(token interface{}, newComment comment.Core, contentID uint) (comment.Core, error) {
	ret := _m.Called(token, newComment, contentID)

	var r0 comment.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, comment.Core, uint) (comment.Core, error)); ok {
		return rf(token, newComment, contentID)
	}
	if rf, ok := ret.Get(0).(func(interface{}, comment.Core, uint) comment.Core); ok {
		r0 = rf(token, newComment, contentID)
	} else {
		r0 = ret.Get(0).(comment.Core)
	}

	if rf, ok := ret.Get(1).(func(interface{}, comment.Core, uint) error); ok {
		r1 = rf(token, newComment, contentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: token, commentID
func (_m *CommentService) Delete(token interface{}, commentID uint) error {
	ret := _m.Called(token, commentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint) error); ok {
		r0 = rf(token, commentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCommentService creates a new instance of CommentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommentService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommentService {
	mock := &CommentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
