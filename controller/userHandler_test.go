package controller

import (
	"encoding/json"
	"net/http"
	"nickPay/wallet/internal/domain"
	"nickPay/wallet/server"
	"strings"

	"net/http/httptest"
	errors "nickPay/wallet/internal/errors"
	"nickPay/wallet/internal/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	service *mocks.WalletService
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

func (suite *UserHandlerTestSuite) SetupTest() {
	suite.service = &mocks.WalletService{}
}
func (suite *UserHandlerTestSuite) TeardownTest() {
	suite.service.AssertExpectations(suite.T())
}

func (suite *UserHandlerTestSuite) TestRegisterUserHandler() {
	t := suite.T()
	t.Run("Register Valid User", func(t *testing.T) {
		// Arrange
		request := domain.RegisterUserRequest{
			Name:        "John Doe",
			Email:       "john.doe@gmail.com",
			PhoneNumber: "9993679833",
			Password:    "12345678",
		}
		reqBody, err := json.Marshal(request)
		assert.NoError(t, err)

		jsonRequest := string(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		expectedResponse := domain.RegisterUserResponse{
			Message: "User Registered Successfully",
		}
		exp, err := json.Marshal(expectedResponse)
		if err != nil {
			t.Errorf("Error while marshalling expected response: %v", err)
		}

		user := domain.User{
			Name:        "John Doe",
			Email:       "john.doe@gmail.com",
			PhoneNumber: "9993679833",
			Password:    "12345678",
		}

		// Act
		suite.service.On("RegisterUser", req.Context(), user).Return(nil).Once()
		deps := server.Dependencies{
			NikPay: suite.service,
		}
		// Assert
		got := RegisterUser(deps.NikPay)
		got.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, string(exp), res.Body.String())
	})
	t.Run("expect error 400 when user had invalid Email", func(t *testing.T) {
		// Arrange
		request := domain.RegisterUserRequest{
			Name:        "John Doe",
			Email:       "john1mail.com",
			PhoneNumber: "9993679833",
			Password:    "12345678",
		}
		reqBody, err := json.Marshal(request)
		assert.NoError(t, err)

		jsonRequest := string(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// TODO - make error object - a more client-centric, enhancing readabililty

		expectedResponse := domain.RegisterUserResponse{
			Message: "invalid email",
		}
		exp, err := json.Marshal(expectedResponse)
		if err != nil {
			t.Errorf("Error while marshalling expected response: %v", err)
		}

		user := domain.User{
			Name:        "John Doe",
			Email:       "john1mail.com",
			PhoneNumber: "8123467890",
			Password:    "12345678",
		}

		// Act
		suite.service.On("RegisterUser", req.Context(), user).Return(errors.ErrInvalidEmail).Once()
		deps := server.Dependencies{
			NikPay: suite.service,
		}
		// Assert
		got := RegisterUser(deps.NikPay)
		got.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), res.Body.String())
	})

	t.Run("Register User with Invalid Phone Number", func(t *testing.T) {
		// Arrange
		request := domain.RegisterUserRequest{
			Name:        "John Doe",
			Email:       "john1@mail.com",
			PhoneNumber: "812346789",
			Password:    "12345678",
		}
		reqBody, err := json.Marshal(request)
		assert.NoError(t, err)

		jsonRequest := string(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		expectedResponse := domain.RegisterUserResponse{
			Message: "invalid phone number",
		}
		exp, err := json.Marshal(expectedResponse)
		if err != nil {
			t.Errorf("Error while marshalling expected response: %v", err)
		}

		user := domain.User{
			Name:        "John Doe",
			Email:       "john1@mail.com",
			PhoneNumber: "812346789",
			Password:    "12345678",
		}

		// Act
		suite.service.On("RegisterUser", req.Context(), user).Return(errors.ErrInvalidPhoneNumber).Once()
		deps := server.Dependencies{
			NikPay: suite.service,
		}
		// Assert
		got := RegisterUser(deps.NikPay)
		got.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), res.Body.String())
	})
}

func (suite *UserHandlerTestSuite) TestLoginUserHandler() {
	t := suite.T()
	t.Run("Login Valid User", func(t *testing.T) {
		// Arrange
		request := domain.LoginUserRequest{
			Email:    "john1@gmail.com",
			Password: "12345678",
		}
		reqBody, err := json.Marshal(request)
		assert.NoError(t, err)

		jsonRequest := string(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		expectedResponse := domain.LoginUserResponse{
			Message: "User Logged In Successfully",
			Token:   "token",
		}
		exp, err := json.Marshal(expectedResponse)
		if err != nil {
			t.Errorf("Error while marshalling expected response: %v", err)
		}

		loginRequest := domain.LoginUserRequest{
			Email:    "john1@gmail.com",
			Password: "12345678",
		}

		// Act
		suite.service.On("LoginUser", req.Context(), loginRequest).Return("token", nil).Once()
		deps := server.Dependencies{
			NikPay: suite.service,
		}
		// Assert
		got := LoginUser(deps.NikPay)
		got.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(exp), res.Body.String())
	})

	t.Run("Login User with Invalid Email", func(t *testing.T) {
		//Arrange
		request := domain.LoginUserRequest{
			Email:    "john1mail.com",
			Password: "12345678",
		}
		reqBody, err := json.Marshal(request)
		assert.NoError(t, err)

		jsonRequest := string(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		expectedResponse := domain.LoginUserResponse{
			Message: "invalid email",
		}
		exp, err := json.Marshal(expectedResponse)
		if err != nil {
			t.Errorf("Error while marshalling expected response: %v", err)
		}

		loginRequest := domain.LoginUserRequest{
			Email:    "john1mail.com",
			Password: "12345678",
		}

		// Act
		suite.service.On("LoginUser", req.Context(), loginRequest).Return("", errors.ErrInvalidEmail).Once()
		deps := server.Dependencies{
			NikPay: suite.service,
		}

		// Assert
		got := LoginUser(deps.NikPay)
		got.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), res.Body.String())
	})
}