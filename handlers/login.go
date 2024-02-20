package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func LoginHandler(secretKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := GenerateJWT(secretKey)
		if err != nil {
			fmt.Println("Error generating token string")
		} else {
			fmt.Println("Generated token string: ", token)
		}
		// Return cookie
		cookie := &http.Cookie{
			Name:    "session_token",
			Value:   token, // replace with the actual session token
			Expires: time.Now().Add(24 * time.Hour),

			MaxAge:   24 * 60 * 60,
			HttpOnly: true,
			Secure:   false,
			Path:     "/"}

		// Set the cookie in the response header
		http.SetCookie(w, cookie)
		w.Write([]byte("Login successful"))
	}
}
func GenerateJWT(SecretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "John Doe"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
