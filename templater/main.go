package main

import (
	"fmt"
	"os"
	iofunctions "templater/io_functions"
)

func main() {
	lang, err := iofunctions.PromptLang()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	username, err := iofunctions.PromptUsername()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lang_template, WF_PATH := iofunctions.CtWrapper(username, lang)

	err1 := iofunctions.WriteYAML(WF_PATH, lang_template)
	if err1 != nil {
		fmt.Println(err1)
	}
}
