package auth

import (
	"api/internal/config"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a token with the given user permissions
func GenerateToken(userID int) (string, error) {
	permissions := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 6).Unix(),
		"userID":     userID,
	}

	JWTToken := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return JWTToken.SignedString([]byte(config.SecretKey))
}

// GetUserID extracts the user ID from JWT token an return it
func GetUserIDFromToken(authorization string) (userID int, err error) {
	JWTTokenString := getTokenFromHeaders(authorization)
	JWTToken, err := jwt.Parse(JWTTokenString, returnSecrectKey)
	if err != nil {
		return
	}

	if permissions, ok := JWTToken.Claims.(jwt.MapClaims); ok && JWTToken.Valid {
		var ID int64

		ID, err = strconv.ParseInt(fmt.Sprintf("%.f", permissions["userID"]), 10, 64)
		if err != nil {
			return
		}

		userID = int(ID)

		return
	}

	err = errors.New("invalid token")
	return
}

func getTokenFromHeaders(authorization string) (JWTToken string) {
	JWTToken = strings.Split(authorization, " ")[1]
	return
}

func returnSecrectKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method. received method: %v", token.Header["alg"])
	}

	return []byte(config.SecretKey), nil
}
