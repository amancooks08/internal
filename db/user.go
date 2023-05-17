package db

import (
	"context"
	"nickPay/wallet/internal/domain"

	logger "github.com/sirupsen/logrus"
)

type User struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (s *pgStore) RegisterUser(ctx context.Context, user domain.User) error {
	err := s.db.QueryRowxContext(ctx, `INSERT INTO "user" (name, email, number, password) VALUES ($1, $2, $3, $4) RETURNING id`, user.Name, user.Email, user.PhoneNumber, user.Password).Scan(&user.ID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot insert user")
		return err
	}
	return nil
}
