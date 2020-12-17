package auth

import "golang.org/x/crypto/bcrypt"

// GeneratePasswordHash hashes given password
func GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// IsValidPassword compares raw password with it's hashed values
func IsValidPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
