// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "nickPay/wallet/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// WalletService is an autogenerated mock type for the WalletService type
type WalletService struct {
	mock.Mock
}

// CreditWallet provides a mock function with given fields: _a0, _a1, _a2
func (_m *WalletService) CreditWallet(_a0 context.Context, _a1 int64, _a2 float64) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, float64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DebitWallet provides a mock function with given fields: _a0, _a1, _a2
func (_m *WalletService) DebitWallet(_a0 context.Context, _a1 int64, _a2 float64) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, float64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetWallet provides a mock function with given fields: _a0, _a1
func (_m *WalletService) GetWallet(_a0 context.Context, _a1 int64) (domain.Wallet, error) {
	ret := _m.Called(_a0, _a1)

	var r0 domain.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.Wallet, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Wallet); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(domain.Wallet)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: _a0, _a1
func (_m *WalletService) LoginUser(_a0 context.Context, _a1 domain.LoginUserRequest) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.LoginUserRequest) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.LoginUserRequest) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.LoginUserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: _a0, _a1
func (_m *WalletService) RegisterUser(_a0 context.Context, _a1 domain.User) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewWalletService interface {
	mock.TestingT
	Cleanup(func())
}

// NewWalletService creates a new instance of WalletService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWalletService(t mockConstructorTestingTNewWalletService) *WalletService {
	mock := &WalletService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
