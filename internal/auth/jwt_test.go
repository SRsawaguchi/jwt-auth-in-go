package auth

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"strings"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	username := "Zola Russel"
	secret := "9FdZ&T*x4gu9"
	duration := time.Hour * 24
	expiresAt := time.Now().Add(duration).Unix()

	jwtString, err := GenerateToken(username, secret, expiresAt)
	if err != nil {
		t.Errorf(err.Error())
	}

	token := strings.Split(jwtString, ".")
	payload, _ := base64.StdEncoding.DecodeString(token[1])
	claim := NamedClaims{}
	if err := json.Unmarshal(payload, &claim); err != nil {
		t.Error(err.Error())
	}

	if claim.Name != username {
		t.Errorf("Invalid name: expected '%v' but got '%v'", username, claim.Name)
	}

	if math.Abs(float64(claim.ExpiresAt-expiresAt)) != 0 {
		t.Errorf("Invalid ExpiresAt: expected '%v' but got '%v' (%v)", expiresAt, claim.ExpiresAt, expiresAt-claim.ExpiresAt)
	}
}

func TestParseToken(t *testing.T) {
	expectedName := "Rahsaan Schiller"
	secret := "4MafgDmBsSa55&5Sf2Rt8tbnG6CbW@q2"
	// generated by https://jwt.io/
	jwtString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlJhaHNhYW4gU2NoaWxsZXIiLCJpYXQiOjE1MTYyMzkwMjJ9.H_bSqiBeAOu3A_X4XmTBJW-NyPN2tI3SVjHK3DLgv1I"

	name, err := ParseToken(jwtString, secret)
	if err != nil {
		t.Error(err.Error())
	}
	if name != expectedName {
		t.Errorf("Invalid name: expected '%v' but got '%v'", expectedName, name)
	}
}
