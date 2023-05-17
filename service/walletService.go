package service

import (
	"context"
	"nickPay/wallet/internal/db"
	"nickPay/wallet/internal/domain"
)

type WalletService interface {
	RegisterUser(context.Context, domain.User) error
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
