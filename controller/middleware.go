package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		header := req.Header.Get("Authorization")

		// Check if the header is missing or invalid
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Parse the JWT token from the header
		token, err := jwt.Parse(strings.TrimPrefix(header, "Bearer "), func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			// Set the secret key for the token
			return []byte("secret@987"), nil
		})

		// Check if there was an error parsing the token
		if err != nil {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid and has not expired
		if !token.Valid {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Get the claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(context.Background(), "id", claims["user_id"])
		req = req.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(rw, req)
	})
}