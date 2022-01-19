package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashSalt(password string) (string, error) {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashs(hashedPwd, password string) bool {
	byteHash := []byte(hashedPwd)
	plainPwd := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
