package db

import (
	"context"
	"nickPay/wallet/internal/domain"
)

type Storer interface {
	RegisterUser(context.Context, domain.User) error
}
