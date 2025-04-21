package auth

import (
	"os"
	"time"

	"github.com/dionisioedu/StickerVerse/internal/user"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(u *user.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "changeme"
	}

	claims := jwt.MapClaims{
		"sub":   u.ID,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
