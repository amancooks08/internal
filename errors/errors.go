package errors

import (
	"errors"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrGenJWTToken = errors.New("error generating jwt token")
	ErrLoggingIn = errors.New("error logging in")
	ErrNoWallet = errors.New("no wallet found")
	ErrFetchingWallet = errors.New("error fetching wallet")
	ErrCreditingWallet = errors.New("error crediting wallet")
	ErrUpdatingWallet = errors.New("error updating wallet")
	ErrRegisteringUser = errors.New("error registering user")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrFetchingBalance = errors.New("error fetching balance from wallet")
	ErrDebitingWallet = errors.New("error debiting wallet")
)