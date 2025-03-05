package main

import (
	"fmt"
	"os"
)

// create & write to .yml file
func writeYAML(filename string, contents string) {
	//create file
	df_file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer df_file.Close()

	//write contents
	_, err = df_file.WriteString(contents)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
