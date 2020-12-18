package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// NamedClaims represents payload of JWT token
type NamedClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenerateToken generates JWT token with passed secretV
func GenerateToken(username, secret string, expiresAt int64) (string, error) {
	claim := NamedClaims{
		Name: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses token and return username
func ParseToken(tokenString, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&NamedClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if claim, ok := token.Claims.(*NamedClaims); ok && token.Valid {
		return claim.Name, nil
	}

	return "", err
}
