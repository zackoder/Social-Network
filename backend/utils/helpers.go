package utils

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Hashpass(password string) string {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return ""
	}
	return string(hashedPasswd)
}

func CheckExtension(FileExtension string) bool {
	extensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp", ".ico", ".heif", ".apng"}
	for _, extension := range extensions {
		if strings.Contains(FileExtension, extension) {
			return true
		}
	}
	return false
}
