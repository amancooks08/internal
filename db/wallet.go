package db

import (
	"context"
	"database/sql"
	"nickPay/wallet/internal/domain"
	"nickPay/wallet/internal/errors"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (s *pgStore) CreateWallet(ctx context.Context, userID int64) (err error){
	rows, err := s.db.QueryContext(ctx, `INSERT INTO "wallet" (user_id, balance, creation_date, last_updated, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`, &userID, 0.0, time.Now().Local().Format("2006-01-02"), time.Now().Local().Format("2006-01-02 15:04:05"), "active")
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cannot insert wallet")
		return err
	}
	defer rows.Close()
	return nil
}


func (s *pgStore) GetWallet(ctx context.Context, userID int64) (wallet domain.Wallet, err error) {
	wallet = domain.Wallet{}
	rows, err := s.db.Query("SELECT * FROM wallet where user_id = $1", &userID)
	if err != nil && err == sql.ErrNoRows {
		logger.WithField("err", err.Error()).Error(errors.ErrNoWallet.Error())
		return wallet, errors.ErrNoWallet
	} else if err != nil {
		logger.WithField("err", err.Error()).Error(errors.ErrFetchingWallet.Error())
		return wallet, errors.ErrFetchingWallet
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&wallet.ID, &wallet.UserID, &wallet.Balance, &wallet.CreationDate, &wallet.LastUpdated, &wallet.Status)
		if err != nil {
			logger.WithField("err", err.Error()).Error(errors.ErrFetchingWallet.Error())
			return wallet, errors.ErrFetchingWallet
		}
	}
	return
}

func (s *pgStore) CreditWallet(ctx context.Context, userID int64, amount float64) (err error) {
	result, err := s.db.Exec(`UPDATE "wallet" SET balance = balance + $1, last_updated = $2 WHERE user_id = $3`, &amount, time.Now().Local().Format("2006-01-02 15:04:05"), &userID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(errors.ErrUpdatingWallet.Error())
		return errors.ErrUpdatingWallet
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.WithField("err", err.Error()).Error(errors.ErrUpdatingWallet.Error())
		return errors.ErrUpdatingWallet
	}
	if rowsAffected == 0 && rowsAffected != 1 {
		logger.WithField("err", err.Error()).Error(errors.ErrUpdatingWallet.Error())
		return errors.ErrUpdatingWallet
	}
	return
}


func (s *pgStore) DebitWallet(ctx context.Context, userID int64, amount float64) (err error) {
	result, err := s.db.Exec(`UPDATE "wallet" SET balance = balance - $1, last_updated = $2 WHERE user_id = $3`, &amount, time.Now().Local().Format("2006-01-02 15:04:05"), &userID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(err.Error())
		return errors.ErrUpdatingWallet
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.WithField("err", err.Error()).Error(err.Error())
		return errors.ErrUpdatingWallet
	}
	if rowsAffected == 0 && rowsAffected != 1 {
		logger.WithField("err", err.Error()).Error(err.Error())
		return errors.ErrUpdatingWallet
	}
	return
}