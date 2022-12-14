// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "ewallet/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// GetDetail provides a mock function with given fields: id
func (_m *UserUsecase) GetDetail(id int) (*entity.UserDetailResponse, error) {
	ret := _m.Called(id)

	var r0 *entity.UserDetailResponse
	if rf, ok := ret.Get(0).(func(int) *entity.UserDetailResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserDetailResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email, password
func (_m *UserUsecase) Login(email string, password string) (string, error) {
	ret := _m.Called(email, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: _a0
func (_m *UserUsecase) Register(_a0 *entity.User) (*entity.UserRegisterResponse, error) {
	ret := _m.Called(_a0)

	var r0 *entity.UserRegisterResponse
	if rf, ok := ret.Get(0).(func(*entity.User) *entity.UserRegisterResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserRegisterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
