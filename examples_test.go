package dirtmpl_test

import (
	"embed"
	"fmt"
	"log"

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
			fmt.Println("\t", filename)
		}
	}

	// Output:
	// section_a.md
	// 	 samples/simpletxt/_author.md
	// 	 samples/simpletxt/_base.md
	// 	 samples/simpletxt/_comp/comp_x.txt
	// 	 samples/simpletxt/section_a.md
	// section_b/subsection_b1.md
	// 	 samples/simpletxt/_author.md
	// 	 samples/simpletxt/_base.md
	// 	 samples/simpletxt/section_b/_base.md
	// 	 samples/simpletxt/section_b/_comp/comp_y.txt
	// 	 samples/simpletxt/section_b/subsection_b1.md
}

func ExampleEntries_withGoEmbed() {
	entries, err := dirtmpl.EntriesFS(SimpleTxtFS, "samples/simpletxt")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Println(entry.Key)
		for _, filename := range entry.Filenames {
			fmt.Println("\t", filename)
		}
	}

	// Output:
	// section_a.md
	// 	 samples/simpletxt/_author.md
	// 	 samples/simpletxt/_base.md
	// 	 samples/simpletxt/_comp/comp_x.txt
	// 	 samples/simpletxt/section_a.md
	// section_b/subsection_b1.md
	// 	 samples/simpletxt/_author.md
	// 	 samples/simpletxt/_base.md
	// 	 samples/simpletxt/section_b/_base.md
	// 	 samples/simpletxt/section_b/_comp/comp_y.txt
	// 	 samples/simpletxt/section_b/subsection_b1.md
}
