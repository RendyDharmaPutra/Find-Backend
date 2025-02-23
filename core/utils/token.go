package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateRandomToken() (string, error) {
	bytes := make([]byte, 32)
	
	_, err := rand.Read(bytes)
	if err != nil {
		log.Printf("Gagal membuat token, error tidak diketahui  : %v", err.Error())

		return "", errors.New("error tidak diketahui")
	}

	return hex.EncodeToString(bytes), nil
}

func GenerateJWT(id string) (string, error) {
	key := os.Getenv("SECRET_KEY")
	
	claims := jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result,err := token.SignedString([]byte(key))
	if err != nil {
		log.Printf("Gagal membuat token, error tidak diketahui  : %v", err.Error())

		result = ""
		err = errors.New("error tidak diketahui")
	}


	return result, err
}