package dirtmpl_test

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/imantung/dirtmpl"
)

var (
	//go:embed all:samples/simpletxt/*
	SimpleTxtFS embed.FS
)

func ExampleEntries_withOSReadDir() {
	entries, err := dirtmpl.Entries("samples/simpletxt")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Println(entry.Key)
		for _, filename := range entry.Filenames {
			fmt.Println("↳", filename)
		}
	}

	// Output:
	// section_a.md
	// ↳ samples/simpletxt/_base/__base.md
	// ↳ samples/simpletxt/_base/_author.md
	// ↳ samples/simpletxt/section_a.md
	// ↳ samples/simpletxt/_private/comp_x.txt
	// ↳ samples/simpletxt/section_a.md
	// section_b/sub_section_b1.md
	// ↳ samples/simpletxt/_base/__base.md
	// ↳ samples/simpletxt/_base/_author.md
	// ↳ samples/simpletxt/section_a.md
	// ↳ samples/simpletxt/section_b/_base/__base.md
	// ↳ samples/simpletxt/section_b/sub_section_b1.md
	// ↳ samples/simpletxt/section_b/_private/comp_y.txt
	// ↳ samples/simpletxt/section_b/sub_section_b1.md
}

func ExampleEntries_withGoEmbed() {
	entries, err := dirtmpl.EntriesFS(SimpleTxtFS, "samples/simpletxt")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Println(entry.Key)
		for _, filename := range entry.Filenames {
			fmt.Println("↳", filename)
		}
	}

	// Output:
	// section_a.md
	// ↳ samples/simpletxt/_base/__base.md
	// ↳ samples/simpletxt/_base/_author.md
	// ↳ samples/simpletxt/section_a.md
	// ↳ samples/simpletxt/_private/comp_x.txt
	// ↳ samples/simpletxt/section_a.md
	// section_b/sub_section_b1.md
	// ↳ samples/simpletxt/_base/__base.md
	// ↳ samples/simpletxt/_base/_author.md
	// ↳ samples/simpletxt/section_a.md
	// ↳ samples/simpletxt/section_b/_base/__base.md
	// ↳ samples/simpletxt/section_b/sub_section_b1.md
	// ↳ samples/simpletxt/section_b/_private/comp_y.txt
	// ↳ samples/simpletxt/section_b/sub_section_b1.md
}

func ExampleTextTemplates() {
	m, err := dirtmpl.TextTemplates("samples/simpletxt")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, ok := m["section_b/sub_section_b1.md"]
	if !ok {
		log.Fatal("template not found")
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// # Simple Text Example
	//
	// ## Section B
	//
	// ### Subsection B1
	// Component Y
	// Component Y
	//
	//
	// ## Author
	// By iman.tung@gmail.com
}

func ExampleTextTemplatesFS() {
	m, err := dirtmpl.TextTemplatesFS(SimpleTxtFS, "samples/simpletxt")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, ok := m["section_a.md"]
	if !ok {
		log.Fatal("template not found")
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// # Simple Text Example
	//
	// ## Section A
	// Component X
	// Component X
	//
	// ## Author
	// By iman.tung@gmail.com
}
