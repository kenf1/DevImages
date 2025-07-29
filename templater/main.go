package main

import iofunctions "templater/io_functions"

func main() {
	lang := iofunctions.PromptLang()
	lang_template, WF_PATH := iofunctions.CtWrapper(lang)
	iofunctions.WriteYAML(WF_PATH, lang_template)
}
