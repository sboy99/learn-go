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
