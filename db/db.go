package db

import (
	"context"
	"nickPay/wallet/internal/domain"
)

type Storer interface {
	RegisterUser(context.Context, domain.User) error
	LoginUser(context.Context, string) (domain.LoginDbResponse, error)
	GetWallet(context.Context, int) (domain.Wallet, error)
}
