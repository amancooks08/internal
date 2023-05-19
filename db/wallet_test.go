package db

import (
	"context"
	"errors"
	"nickPay/wallet/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func (suite *StoreTestSuite) Test_pgStore_GetWallet() {
	t := suite.T()
	type args struct {
		ctx    context.Context
		userId int
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
				userId: 1,
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
				userId: 2,
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

			rows := sqlxmock.NewRows([]string{"id", "user_id", "balance", "creation_date", "last_updated", "status"}).AddRow(1, 1, 1000.0, time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02 15:04:05"), "active")
			suite.mock.ExpectQuery(`SELECT \* FROM wallet`).WithArgs(tt.args.userId).WillReturnError(err).WillReturnRows(rows)

			wallet, err := suite.repo.GetWallet(tt.args.ctx, tt.args.userId)
			if (err != nil) == tt.wantErr {
				require.Equal(t, tt.want, wallet)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
