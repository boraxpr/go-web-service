package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler authenticates a user and returns a session token.
// @Summary User login
// @Description Authenticates a user and returns a session token.
// @Accept  json
// @Produce  json
// @Param   secretKey     query    string     true  "Secret Key"
// @Success 200 {string} string "Login successful"
// @Failure 400 {string} string "Error generating token string"
// @Router /login [post]
func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		secretKey := os.Getenv("SECRET_KEY")
		// TODO: receive username and password from body
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
			SameSite: http.SameSiteStrictMode,
			Path:     "/"}

		// Set the cookie in the response header
		http.SetCookie(w, cookie)
		w.Write([]byte("Login successful"))
	}
}

// TODO: Pass needed parameters in the request
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
