package db

import (
	"context"
	"database/sql"
	"nickPay/wallet/internal/domain"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (s *pgStore) RegisterUser(ctx context.Context, user domain.User) error {
	err := s.db.QueryRowxContext(ctx, `INSERT INTO "user" (name, email, number, password) VALUES ($1, $2, $3, $4) RETURNING id`, user.Name, user.Email, user.PhoneNumber, user.Password).Scan(&user.ID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot insert user")
		return err
	}
	_, err = s.db.Exec(`INSERT INTO "wallet" (user_id, balance, creation_date, last_updated, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`, user.ID, 0.0, time.Now().Format("2006-01-02"), time.Now().Local().Format("2006-01-02 15:04:05"), "active")
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot insert wallet")
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
