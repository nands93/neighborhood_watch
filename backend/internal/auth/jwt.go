package auth

import (
	"fmt"
	"os"
	"time"
	"vizinhanca/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

type AppClaims struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *model.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := AppClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "neighborhood-watch",
			Subject:   fmt.Sprintf("%d", user.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
