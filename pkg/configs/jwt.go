package configs

import (
	"encoding/base64"
	"os"
)

func JwtConfig() (secretKey string) {
	secret := os.Getenv("JWT_SECRET_KEY")
	secretKey = base64.StdEncoding.EncodeToString([]byte(secret))
	return secretKey
}
