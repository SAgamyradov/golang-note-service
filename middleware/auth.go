package middleware

import (
	"encoding/json"
	"net/http"
	"note-service/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

var (
	users = map[string]string{
		"user": "password", // simple hardcoded user
	}
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds map[string]string
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	password, ok := users[creds["username"]]
	if !ok || password != creds["password"] {
		utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Generation JWT
	token, err := GenerateToken(creds["username"])
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Could not generate token")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
