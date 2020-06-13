package main

import (
    "fmt"
    "io/ioutil"
    // "log"
    /////
    "os"
    "text/template"
)

type Todo struct {
    Name string
    Description string
}

func main() {
    data := Todo{"test templates", "hello here"}
    t, err := template.New("todos").Parse("you have task \"{{ .Name}}\" with description: \"{{ .Description}}\"")
    if err != nil {
        panic(err)
    }
    err = t.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }

    // directory := "."
    // files, err := ioutil.ReadDir(directory)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // for _, file := range files {
    //     fmt.Println(file.Name())
    // }

}
