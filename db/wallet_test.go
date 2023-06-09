package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"nickPay/wallet/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func (suite *StoreTestSuite) Test_pgStore_CreateWallet() {
	t := suite.T()
	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Create Valid Wallet",
			args: args{
				userID: 1,
			},
			wantErr: false,
		},
		{
			name: "Create Invalid Wallet",
			args: args{
				userID: 2,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.wantErr {
				err = errors.New("mocked error")
			} else {
				err = nil
			}

			rows := sqlxmock.NewRows([]string{"id"}).AddRow(1)
			suite.mock.ExpectQuery(`INSERT INTO "wallet"`).
				WithArgs(tt.args.userID, 0.0, time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02 15:04:05"), "active").
				WillReturnRows(rows).
				WillReturnError(err)

			err = suite.repo.CreateWallet(context.Background(), tt.args.userID)
			require.Equal(t, tt.wantErr, err != nil)
		})
	}

}

func (suite *StoreTestSuite) Test_pgStore_GetWallet() {
	t := suite.T()
	type args struct {
		ctx    context.Context
		userID int64
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Wallet
		wantErr bool
	}{
		{
			name: "Get Valid Wallet",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			want: domain.Wallet{
				ID:           1,
				UserID:       1,
				Balance:      1000.0,
				CreationDate: time.Now().Format("2006-01-02"),
				LastUpdated:  time.Now().Format("2006-01-02 15:04:05"),
				Status:       "active",
			},
			wantErr: false,
		},
		{
			name: "wallet not found",
			args: args{
				ctx:    context.Background(),
				userID: 2,
			},
			want:    domain.Wallet{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.wantErr {
				err = errors.New("mocked error")
			} else {
				err = nil
			}

			rows := sqlxmock.NewRows([]string{"id", "user_id", "balance", "creation_date", "last_updated", "status"})
			rows = rows.AddRow(tt.want.ID, tt.want.UserID, tt.want.Balance, tt.want.CreationDate, tt.want.LastUpdated, tt.want.Status)

			suite.mock.ExpectQuery(`SELECT \* FROM wallet`).WithArgs(tt.args.userID).
			WillReturnRows(rows).
			WillReturnError(err)

			wallet, err := suite.repo.GetWallet(tt.args.ctx, tt.args.userID)
			if (err != nil) == tt.wantErr  && err == sql.ErrNoRows{
				require.Equal(t, tt.want, wallet)
			} else if (err != nil) == tt.wantErr {
				require.Equal(t, tt.want, wallet)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func (suite *StoreTestSuite) Test_pgStore_CreditWallet() {
	t := suite.T()
	type args struct {
		ctx    context.Context
		userID int64
		amount float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Credit Valid Wallet",
			args: args{
				ctx:    context.Background(),
				userID: 1,
				amount: 1000.0,
			},
			wantErr: false,
		},
		{
			name: "Credit Invalid Wallet",
			args: args{
				ctx:    context.Background(),
				userID: 2,
				amount: 1000.0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.wantErr {
				err = errors.New("mocked error")
			} else {
				err = nil
			}

			result := sqlxmock.NewResult(0, 1)
			suite.mock.ExpectExec(`UPDATE "wallet"`).WithArgs(tt.args.amount, time.Now().Local().Format("2006-01-02 15:04:05"), tt.args.userID).WillReturnError(err).WillReturnResult(result)
			fmt.Println(err)
			err = suite.repo.CreditWallet(tt.args.ctx, tt.args.userID, tt.args.amount)
			if (err != nil) == tt.wantErr {
				require.Equal(t, err, nil)
			} else {
				require.NoError(t, err)
			}
		})
	}
}


func (suite *StoreTestSuite) Test_pgStore_DebitWallet() {
	t := suite.T()
	type args struct {
		ctx    context.Context
		userID int64
		amount float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// {
		// 	name: "Debit Valid Wallet",
		// 	args: args{
		// 		ctx:    context.Background(),
		// 		userID: 1,
		// 		amount: 1000.0,
		// 	},
		// 	wantErr: false,
		// },
		{
			name: "Debit Invalid Wallet",
			args: args{
				ctx:    context.Background(),
				userID: -2,
				amount: 1000.0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			if tt.wantErr {
				err = errors.New("mocked error")
				fmt.Println(err)
			} else {
				err = nil
			}
			result := sqlxmock.NewResult(0, 1)
			suite.mock.ExpectExec(`UPDATE "wallet"`).WithArgs(tt.args.amount, time.Now().Local().Format("2006-01-02 15:04:05"), tt.args.userID).WillReturnError(err).WillReturnResult(result)
			err = suite.repo.DebitWallet(tt.args.ctx, tt.args.userID, tt.args.amount)
			fmt.Println(err)
			if (err != nil) == tt.wantErr {
				require.Equal(t, err, nil)
			} else {
				require.NoError(t, err)
			}
		})
	}
}