package controller

import (
	"encoding/json"
	"net/http"
	"nickPay/wallet/internal/domain"
	"nickPay/wallet/internal/service"
)

func GetWallet(NikPay service.WalletService) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		userID := r.Context().Value("id").(int)
		var wallet domain.Wallet
		wallet, err := NikPay.GetWallet(r.Context(), userID)
		if err != nil {
			message := domain.Message{
				Message: err.Error(),
			}
			rw.WriteHeader(http.StatusBadRequest)
			resp, err := json.Marshal(message)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(resp)
			return
		}
		message := domain.GetWalletResponse{
			ID:           wallet.ID,
			Balance:      wallet.Balance,
			CreationDate: wallet.CreationDate,
			LastUpdated:  wallet.LastUpdated,
			Status:       wallet.Status,
		}
		resp, err := json.Marshal(message)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(resp)
	})
}

func CreditWallet(NikPay service.WalletService) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		userID := r.Context().Value("id").(int)
		var credit domain.Credit
		err := json.NewDecoder(r.Body).Decode(&credit)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		err = NikPay.CreditWallet(r.Context(), userID, credit.Amount)
		if err != nil {
			message := domain.Message{
				Message: err.Error(),
			}
			rw.WriteHeader(http.StatusBadRequest)
			resp, err := json.Marshal(message)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(resp)
			return
		}
		message := domain.Message{
			Message: "Wallet credited successfully",
		}
		resp, err := json.Marshal(message)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(resp)
	})
}