package utils

import "golang.org/x/crypto/bcrypt"

func Hashpass(password string) string {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return ""
	}
	return string(hashedPasswd)
}
