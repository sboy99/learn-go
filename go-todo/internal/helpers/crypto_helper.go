package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// -----------------------------STRUCTS----------------------------------- //

type CryptoHelper struct{}

// -----------------------------PUBLIC_METHODS----------------------------------- //

func (h *CryptoHelper) HashStr(str string) (string, error) {
	buff, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

func (h *CryptoHelper) CompareHash(str string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	if err != nil {
		return false
	}
	return true
}
