// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	comment "github.com/garindradeksa/socialmedia-mini/features/comment"
	mock "github.com/stretchr/testify/mock"
)

// CommentData is an autogenerated mock type for the CommentData type
type CommentData struct {
	mock.Mock
}

// Add provides a mock function with given fields: userID, newComment, contentID
func (_m *CommentData) Add(userID uint, newComment comment.Core, contentID uint) (comment.Core, error) {
	ret := _m.Called(userID, newComment, contentID)

	var r0 comment.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, comment.Core, uint) (comment.Core, error)); ok {
		return rf(userID, newComment, contentID)
	}
	if rf, ok := ret.Get(0).(func(uint, comment.Core, uint) comment.Core); ok {
		r0 = rf(userID, newComment, contentID)
	} else {
		r0 = ret.Get(0).(comment.Core)
	}

	if rf, ok := ret.Get(1).(func(uint, comment.Core, uint) error); ok {
		r1 = rf(userID, newComment, contentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID, commentID
func (_m *CommentData) Delete(userID uint, commentID uint) error {
	ret := _m.Called(userID, commentID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userID, commentID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCommentData creates a new instance of CommentData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommentData(t interface {
	mock.TestingT
	Cleanup(func())
}) *CommentData {
	mock := &CommentData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
