package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (hashed []byte, err error) {
	hashed, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
