package main

import (
	"io/ioutil"
)

func main() {
	

}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func renderTemplate() string {
	paths := []string{
		"template.tmpl",
	  }
	  
	  template = "template.tmpl"
	  t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	//   contents := readFile()

	  err = t.Execute(os.Stdout, contents)
	  bytesToWrite := []byte (buffer.String())
	  if err != nil {
		panic(err)
	  }
 
	return string(contents)

}
