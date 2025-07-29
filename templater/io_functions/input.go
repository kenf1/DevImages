package iofunctions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func sharedPromptText(promptText, emptyValueText string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(promptText)

	input_value, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	input_value = strings.TrimSpace(input_value)
	if input_value == "" {
		return "", fmt.Errorf("%s", emptyValueText)
	}

	return input_value, nil
}

func PromptLang() (string, error) {
	lang, err := sharedPromptText("Enter language name: ", "no language name entered")
	if err != nil {
		return "", err
	}

	return strings.ToLower(lang), nil
}

func PromptUsername() (string, error) {
	username, err := sharedPromptText("Enter username: ", "no username entered")
	if err != nil {
		return "", err
	}

	return username, nil
}

func capFirstLetter(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", fmt.Errorf("empty string provided")
	}

	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size == 0 {
		return "", fmt.Errorf("invalid input string")
	}

	return strings.ToUpper(string(r)) + s[size:], nil
}
