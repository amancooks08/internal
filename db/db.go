package db

import (
	"context"
	"nickPay/wallet/internal/domain"
)

type Storer interface {
	RegisterUser(context.Context, domain.User) error
	LoginUser(context.Context, string) (domain.LoginDbResponse, error)
	CreateWallet(context.Context, int64) error
	GetWallet(context.Context, int) (domain.Wallet, error)
	CreditWallet(context.Context, int, float64) error
}
