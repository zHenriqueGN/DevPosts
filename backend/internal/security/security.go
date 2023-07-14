package security

import "golang.org/x/crypto/bcrypt"

// HashPassoword receives a password and add a hash to it
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ComparePasswordWithHash compare a hash and a passoword and return if them are equal
func ComparePasswordWithHash(passwordWithHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(password))
}
