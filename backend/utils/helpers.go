package utils

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Hashpass(password string) string {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
	return string(hashedPasswd)
}

func RemoveIMG(filepath string) {
	if filepath == "" {
		return
	}
	if err := os.Remove(filepath); err != nil {
		fmt.Println("removing error:", err)
	}
}

func CheckExtension(fileContent string) bool {
	ImageTypes := []string{
		"\xFF\xD8\xFF",      // jpg/jpeg
		"\x89PNG\r\n\x1A\n", // png
		"GIF87a",            // gif (variant 1)
		"GIF89a",            // gif (variant 2)
		"BM",                // bmp
		"II*\x00",           // tiff (little endian)
		"MM\x00*",           // tiff (big endian)
		"RIFF",              // webp (check for "WEBP" at offset 8)
		"\x00\x00\x01\x00",  // ico
		"ftypheic",          // heif/heic (starts after offset 4)
		"\x89PNG\r\n\x1A\n", // apng (same as png, distinguish by 'acTL' chunk)
	}
	for _, ImageType := range ImageTypes {
		// log.Println(fileContent[:len(ImageType)] == ImageType)
		if len(fileContent) > len(ImageType)-1 && fileContent[:len(ImageType)] == ImageType {
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

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func GenerateSessionID() (string, error) {
	sessionID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return sessionID.String(), nil
}

func HashPassword(password *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	*password = string(bytes)
	return err
}

func CheckPasswordHash(password, hash *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}

func CheckName(name string) bool {
	re, err := regexp.Compile(`^[a-zA-Z\s]{4,30}$`)
	if err != nil {
		return false
	}
	return re.MatchString(name)
}

// func CheckAge(age uint8) bool {
// 	if age < 13 || age > 120 {
// 		return false
// 	}
// 	return true
// }

func CheckGender(gender string) bool {
	gender = strings.ToLower(gender)
	genders := []string{"male", "female"}
	return slices.Contains(genders, gender)
}

func CheckNickName(nickname string) bool {
	re, err := regexp.Compile(`^[a-zA-Z1-9-_]{4,30}$`)
	if err != nil {
		return false
	}
	return re.MatchString(nickname)
}

