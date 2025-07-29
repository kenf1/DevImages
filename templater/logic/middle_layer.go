package logic

import "strings"

func replacePlaceholders(rawTemplate string, replacements map[string]string) string {
	result := rawTemplate

	for key, val := range replacements {
		placeholder := "${" + key + "}"
		result = strings.ReplaceAll(result, placeholder, val)
	}

	return result
}

func CtLogic(name, ulang, lang string) string {
	raw_template := `name: Build ${ulang}Dev

on:
  workflow_dispatch:
  push:
    branches:
      - main

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
          file: ./Dockerfiles/${ulang}_dev.dockerfile
          push: true
          platforms: linux/amd64
          tags: ghcr.io/${name}/${lang}dev:latest
`

	replacements := map[string]string{
		"name":  name,
		"ulang": ulang,
		"lang":  lang,
	}

	return replacePlaceholders(raw_template, replacements)
}
