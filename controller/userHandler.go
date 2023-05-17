package controller

import (
	"encoding/json"
	"net/http"
	"nickPay/wallet/internal/domain"
	service "nickPay/wallet/internal/service"
)

func RegisterUser(NikPay service.WalletService) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var user domain.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		err = NikPay.RegisterUser(r.Context(), user)

		if err != nil {
			message := service.RegisterUserResponse{
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
		message := service.RegisterUserResponse{
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
