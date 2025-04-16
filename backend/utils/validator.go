package utils

import (
	"fmt"
	"regexp"
)

func ValidatNames(names ...any) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z]{1,}\w+`)
	for i, name := range names {
		str, ok := name.(string)
		if i != 2 || ok {
			fmt.Println(str)
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
