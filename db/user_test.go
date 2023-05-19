package db

import (
	"context"
	"errors"
	"nickPay/wallet/internal/domain"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

type StoreTestSuite struct {
	suite.Suite
	repo Storer
	mock sqlxmock.Sqlmock
}

func TestUserStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}

func (suite *StoreTestSuite) SetupTest() {
	var err error
	var db *sqlx.DB
	db, suite.mock, err = sqlxmock.Newx()
	suite.Require().NoError(err)
	suite.repo = NewPgStore(db)
}

func (suite *StoreTestSuite) TearDownTest() {
	suite.mock.ExpectClose()
}

func (suite *StoreTestSuite) Test_pgStore_RegisterUser() {
	t := suite.T()
	type args struct {
		ctx  context.Context
		user domain.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Register Valid User",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					ID:          1,
					Name:        "John Doe",
					Email:       "john1@mail.com",
					PhoneNumber: "8123467890",
					Password:    "12345678",
				},
			},
			wantErr: false,
		},
		{
			name: "Register User with Invalid Email",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					ID:          2,
					Name:        "John Dee",
					Email:       "john2mail.com",
					PhoneNumber: "8123467891",
					Password:    "12345678",
				},
			},
			wantErr: true,
		},
		{
			name: "Register User with Invalid Phone Number",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					ID:          3,
					Name:        "John Doe",
					Email:       "john1@maiil.com",
					PhoneNumber: "812346789",
					Password:    "12345678",
				},
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

			suite.mock.ExpectQuery(`INSERT INTO "user"`).WithArgs(tt.args.user.Name, tt.args.user.Email, tt.args.user.PhoneNumber, tt.args.user.Password).WillReturnError(err).WillReturnRows(rows)

			if err := suite.repo.RegisterUser(tt.args.ctx, tt.args.user); (err != nil) == tt.wantErr {
				if tt.wantErr {
					require.EqualError(t, err, "mocked error")
				} else {
					require.NoError(t, err)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func (suite *StoreTestSuite) Test_pgStore_LoginUser() {
	t := suite.T()
	type args struct {
		ctx  context.Context
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.LoginDbResponse
		wantErr bool
	}{
		{
			name: "Login Valid User",
			args: args{
				ctx: context.Background(),
				email: "john1@gmail.com",
			},
			want: domain.LoginDbResponse{
				ID:          1,
				Password:   "12345678",
			},
			wantErr: false,
		},
		{
			name: "Login Invalid User",
			args: args{
				ctx: context.Background(),
				email: "john2@mail.com",
			},
			want: domain.LoginDbResponse{},
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

			rows := sqlxmock.NewRows([]string{"id", "password"}).AddRow(1, "12345678")

			suite.mock.ExpectQuery(`SELECT id, password FROM "user"`).WithArgs(tt.args.email).WillReturnError(err).WillReturnRows(rows)

			got, err := suite.repo.LoginUser(tt.args.ctx, tt.args.email)
			if (err != nil) == tt.wantErr {
				if tt.wantErr {
					require.EqualError(t, err, "mocked error")
				} else {
					require.NoError(t, err)
				}
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, got)
		})
	}
}