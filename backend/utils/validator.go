package utils

import (
	"regexp"
	"unicode"
)

func ValidatNames(names ...any) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z]{1,}\w+`)
	for i, name := range names {
		str, ok := name.(string)
		if i != 2 || ok {
			if len(str) > 10 {
				return false
			}
			if !pattern.MatchString(str) {
				return false
			}
		}
	}
	return true
}

func ValidEmail(email string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return len(email) < 50 && pattern.MatchString(email)
}

func ValidatePassWord(s string) bool {
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, ch := range s {
		if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		} else if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}
