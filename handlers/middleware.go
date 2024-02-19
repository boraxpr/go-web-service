package handlers

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"

	"net/http"
	"strings"
)

func VerifyToken(requestToken string, SecretKey string) bool {
	splitToken := strings.Split(requestToken, "Bearer ")
	if len(splitToken) != 2 {
		fmt.Println("malformed token")
		return false
	}
	requestToken = splitToken[1]

	token, err := jwt.Parse(requestToken,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

	if err == nil && token.Valid {
		fmt.Println("valid token")
		return true
	} else {
		fmt.Println("invalid token")
		return false
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Logging request:", r.URL.Path)
			next.ServeHTTP(w, r)
		})
}

func AuthMiddleware(next http.Handler, SecretKey string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			requestToken := r.Header.Get("Authorization")
			if !VerifyToken(requestToken, SecretKey) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
}
