package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	lang := promptLang()
	fmt.Println("Entered:", lang)

	WF_PATH, DF_PATH := createTemplate(lang)
	fmt.Println(WF_PATH, DF_PATH)
}

// prompt user for lang input, returns all lowercase lang
func promptLang() string {
	var lang string
	fmt.Println("Enter language name: ")
	fmt.Scanf("%s", &lang)

	return strings.ToLower(lang)
}

// capitalize 1st letter of string
func firstLetterCap(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// create template using all lowercase lang name
func createTemplate(lang string) (string, string) {
	ulang := firstLetterCap(lang)

	DF_PATH := fmt.Sprintf("./Dockerfiles/%sDev", ulang)
	WF_PATH := fmt.Sprintf("./.github/workflows/Build%sDev.yml", lang)

	return WF_PATH, DF_PATH
}
