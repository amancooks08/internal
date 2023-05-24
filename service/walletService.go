package service

import (
	"context"
	"nickPay/wallet/internal/db"
	"nickPay/wallet/internal/domain"
	errors "nickPay/wallet/internal/errors"

	logger "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type WalletService interface {
	RegisterUser(context.Context, domain.User) error
	LoginUser(context.Context, domain.LoginUserRequest) (string, error)
	GetWallet(context.Context, int64) (domain.Wallet, error)
	CreditWallet(context.Context, int64, float64) error
	DebitWallet(context.Context, int64, float64) error
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
		user.Password = HashPassword(user.Password)
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
		return "", errors.ErrInvalidPassword
	}

	if err != nil {
		logger.WithField("err", err).Error("Error while logging in user")
		return "", err
	}
	token, err = GenerateToken(loginResponse)
	if err != nil {
		logger.WithField("err", err.Error()).Error(errors.ErrGenJWTToken.Error())
		return "", errors.ErrGenJWTToken
	}
	return token, nil
}

func (w *walletService) GetWallet(ctx context.Context, userID int64) (wallet domain.Wallet, err error) {
	wallet, err = w.store.GetWallet(ctx, userID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(errors.ErrFetchingWallet.Error())
		return domain.Wallet{}, errors.ErrFetchingWallet
	}
	return wallet, nil
}

func (w *walletService) CreditWallet(ctx context.Context, userID int64, amount float64) (err error) {
	err = w.store.CreditWallet(ctx, userID, amount)
	if err != nil {
		logger.WithField("err", err.Error()).Error(errors.ErrCreditingWallet.Error())
		return errors.ErrCreditingWallet
	}
	return nil
}

func (w *walletService) DebitWallet(ctx context.Context, userID int64, amount float64) (err error) {
	wallet, err := w.store.GetWallet(ctx, userID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(errors.ErrFetchingBalance.Error())
		return errors.ErrFetchingBalance
	}
	if wallet.Balance < amount {
		logger.WithField("err", err.Error()).Error(errors.ErrInsufficientBalance.Error())
		return errors.ErrInsufficientBalance
	}
	err = w.store.DebitWallet(ctx, userID, amount)
	if err != nil {
		logger.WithField("err", err.Error()).Error(errors.ErrDebitingWallet.Error())
		return errors.ErrDebitingWallet
	}
	return nil
}
