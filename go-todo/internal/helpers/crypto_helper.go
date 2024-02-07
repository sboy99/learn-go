package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashStr(str string) (string, error) {
	buff, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

func CompareHash(str string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	if err != nil {
		return false
	}
	return true
}
