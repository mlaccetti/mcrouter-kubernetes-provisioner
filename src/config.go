package main

import (
	"log"
	"os"
	"text/template"
)
func parse(inputTemplate string, outputFile string) {
	_template, err := template.ParseFiles(inputTemplate)
	if err != nil {
		log.Print(err)
		return
	}

	_outputFile, err := os.Create(outputFile)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	// A sample config
	config := map[string]string{
		"textColor":      "#abcdef",
		"linkColorHover": "#ffaacc",
	}

	err = _template.Execute(_outputFile, config)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

	_outputFile.Close()
}
