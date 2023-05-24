package db

import (
	"context"
	"nickPay/wallet/internal/domain"
)

type Storer interface {
	RegisterUser(context.Context, domain.User) error
	LoginUser(context.Context, string) (domain.LoginDbResponse, error)
	CreateWallet(context.Context, int64) error
	GetWallet(context.Context, int64) (domain.Wallet, error)
	CreditWallet(context.Context, int64, float64) error
	DebitWallet(context.Context, int64, float64) error
}
