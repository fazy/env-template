package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

var templateFile string

func init() {
	if len(os.Args) != 2 {
		exitError(fmt.Errorf("Usage: %s <template>\n", os.Args[0]))
	}

	templateFile = os.Args[1]
}

func main() {
	getEnv := func(key string) string {
		value := os.Getenv(key)
		return value
	}

	funcs := template.FuncMap{"env": getEnv}

	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		exitError(err)
	}

	parsedTemplate, err := template.New("test").Funcs(funcs).Parse(string(templateContent))
	if err != nil {
		exitError(err)
	}

	var output bytes.Buffer
	err = parsedTemplate.Execute(&output, nil)
	if err != nil {
		exitError(err)
	}

	fmt.Print(output.String())
}

func exitError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
