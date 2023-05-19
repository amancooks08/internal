package service

import (
	"context"
	"errors"
	"nickPay/wallet/internal/db/mocks"
	"nickPay/wallet/internal/domain"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	service    WalletService
	repository *mocks.Storer
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.repository = new(mocks.Storer)
	suite.service = NewWalletService(suite.repository)
}

func (suite *ServiceTestSuite) TearDownTest() {
	suite.repository.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestWalletService_RegisterUser() {
	t := suite.T()
	type args struct {
		ctx  context.Context
		user domain.User
	}
	type test struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}
	tests := []test{
		{
			name: "Register Valid User",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Name:        "John Doe",
					Email:       "john1@gmail.com",
					PhoneNumber: "8123467890",
					Password:    "12345678",
				},
			},
			wantErr: false,
			prepare: func(args args, mock *mocks.Storer) {
				args.user.Password = HashPassword(args.user.Password)
				mock.On("RegisterUser", args.ctx, args.user).Return(nil).Once()
			},
		},
		{
			name: "Register User with Invalid Email",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Name:        "John Doe",
					Email:       "john1mail.com",
					PhoneNumber: "8123467890",
					Password:    "1234567",
				},
			},
			wantErr: true,
			prepare: func(args args, s *mocks.Storer) {
				args.user.Password = HashPassword(args.user.Password)
				s.On("RegisterUser", args.ctx, args.user).Return(nil).Once()
			},
		},
		{
			name: "Register User with Invalid Phone Number",
			args: args{
				ctx: context.Background(),
				user: domain.User{
					Name:        "John Doe",
					Email:       "john1@mail.com",
					PhoneNumber: "812346789",
					Password:    "$12345",
				},
			},
			wantErr: true,
			prepare: func(args args, s *mocks.Storer) {
				args.user.Password = HashPassword(args.user.Password)
				s.On("RegisterUser", args.ctx, args.user).Return(nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.user.Password = HashPassword(tt.args.user.Password)
			tt.prepare(tt.args, suite.repository)
			err := suite.service.RegisterUser(tt.args.ctx, tt.args.user)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func (suite *ServiceTestSuite) TestWalletService_LoginUser() {
	t := suite.T()
	type args struct {
		ctx          context.Context
		loginRequest domain.LoginUserRequest
	}
	type test struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}
	tests := []test{
		{
			name: "Login Valid User",
			args: args{
				ctx: context.Background(),
				loginRequest: domain.LoginUserRequest{
					Email:    "john1@mail.com",
					Password: "12345678",
				},
			},
			wantErr: false,
			prepare: func(args args, mock *mocks.Storer) {
				mock.On("LoginUser", args.ctx, args.loginRequest.Email).Return(domain.LoginDbResponse{}, nil).Once()
			},
		},
		{
			name: "Login User with Invalid Email",
			args: args{
				ctx: context.Background(),
				loginRequest: domain.LoginUserRequest{
					Email:    "john1mail.com",
					Password: "12345678",
				},
			},
			wantErr: true,
			prepare: func(args args, s *mocks.Storer) {
				s.On("LoginUser", args.ctx, args.loginRequest.Email).Return(domain.LoginDbResponse{}, errors.New("mocked error")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, suite.repository)
			_, err := suite.service.LoginUser(tt.args.ctx, tt.args.loginRequest)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func (suite *ServiceTestSuite) TestWalletService_GetWallet() {
	t := suite.T()
	type args struct {
		ctx    context.Context
		userID int
	}
	type test struct {
		name    string
		args    args
		wantErr bool
		prepare func(args, *mocks.Storer)
	}
	tests := []test{
		{
			name: "Valid Request to Get Wallet",
			args: args{
				ctx:    context.WithValue(context.Background(), "id", 1),
				userID: 1,
			},
			wantErr: false,
			prepare: func(args args, s *mocks.Storer) {

				s.On("GetWallet", args.ctx, args.userID).Return(domain.Wallet{}, nil).Once()
			},
		},
		{
			name: "Invalid Request to Get Wallet",
			args: args{
				ctx:    context.WithValue(context.Background(), "id", 0),
				userID: 0,
			},
			wantErr: true,
			prepare: func(args args, s *mocks.Storer) {
				s.On("GetWallet", args.ctx, args.userID).Return(domain.Wallet{}, errors.New("mocked error")).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.args, suite.repository)
			_, err := suite.service.GetWallet(tt.args.ctx, tt.args.userID)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
