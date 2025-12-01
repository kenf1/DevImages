package main

import (
	"log"

	iofunctions "templater/io_functions"
)

func main() {
	lang, err := iofunctions.PromptLang()
	if err != nil {
		log.Fatal(err)
	}

	username, err := iofunctions.PromptUsername()
	if err != nil {
		log.Fatal(err)
	}

	lang_template, WF_PATH := iofunctions.CtWrapper(username, lang)

	err1 := iofunctions.WriteYAML(WF_PATH, lang_template)
	if err1 != nil {
		log.Fatal(err1)
	}
}
