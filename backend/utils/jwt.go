package utils

import (
	"log"
	"time"

<<<<<<< HEAD
	"github.com/imperionite/reeisto/backend/config"
=======
	"github.com/imperionite/reetis/backend/config"
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, role string) (string, error) {
	secret := config.App.JWTSecret
	if secret == "" {
        log.Fatal("JWT_SECRET must be set")
    }

	claims := JWTClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}