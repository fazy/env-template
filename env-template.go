package main

import (
    "fmt"
    "os"
    "strings"
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
    parsedTemplate, err := template.ParseFiles(templateFile)
    if err != nil { exitError(err) }

    env := getEnv()
    err = parsedTemplate.Execute(os.Stdout, &env)
    if err != nil { exitError(err) }
}

func getEnv() map[string]string {
    env := make(map[string]string)

    for _, i := range os.Environ() {
        parts := strings.Split(i, "=")
        env[parts[0]] = parts[1]
    }

    return env
}

func exitError(err error) {
    fmt.Println(err)
    os.Exit(1)
}
