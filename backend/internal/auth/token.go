package auth

import (
	"api/internal/config"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// GenerateToken creates a token with the given user permissions
func GenerateToken(userID int) (string, error) {
	permissions := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 6).Unix(),
		"userID":     userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// GetToken extracts the JWT token from request header and return it
func GetToken(c *fiber.Ctx) (JWTToken string) {
	AuthorizationValue := c.GetReqHeaders()["Authorization"]
	JWTToken = strings.Split(AuthorizationValue, " ")[1]
	return
}
