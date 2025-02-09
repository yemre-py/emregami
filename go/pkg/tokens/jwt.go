package tokens

import (
	"emregami/pkg/configs"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func generateAccessToken(id string) (string, error) {
	now := time.Now()
	token := &claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(configs.NewTokenConfig().AccessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "emregami",
			Subject:   id,
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, token).
		SignedString([]byte(configs.NewTokenConfig().AccessTokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAccessToken(token string) (*claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
		}

		return []byte(configs.NewTokenConfig().AccessTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*claims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func generateRefreshToken(id string) (string, error) {
	now := time.Now()
	token := &claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "emregami",
			Subject:   id,
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, token).
		SignedString([]byte(configs.NewTokenConfig().RefreshTokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateRefreshToken(token string) (*claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
		}

		return []byte(configs.NewTokenConfig().RefreshTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*claims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func GenerateTokens(id string) (string, string, error) {
	accessToken, err := generateAccessToken(id)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
