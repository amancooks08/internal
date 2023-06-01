package service

import (
	"crypto/sha256"
	"encoding/hex"
	"nickPay/wallet/internal/domain"
	errors "nickPay/wallet/internal/errors"
	"regexp"
)

func HashPassword(password string) string {
	
	hash := sha256.New()

	hash.Write([]byte(password))

	hashBytes := hash.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func Validate(user domain.User) error {
	if !ValidateEmail(user.Email) {
		return errors.ErrInvalidEmail
	}
	if !ValidatePhoneNumber(user.PhoneNumber) {
		return errors.ErrInvalidPhoneNumber
	}
	if user.Name == "" {
		return errors.ErrInvalidName
	}
	if user.Password == "" {
		return errors.ErrInvalidPassword
	}
	return nil
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidatePhoneNumber(phoneNumber string) bool {
	re := regexp.MustCompile(`^[0-9]{10}$`)
	return re.MatchString(phoneNumber)
}
