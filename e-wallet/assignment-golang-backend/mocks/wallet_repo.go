// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "ewallet/entity"

	mock "github.com/stretchr/testify/mock"
)

// WalletRepo is an autogenerated mock type for the WalletRepo type
type WalletRepo struct {
	mock.Mock
}

// AddBalance provides a mock function with given fields: walletID, amount
func (_m *WalletRepo) AddBalance(walletID int, amount int) error {
	ret := _m.Called(walletID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(walletID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWallet provides a mock function with given fields:
func (_m *WalletRepo) CreateWallet() (*entity.Wallet, error) {
	ret := _m.Called()

	var r0 *entity.Wallet
	if rf, ok := ret.Get(0).(func() *entity.Wallet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByWalletID provides a mock function with given fields: walletid
func (_m *WalletRepo) GetUserByWalletID(walletid int) (*entity.User, error) {
	ret := _m.Called(walletid)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(int) *entity.User); ok {
		r0 = rf(walletid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(walletid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWallet provides a mock function with given fields: walletID
func (_m *WalletRepo) GetWallet(walletID int) (*entity.Wallet, error) {
	ret := _m.Called(walletID)

	var r0 *entity.Wallet
	if rf, ok := ret.Get(0).(func(int) *entity.Wallet); ok {
		r0 = rf(walletID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(walletID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReduceBalance provides a mock function with given fields: walletID, amount
func (_m *WalletRepo) ReduceBalance(walletID int, amount int) error {
	ret := _m.Called(walletID, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(walletID, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewWalletRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewWalletRepo creates a new instance of WalletRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWalletRepo(t mockConstructorTestingTNewWalletRepo) *WalletRepo {
	mock := &WalletRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}