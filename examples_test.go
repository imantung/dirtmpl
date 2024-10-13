package dirtmpl_test

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/imantung/dirtmpl"
)

var (
	SimpleTxtRoot = "samples/simpletxt"

	//go:embed all:samples/simpletxt/*
	SimpleTxtFS embed.FS
)

func ExampleTemplateDir_Entries_withOSReadDir() {
	td := dirtmpl.TemplateDir{
		Root:    "samples/simpletxt",
		ReadDir: os.ReadDir,
	}

	m, err := td.Entries()
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range m {
		fmt.Println(entry.Key)
		for _, filename := range entry.Filenames {
			fmt.Println("\t", filename)
		}
	}

	// Output:
	// file_a.txt
	// 	 samples/simpletxt/_base.txt
	// 	 samples/simpletxt/_footer.txt
	// 	 samples/simpletxt/file_a.txt
	// file_b.txt
	// 	 samples/simpletxt/_base.txt
	// 	 samples/simpletxt/_footer.txt
	// 	 samples/simpletxt/file_b.txt
	// level-2/file_c.txt
	// 	 samples/simpletxt/_base.txt
	// 	 samples/simpletxt/_footer.txt
	// 	 samples/simpletxt/level-2/_comp/comp_x.txt
	// 	 samples/simpletxt/level-2/_comp/comp_y.txt
	// 	 samples/simpletxt/level-2/file_c.txt
	// level-2/level-3/file_d.txt
	// 	 samples/simpletxt/_base.txt
	// 	 samples/simpletxt/_footer.txt
	// 	 samples/simpletxt/level-2/level-3/file_d.txt
}

func ExampleTemplateDir_Entries_withGoEmbed() {

	td := dirtmpl.TemplateDir{
		Root:    SimpleTxtRoot,
		ReadDir: SimpleTxtFS.ReadDir,
	}

	m, err := td.Entries()
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range m {
		fmt.Println(entry.Key)
		for _, filename := range entry.Filenames {
			fmt.Println("\t", filename)
		}
	}

	// Output:
	// file_a.txt
	// 	 samples/simpletxt/_base.txt
	// 	 samples/simpletxt/_footer.txt
	// 	 samples/simpletxt/file_a.txt
	// file_b.txt
	// 	 samples/simpletxt/_base.txt
	// 	 samples/simpletxt/_footer.txt
	// 	 samples/simpletxt/file_b.txt
	// level-2/file_c.txt
	// 	 samples/simpletxt/_base.txt
	// 	 samples/simpletxt/_footer.txt
	// 	 samples/simpletxt/level-2/_comp/comp_x.txt
	// 	 samples/simpletxt/level-2/_comp/comp_y.txt
	// 	 samples/simpletxt/level-2/file_c.txt
	// level-2/level-3/file_d.txt
	// 	 samples/simpletxt/_base.txt
	// 	 samples/simpletxt/_footer.txt
	// 	 samples/simpletxt/level-2/level-3/file_d.txt
}
