package iofunctions

import (
	"fmt"
	"strings"
)

// prompt user for lang input, returns all lowercase lang
func promptLang() string {
	var lang string
	fmt.Println("Enter language name: ")
	fmt.Scanf("%s", &lang)

	return strings.ToLower(lang)
}

// capitalize 1st letter of string
func firstLetterCap(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
