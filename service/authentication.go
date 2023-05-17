package service

import (
	"nickPay/wallet/internal/domain"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret@987")

func GenerateToken(loginResponse domain.LoginDbResponse) (string, error) {
	tokenExpirationTime := time.Now().Add(time.Minute * 30)
	tokenObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": loginResponse.ID,
		"exp":     tokenExpirationTime.Unix(),
	})
	token, err := tokenObject.SignedString(secretKey)
	return token, err
}
