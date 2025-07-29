package iofunctions

import (
	"fmt"
	"os"
)

func WriteYAML(filename string, contents string) error {
	df_file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		err := df_file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	_, err = df_file.WriteString(contents)
	if err != nil {
		return err
	}

	return nil
}

func CtWrapper(lang string) (string, string) {
	ulang := firstLetterCap(lang)
	lang_template := ctLogic(ulang, lang)
	WF_PATH := fmt.Sprintf("../.github/workflows/Build%sDev.yml", ulang)

	return lang_template, WF_PATH
}
