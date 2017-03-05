package lib

import (
	"log"
	"os"
	"text/template"
)

func Parse(inputTemplate string, outputFile string, pods map[string]string) (error) {
	_template, err := template.ParseFiles(inputTemplate)
	if err != nil {
		log.Print(err)
		return err
	}

	_outputFile, err := os.Create(outputFile)
	if err != nil {
		log.Println("create file: ", err)
		return err
	}

	templateConfig := make(map[string]map[string]string)
	templateConfig["servers"] = pods

	err = _template.Execute(_outputFile, templateConfig)
	if err != nil {
		log.Print("execute: ", err)
		return err
	}

	_outputFile.Close()

	return nil
}
