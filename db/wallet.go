package db

import (
	"context"
	"database/sql"
	"nickPay/wallet/internal/domain"
	"nickPay/wallet/internal/errors"

	logger "github.com/sirupsen/logrus"
)

func (s *pgStore) GetWallet(ctx context.Context, userID int) (wallet domain.Wallet, err error) {
	wallet = domain.Wallet{}
	rows, err := s.db.Query("SELECT * FROM wallet where user_id = $1", &userID)
	if err != nil && err == sql.ErrNoRows {
		logger.WithField("err", err.Error()).Error("no wallet found")
		return wallet, errors.ErrNoWallet
	} else if err != nil {
		logger.WithField("err", err.Error()).Error("error fetching venue")
		return wallet, errors.ErrFetchingWallet
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&wallet.ID, &wallet.UserID, &wallet.Balance, &wallet.CreationDate, &wallet.LastUpdated, &wallet.Status)
		if err != nil {
			logger.WithField("err", err.Error()).Error("error fetching wallet")
			return wallet, errors.ErrFetchingWallet
		}
	}
	return
}
