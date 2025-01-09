package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func CreateFileIfNotExists(fileName string, tmpl *template.Template, data interface{}) {
	dir := filepath.Dir(fileName)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("error creating directory: %v", err)
	}

	_, stateFileErr := os.Stat(fileName)

	if os.IsNotExist(stateFileErr) {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("error creating file : %v", err)
		}
		defer file.Close()

		err = tmpl.Execute(file, data)
		if err != nil {
			log.Fatalf("error writing file : %v", err)
		}

	} else {
		fmt.Println("File already exists!")
	}
}
