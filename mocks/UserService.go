// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"

	mock "github.com/stretchr/testify/mock"

	user "github.com/garindradeksa/socialmedia-mini/features/user"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Deactivate provides a mock function with given fields: token
func (_m *UserService) Deactivate(token interface{}) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: username, password
func (_m *UserService) Login(username string, password string) (string, user.Core, error) {
	ret := _m.Called(username, password)

	var r0 string
	var r1 user.Core
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (string, user.Core, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) user.Core); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Get(1).(user.Core)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(username, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Profile provides a mock function with given fields: token
func (_m *UserService) Profile(token interface{}) (user.Core, error) {
	ret := _m.Called(token)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (user.Core, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(interface{}) user.Core); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: newUser
func (_m *UserService) Register(newUser user.Core) error {
	ret := _m.Called(newUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Core) error); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: formHeader, formHeader2, token, updatedProfile
func (_m *UserService) Update(formHeader multipart.FileHeader, formHeader2 multipart.FileHeader, token interface{}, updatedProfile user.Core) (user.Core, error) {
	ret := _m.Called(formHeader, formHeader2, token, updatedProfile)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(multipart.FileHeader, multipart.FileHeader, interface{}, user.Core) (user.Core, error)); ok {
		return rf(formHeader, formHeader2, token, updatedProfile)
	}
	if rf, ok := ret.Get(0).(func(multipart.FileHeader, multipart.FileHeader, interface{}, user.Core) user.Core); ok {
		r0 = rf(formHeader, formHeader2, token, updatedProfile)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(multipart.FileHeader, multipart.FileHeader, interface{}, user.Core) error); ok {
		r1 = rf(formHeader, formHeader2, token, updatedProfile)
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
