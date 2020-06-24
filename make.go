package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"html/template"
	"flag"
	"strings"
	"github.com/labstack/gommon/color"
)

// data where filecontents is stored, dictionary
type content struct {
	Title string
	FileData string
}


func readFile(filename string) string {
	// var content string
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func renderTemplate(filename string, cont string) {
	// var filename string
	var err error
	// bytesToWrite := []byte(filename)
	con := content{FileData: cont}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))
	// err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	err = t.Execute(os.Stdout, con)
	if err != nil {
		panic(err)
	}
	// fmt.Println(err)
	// return fmt.Sprintf("%s file created", err)
}

func extensionAdd(filename string) string {
	filext := strings.Split(filename, ".")[0] + ".html"
	return filext
}

func fileToTemplate(temp string, cont string) {
	con := content{FileData: readFile(cont)}
	t := template.Must(template.New("template.tmpl").ParseFiles(temp))
	
	file := extensionAdd(cont)
	filee, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	err = t.Execute(filee, con)
	if err != nil {
		panic(err)
	}
}

//	check if existent file
func checkFile(filename string) bool {
	if strings.Contains(filename, "."){
		return strings.Split(filename, ".")[1] == "txt"
	} else {
		return false
	}
}
func save() {
	// var err error
	// var content string
	// var filename *string
	// var dirname *string
	// FLAGS
	// var filename string
	// var dirname string

	fileFlag := flag.String("file", "", " latest-post.txt")
	dirFlag := flag.String("dir", "", " all .txt files")
	// flag.StringVar(&filename, "file", "", " latest-post.txt")
	// flag.StringVar(&dirname, "dir", "", " all .txt files")
	flag.Parse()
	// fmt.Println("file: ", *filename)
	if *dirFlag != ""{
		var allFiles = 0
		files, err := ioutil.ReadDir(*dirFlag)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			fname := f.Name()
			// check if extension already exists
			if checkFile(fname) == true {
				renderTemplate("template.tmpl", readFile(fname))
				fileToTemplate("template.tmpl", fname)
				allFiles += 1
				// fmt.Println(*dirname)
			}

		}
		fmt.Printf("new files created")

	}
	if *fileFlag != "" {
		renderTemplate("template.tmpl", readFile(*fileFlag))
		fmt.Println("file: ", *fileFlag)
	} 
	// else {
	// 	err := filepath.Walk(".",
	// 		func(path string, info os.FileInfo, err error) error {
	// 			if err != nil {
	// 				return err
	// 			}
	// 			fmt.Println(path, info.Size())
	// 			return nil
	// 		})
	// 	fmt.Println("directory", *dirname)
	// }

	// if *dirname != ""{
	// 	var allFiles = 0
	// 	files, err := ioutil.ReadDir(*dirname)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for _, f := range files {
	// 		filename := f.Name()
	// 		renderTemplate("template.tmpl", readFile(name))
	// 		fmt.Println(*dirname)
	// 		}
	// 	}
		// err := filepath.Walk(".",
		// 	func(path string, info os.FileInfo, err error) error {
		// 		if err != nil {
		// 			return err
		// 		}
		// 		fmt.Println(path, info.Size())
		// 		return nil
		// 	})
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// allFiles, _ := ioutil.ReadDir(".")
		// // if err != nil {
		// // 	log.Fatal(err)
		// // }
		// for _, f := range allFiles {
		// 	fmt.Println(f)
		// }
	}

	/////////////////////////
	// fileContents, err := ioutil.ReadFile("onepost.txt")
	// fmt.Println(string(fileContents))

	// content := FileData {
	// 	Title: "first-post",
	// 	Content: string(fileContents),
	// }
	// t := template.Must(template.ParseFiles("template.tmpl"))

	// newfile, err := os.Create("onepost.html")
	// err = t.Execute(os.Stdout, content)
	// err = t.Execute(newfile, content)
	// if err != nil {
	// 	panic(err)
	// }
// }

func main() {
	color.Println(color.Red("Welcome to the Static Site Generator!"))
	save()
	// return FileData{Content: string(content)}
}