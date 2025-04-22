package utils

import (
	"net/http"
	"regexp"
	"strings"

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
