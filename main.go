package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"playground/templates/pkg/entity"
	"text/template"
)

type Conf struct {
	OutputPath string `json:"outputPath,omitempty"`
	Entity     string `json:"entity,omitempty"`
}

const (
	dataPath     = "./data/"
	templateFile = "api-template.md"
)

func main() {
	folders, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range folders {
		var data entity.ApplicationsDoc

		if !folder.IsDir() {
			continue
		}

		fn := filepath.Join(dataPath, folder.Name(), "config.json")
		confJSON, err := ioutil.ReadFile(fn)
		if err != nil {
			log.Fatal(folder.Name() + " => ioutil.ReadFile => " + err.Error())
		}

		config := &Conf{}
		err = json.Unmarshal(confJSON, config)
		if err != nil {
			log.Fatal(folder.Name() + " => json.Unmarshal => " + err.Error())
		}

		entries, err := ioutil.ReadDir(dataPath + folder.Name())
		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
				continue
			}

			fn := filepath.Join(dataPath, folder.Name(), entry.Name())
			jsonFile, err := ioutil.ReadFile(fn)
			if err != nil {
				log.Fatal(entry.Name() + " => ioutil.ReadFile => " + err.Error())
			}

			appDoc := &entity.ApplicationDoc{}
			err = json.Unmarshal(jsonFile, appDoc)
			if err != nil {
				log.Fatal(entry.Name() + " => json.Unmarshal => " + err.Error())
			}

			data = append(data, *appDoc)
		}

		tpl := template.Must(template.New(templateFile).ParseFiles(templateFile))
		generate(tpl, config.OutputPath, data)
	}
}

func generate(template *template.Template, outputFileName string, data interface{}) {
	os.MkdirAll(path.Dir(outputFileName), os.ModePerm)

	outputFile, err := os.Create(outputFileName)
	log.Println("Generating file: " + outputFileName)
	if err != nil {
		log.Fatal(outputFileName + " => os.Create => " + err.Error())
	}
	defer outputFile.Close()

	err = template.Execute(outputFile, data)
	if err != nil {
		log.Fatal("Execute => " + err.Error())
	}
}
