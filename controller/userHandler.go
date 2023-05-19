package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nickPay/wallet/internal/domain"
	service "nickPay/wallet/internal/service"
)

func RegisterUser(NikPay service.WalletService) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var user domain.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		err = NikPay.RegisterUser(r.Context(), user)

		if err != nil {
			message := domain.RegisterUserResponse{
				Message: err.Error(),
			}
			rw.WriteHeader(http.StatusBadRequest)
			resp, err := json.Marshal(message)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			rw.WriteHeader(http.StatusBadRequest)
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(resp)
			return
		}
		message := domain.RegisterUserResponse{
			Message: "User Registered Successfully",
		}
		resp, err := json.Marshal(message)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		rw.WriteHeader(http.StatusCreated)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(resp)
	})
}

func LoginUser(NikPay service.WalletService) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var loginRequest domain.LoginUserRequest
		err := json.NewDecoder(r.Body).Decode(&loginRequest)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		token, err := NikPay.LoginUser(r.Context(), loginRequest)

		if err != nil {
			message := domain.LoginUserResponse{
				Message: err.Error(),
				Token:   "",
			}
			rw.WriteHeader(http.StatusBadRequest)
			resp, err := json.Marshal(message)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			rw.WriteHeader(http.StatusBadRequest)
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(resp)
			return
		}
		message := domain.LoginUserResponse{
			Message: "User Logged In Successfully",
			Token:   token,
		}
		resp, err := json.Marshal(message)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(resp)
	})
}
