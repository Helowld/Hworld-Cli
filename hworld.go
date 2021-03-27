package main

import (
	"io/ioutil"

	"github.com/thatisuday/commando"
)

const (
	filePath string = "README.md"
)

func main() {
	read, err := ioutil.ReadFile(filePath)
	isError(err)
	content := string(read)

	commando.
		SetExecutableName("hworld").
		SetVersion("0.0.2").
		SetDescription("Use this tool to update the Readme file after adding new language")

	add(content) // function for adding new language
	search()     // function for searching the language

	commando.Parse(nil)
}
