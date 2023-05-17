package db

import (
	"context"
	"database/sql"
	"nickPay/wallet/internal/domain"

	logger "github.com/sirupsen/logrus"
)


func (s *pgStore) RegisterUser(ctx context.Context, user domain.User) error {
	err := s.db.QueryRowxContext(ctx, `INSERT INTO "user" (name, email, number, password) VALUES ($1, $2, $3, $4) RETURNING id`, user.Name, user.Email, user.PhoneNumber, user.Password).Scan(&user.ID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot insert user")
		return err
	}
	return nil
}

func (s *pgStore) LoginUser(ctx context.Context, requestEmail string) (loginResponse domain.LoginDbResponse, err error) {
	loginResponse = domain.LoginDbResponse{}
	err = s.db.QueryRowContext(ctx, `SELECT id, password FROM "user" WHERE email = $1`, requestEmail).Scan(&loginResponse.ID, &loginResponse.Password)
	if err == sql.ErrNoRows {
		logger.WithField("err", err).Error("User not found")
		return loginResponse, err
	}
	if err != nil {
		logger.WithField("err", err).Error("Error while logging in user")
		return loginResponse, err
	}
	return loginResponse, nil
}