package service

import (
	"context"
	"nickPay/wallet/internal/db"
	"nickPay/wallet/internal/domain"
	"errors"
	errorrs "nickPay/wallet/internal/errors"
	logger "github.com/sirupsen/logrus" 
	"golang.org/x/crypto/bcrypt"
)

type WalletService interface {
	RegisterUser(context.Context, domain.User) error
	LoginUser(context.Context, domain.LoginUserRequest) (string, error)
}

type walletService struct {
	store db.Storer
}

func NewWalletService(storer db.Storer) WalletService {
	return &walletService{store: storer}
}

func (w *walletService) RegisterUser(ctx context.Context, user domain.User) (err error) {
	user = domain.User{
		ID:          0,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}
	err = Validate(user)
	if err == nil {
		err = w.store.RegisterUser(ctx, user)
		if err != nil {
			return
		}
		return
	}
	return
}

func (w *walletService) LoginUser(ctx context.Context, loginRequest domain.LoginUserRequest) (token string, err error) {
	loginResponse, err := w.store.LoginUser(ctx, loginRequest.Email)
	if bcrypt.CompareHashAndPassword([]byte(loginResponse.Password), []byte(loginRequest.Password)) != nil {
		return "", errorrs.ErrInvalidPassword
	}

	if err != nil {
		logger.WithField("err", err).Error("Error while logging in user")
		return "", err
	}
	token, err = GenerateToken(loginResponse)
	if err != nil {
		logger.WithField("err", err.Error()).Error("error generating jwt token for given userId")
		return "", errors.New("error generating jwt token for given userId")
	}
	return token, nil
}
