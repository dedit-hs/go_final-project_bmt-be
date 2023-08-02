package middlewares

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateTokenAdmin(adminId int, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["adminId"] = adminId
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_ADMIN")))
}

func CreateTokenAnggota(anggotaId int, nama string) (string, error) {
	claims := jwt.MapClaims{}
	claims["anggotaId"] = anggotaId
	claims["nama"] = nama
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
