package handlers

import (
	"fmt"
	"os"
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

func SwaggerAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("SECRET_KEY")
		user, token, ok := r.BasicAuth()
		if !ok || user != "admin" || !VerifyToken(token, secretKey) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your token"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Logging request:", r.URL.Path)
			next.ServeHTTP(w, r)
		})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			requestToken := r.Header.Get("Cookie")
			SecretKey := os.Getenv("SECRET_KEY")

			if !VerifyToken(requestToken, SecretKey) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
}
