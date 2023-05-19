package errors

import (
	"errors"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrNoWallet = errors.New("no wallet found")
	ErrFetchingWallet = errors.New("error fetching wallet")
)