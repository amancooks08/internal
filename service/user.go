package service

import (
	"nickPay/wallet/internal/domain"
	errors "nickPay/wallet/internal/errors"
	"regexp"

	bcrypt "golang.org/x/crypto/bcrypt"
)

type User struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Validate(user domain.User) (error) {
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

func ValidateEmail(email string) (bool) {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email) 
}

func ValidatePhoneNumber(phoneNumber string) (bool) {
	re := regexp.MustCompile(`^[0-9]{10}$`)
	return re.MatchString(phoneNumber) 
}
