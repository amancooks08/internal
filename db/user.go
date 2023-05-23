package db

import (
	"context"
	"database/sql"
	"nickPay/wallet/internal/domain"
	"nickPay/wallet/internal/errors"

	logger "github.com/sirupsen/logrus"
)

func (s *pgStore) RegisterUser(ctx context.Context, user domain.User) error {
	rows, err := s.db.QueryContext(ctx, `INSERT INTO "user" (name, email, number, password) VALUES ($1, $2, $3, $4)`, user.Name, user.Email, user.PhoneNumber, user.Password)
	if err != nil {
		logger.WithField("err", err).Error("Error while registering user")
		return errors.ErrRegisteringUser
	}
	defer rows.Close()
	return nil
}

func (s *pgStore) LoginUser(ctx context.Context, requestEmail string) (loginResponse domain.LoginDbResponse, err error) {
	loginResponse = domain.LoginDbResponse{}
	rows, err := s.db.QueryContext(ctx, `SELECT id, password FROM "user" WHERE email = $1`, requestEmail)
	if err == sql.ErrNoRows {
		logger.WithField("err", err).Error("user not found")
		return loginResponse, err
	}
	if err != nil {
		logger.WithField("err", err).Error("Error while logging in user")
		return loginResponse, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&loginResponse.ID, &loginResponse.Password)
		if err != nil {	
			logger.WithField("err", err).Error("Error while scanning login response")
			return loginResponse, err
		}
	}
	return loginResponse, nil
}
