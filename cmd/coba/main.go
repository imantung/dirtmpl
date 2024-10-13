package main

import (
	"fmt"
	"log"
	"os"

	txtTmpl "text/template"

	"github.com/imantung/dirtmpl"
)

func main() {

	// template, err := template.ParseFiles("samples/simpletxt/_base.md", "samples/simpletxt/_comp/comp_x.txt", "samples/simpletxt/_author.md", "samples/simpletxt/section_a.md")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = template.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	entries, err := dirtmpl.Entries("samples/simpletxt")
	if err != nil {
		log.Fatal(err)
	}
	// var m map[string]*txtTmpl.Template = make(map[string]*txtTmpl.Template)
	for _, entry := range entries {
		fmt.Println(entry.Filenames)
		template, err := txtTmpl.ParseFiles(entry.Filenames...)
		if err != nil {
			log.Fatal(err)
		}
		err = template.Execute(os.Stdout, nil)
		if err != nil {
			log.Fatal(err)
		}

	}
}
