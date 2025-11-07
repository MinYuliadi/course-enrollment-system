package utils

import (
	"course-enrollment-system/config"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JWTClaim struct {
	Username string `json:"username"`
	MyId     int    `json:"myId"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string, id int) (string, error) {
	config.InitEnv()

	if err := godotenv.Load(); err != nil {
		return "", err
	}

	jwtKey := []byte(os.Getenv("JWT_KEY"))

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := JWTClaim{
		Username: username,
		MyId:     id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*JWTClaim, error) {
	config.InitEnv()

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	jwtKey := []byte(os.Getenv("JWT_KEY"))

	claims := &JWTClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
