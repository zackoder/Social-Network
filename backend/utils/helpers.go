package utils

import (
	"strings"

	"net/http"
	"regexp"

	"github.com/gofrs/uuid"
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

func IsValidEmail(email *string) bool {
	*email = strings.ToLower(*email)
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(*email)
}

func CheckPasswordHash(password, hash *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}

func GenerateSessionID() (string, error) {
	sessionID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return sessionID.String(), nil
}

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
