package iofunctions

import (
	"fmt"
	"strings"
)

// prompt user for lang input, returns all lowercase lang
func PromptLang() string {
	var lang string
	fmt.Println("Enter language name: ")
	fmt.Scanf("%s", &lang)

	return strings.ToLower(lang)
}

func capFirstLetter(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
