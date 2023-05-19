package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"nickPay/wallet/internal/domain"
	"nickPay/wallet/internal/service/mocks"
	"nickPay/wallet/server"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type WalletHandlerSuite struct {
	suite.Suite
	service *mocks.WalletService
}

func TestWalletHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(WalletHandlerSuite))
}

func (suite *WalletHandlerSuite) SetupTest() {
	suite.service = &mocks.WalletService{}
}

func (suite *WalletHandlerSuite) TeardownTest() {
	suite.service.AssertExpectations(suite.T())
}

func (suite *WalletHandlerSuite) TestWallet_GetWallet() {
	t := suite.T()
	t.Run("Get Wallet using user-id", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/user/wallet", nil)
		rw := httptest.NewRecorder()
		ctx := req.Context()
		ctx = context.WithValue(ctx, "id", 1)
		req = req.WithContext(ctx)
		expectedResponse := domain.Wallet{
			ID:           1,
			UserID:       1,
			Balance:      1000,
			CreationDate: "2021-09-01",
			LastUpdated:  "2021-09-01",
			Status:       "active",
		}
		exp, err := json.Marshal(expectedResponse)
		if err != nil {
			t.Errorf("Error while marshalling expected response: %v", err)
		}

		// Act
		suite.service.On("GetWallet", ctx, 1).Return(expectedResponse, nil).Once()
		deps := server.Dependencies{
			NikPay: suite.service,
		}

		// Assert
		got := GetWallet(deps.NikPay)
		got.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Equal(t, string(exp), rw.Body.String())
	})

	t.Run("Invalid request to get Wallet", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(http.MethodGet, "/user/wallet", nil)
		rw := httptest.NewRecorder()
		ctx := req.Context()
		ctx = context.WithValue(ctx, "id", 1)
		req = req.WithContext(ctx)
		expectedResponse := domain.Message{
			Message: "invalid request",
		}
		exp, err := json.Marshal(expectedResponse)
		if err != nil {
			t.Errorf("Error while marshalling expected response: %v", err)
		}

		// Act
		suite.service.On("GetWallet", ctx, 1).Return(domain.Wallet{}, errors.New("invalid request")).Once()
		deps := server.Dependencies{
			NikPay: suite.service,
		}

		// Assert
		got := GetWallet(deps.NikPay)
		got.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusBadRequest, rw.Code)
		assert.Equal(t, string(exp), rw.Body.String())
	})
}
