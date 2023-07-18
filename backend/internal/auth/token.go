package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken creates a token with the given user permissions
func GenerateToken(userID int) (string, error) {
	permissions := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 6).Unix(),
		"userID":     userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString("Secret")
}
