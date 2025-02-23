package utils

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain text password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error tidak diketahui : %v", err.Error())

		return "", errors.New("error tidak diketahui")
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		if err  == bcrypt.ErrMismatchedHashAndPassword {
			err = errors.New("password tidak cocok")
		} else {
			log.Printf("Error tidak diketahui : %v", err.Error())

			err = errors.New("error tidak diketahui")
		}
	}
	
	return err
}