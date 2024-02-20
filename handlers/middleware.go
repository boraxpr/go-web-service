package handlers

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"net/http"
)

// important key is to fetch from frontend by using credentials :"include"
func VerifyToken(requestToken string, SecretKey string) bool {
	if requestToken == "" {
		return false
	}
	strings.Split(requestToken, "=")
	requestToken = requestToken[strings.Index(requestToken, "=")+1:]

	fmt.Println("Verifying token: ", requestToken)
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
			requestToken := r.Header.Get("Cookie")

			if !VerifyToken(requestToken, SecretKey) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
}
