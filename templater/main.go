package main

func main() {
	//get user input
	lang := promptLang()

	//create path + content
	lang_template, WF_PATH := ctWrapper(lang)

	//write .yml
	writeYAML(WF_PATH, lang_template)
}
