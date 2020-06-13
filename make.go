package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"html/template"
	"flag"
	"strings"
)

// data where filecontents is stored, dictionary
type FileData struct {
	Title string
	Content string
}

func main() {
	save()
	// return FileData{Content: string(content)}
}


func readFile(filename string) FileData {
	var content string
	fileContents, err := ioutil.ReadFile("latest-post.txt")
	if err != nil {
		panic(err)
	}
	content = string(fileContents)
	return FileData{Content: strings.TrimSpace(string(content))}
}

func renderTemplate() string{
	var filename string
	bytesToWrite := []byte(filename)
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	return fmt.Sprintf("%s file created", err)
}

func save() {
	var err error
	// var content string
	
	// FLAGS
	filename := flag.String("file", "", " the name of your .txt file")
	flag.Parse()
	fmt.Println("file: ", *filename)

	fileContents, err := ioutil.ReadFile("latest-post.txt")
	fmt.Println(string(fileContents))
	////////

	content := FileData {
		Title: "latest-post",
		Content: string(fileContents),
	}
	t := template.Must(template.ParseFiles("template.tmpl"))

	newfile, err := os.Create("latest-post.html")
	err = t.Execute(os.Stdout, content)
	err = t.Execute(newfile, content)
	if err != nil {
		panic(err)
	}
}