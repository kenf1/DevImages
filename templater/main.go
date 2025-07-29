package main

import iofunctions "templater/io_functions"

func main() {
	lang := promptLang()
	lang_template, WF_PATH := CtWrapper(lang)
	iofunctions.WriteYAML(WF_PATH, lang_template)
}
