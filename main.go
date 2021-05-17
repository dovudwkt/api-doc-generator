// Golang program to illustrate the
// concept of text/templates
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

// get the json file names
// open each file and append the content to the struct
// open template and generate the file

type Doc struct {
	Title     string `json:"title,omitempty"`
	Endpoints []struct {
		EndpointName   string   `json:"endpoint_name,omitempty"`
		Description    string   `json:"description,omitempty"`
		Method         string   `json:"method,omitempty"`
		URL            string   `json:"url,omitempty"`
		RequestBody    string   `json:"request_body,omitempty"`
		RequestFields  []Field  `json:"request_fields,omitempty"`
		Params         []string `json:"params,omitempty"`
		Response       string   `json:"response,omitempty"`
		ResponseFields []Field  `json:"response_fields,omitempty"`
	} `json:"endpoints,omitempty"`
}
type Field struct {
	Field       string `json:"field,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
}
type Docs []Doc

const (
	dataPath     = "./data/"
	templateFile = "api-template.md"
)

func main() {
	var data Docs

	// get data file names
	entries, err := ioutil.ReadDir(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		fn := filepath.Join(dataPath, entry.Name())
		jsonFile, err := ioutil.ReadFile(fn)
		if err != nil {
			fmt.Println(err)
		}

		singleDoc := &Doc{}
		err = json.Unmarshal(jsonFile, singleDoc)
		if err != nil {
			fmt.Println(err)
		}

		data = append(data, *singleDoc)
	}

	tpl := template.Must(template.New(templateFile).ParseFiles(templateFile))

	generateFile(tpl, "./out.md", data)

}

func generateFile(template *template.Template, outputFileName string, data interface{}) {
	os.MkdirAll(path.Dir(outputFileName), os.ModePerm)

	outputFile, err := os.Create(outputFileName)
	fmt.Println("Generating file : " + outputFileName)
	defer outputFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	err = template.Execute(outputFile, data)
	if err != nil {
		fmt.Println(err)
	}
}

// ----------------------------v1----------------------------------------------------

// type Doc struct {
// 	Title     string `json:"title,omitempty"`
// 	Endpoints []struct {
// 		EndpointName   string   `json:"endpoint_name,omitempty"`
// 		Description    string   `json:"description,omitempty"`
// 		Method         string   `json:"method,omitempty"`
// 		URL            string   `json:"url,omitempty"`
// 		RequestBody    string   `json:"request_body,omitempty"`
// 		RequestFields  []Field  `json:"request_fields,omitempty"`
// 		Params         []string `json:"params,omitempty"`
// 		Response       string   `json:"response,omitempty"`
// 		ResponseFields []Field  `json:"response_fields,omitempty"`
// 	} `json:"endpoints,omitempty"`
// }

// type Docs []Doc

// type Field struct {
// 	Field       string `json:"field,omitempty"`
// 	Description string `json:"description,omitempty"`
// 	Required    bool   `json:"required,omitempty"`
// }

// // main function
// func main() {
// 	var data Docs
// 	fileName := "./data.json"

// 	jsonFile, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer jsonFile.Close()

// 	fmt.Println("openning ", fileName, "...")
// 	byteValue, err := ioutil.ReadAll(jsonFile)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = json.Unmarshal(byteValue, &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	tpl := template.Must(template.New("api-template.tpl").ParseFiles("api-template.tpl"))

// 	generateFile(tpl, "./out.md", data)

// }

// func generateFile(template *template.Template, outputFileName string, data interface{}) {
// 	os.MkdirAll(path.Dir(outputFileName), os.ModePerm)

// 	outputFile, err := os.Create(outputFileName)
// 	fmt.Println("Generating file : " + outputFileName)
// 	defer outputFile.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// --------------------------------
// 	// outputFile, err := os.OpenFile(outputFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// defer outputFile.Close()
// 	// --------------------------------

// 	err = template.Execute(outputFile, data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
