package tokens

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type claims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id string) (string, error) {
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

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string) (*Claims, error) {
	claims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
}
