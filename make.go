package main

import (
	// s"fmt"
	"io/ioutil"
)

func main() {
	ex := readFile("first-post.html")
	return ex
}

// data where filecontents is stored, dictionary
type data struct {
	content string
}
func readFile(name string) data {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	var content string
	content = string(fileContents)
	return data{content: string(content)}
}

func renderTemplate() string{
	return "hello"
}