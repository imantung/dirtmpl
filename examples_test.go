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
	// ↳ samples/simpletxt/__base.md
	// ↳ samples/simpletxt/_author.md
	// ↳ samples/simpletxt/_comp/comp_x.txt
	// ↳ samples/simpletxt/section_a.md
	// section_b/subsection_b1.md
	// ↳ samples/simpletxt/__base.md
	// ↳ samples/simpletxt/_author.md
	// ↳ samples/simpletxt/section_b/__base.md
	// ↳ samples/simpletxt/section_b/_author.md
	// ↳ samples/simpletxt/section_b/_comp/comp_y.txt
	// ↳ samples/simpletxt/section_b/subsection_b1.md
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
	// ↳ samples/simpletxt/__base.md
	// ↳ samples/simpletxt/_author.md
	// ↳ samples/simpletxt/_comp/comp_x.txt
	// ↳ samples/simpletxt/section_a.md
	// section_b/subsection_b1.md
	// ↳ samples/simpletxt/__base.md
	// ↳ samples/simpletxt/_author.md
	// ↳ samples/simpletxt/section_b/__base.md
	// ↳ samples/simpletxt/section_b/_author.md
	// ↳ samples/simpletxt/section_b/_comp/comp_y.txt
	// ↳ samples/simpletxt/section_b/subsection_b1.md
}

func ExampleTextTemplates() {
	m, err := dirtmpl.TextTemplates("samples/simpletxt")
	if err != nil {
		log.Fatal(err)
	}

	err = m["section_b/subsection_b1.md"].Execute(os.Stdout, nil)
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
	//
	// By iman.tung@gmail.com
}

func ExampleTextTemplatesFS() {
	m, err := dirtmpl.TextTemplatesFS(SimpleTxtFS, "samples/simpletxt")
	if err != nil {
		log.Fatal(err)
	}

	err = m["section_a.md"].Execute(os.Stdout, nil)
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
