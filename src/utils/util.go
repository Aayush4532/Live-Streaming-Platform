package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email, role string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET not set")
	}

	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(2 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GenerateStreamKey () (string, error){
	buff := make([] byte, 5);
	_, err := rand.Read(buff);
	if err != nil {
		return "", err;
	}
	return hex.EncodeToString(buff), nil;
}