package controller

import (
	server "nickPay/wallet/server"

	"github.com/gorilla/mux"
)

func InitRouter(deps *server.Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/register", RegisterUser(deps.NikPay)).Methods("POST")
	router.HandleFunc("/login", LoginUser(deps.NikPay)).Methods("POST")
	router.HandleFunc("/wallet", authMiddleware(GetWallet(deps.NikPay))).Methods("GET")
	router.HandleFunc("/wallet/credit", authMiddleware(CreditWallet(deps.NikPay))).Methods("POST")
	return
}
