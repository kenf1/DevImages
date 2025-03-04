package main

import (
	"fmt"
	"strings"
)

func main() {
	lang := promptLang()
	fmt.Println("Entered:", lang)

	lang_template, WF_PATH := ctWrapper(lang)
	fmt.Println(WF_PATH, "\n", lang_template)
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
	return strings.ToUpper(s[:1]) + s[1:]
}

// create template wrapper (only req user input lang)
func ctWrapper(lang string) (string, string) {
	ulang := firstLetterCap(lang)
	lang_template := ctLogic(ulang, lang)
	WF_PATH := fmt.Sprintf("./.github/workflows/Build%sDev.yml", ulang)

	return lang_template, WF_PATH
}

// create template logic (to be used inside ctWrapper) (wip)
func ctLogic(ulang string, lang string) string {
	raw_template := `name: Build ${ulang}Dev

	on:
	  workflow_dispatch:
	
	jobs:
	  build-and-push:
		runs-on: ubuntu-latest
		steps:
		  - name: Checkout code
			uses: actions/checkout@v4
	
		  - name: Setup QEMU
			uses: docker/setup-qemu-action@v3
	
		  - name: Setup Docker Buildx
			uses: docker/setup-buildx-action@v3
	
		  - name: Login to GHCR
			uses: docker/login-action@v3
			with:
			  registry: ghcr.io
			  username: ${{ github.actor }}
			  password: ${{ secrets.GITHUB_TOKEN }}
	
		  - name: Build and push Docker image
			uses: docker/build-push-action@v6
			with:
			  context: .
			  file: ./Dockerfiles/${ulang}Dev
			  push: true
			  platforms: linux/amd64
			  tags: ghcr.io/kenf1/${lang}dev:latest`

	//replace placeholders (ulang, lang)
	return strings.ReplaceAll(
		strings.ReplaceAll(raw_template, "${ulang}", ulang),
		"${lang}", lang,
	)
}
