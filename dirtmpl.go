package dirtmpl

import (
	"embed"
	"os"
)

var (
	DefaultBasePrefix = "_"
)

func Entries(root string) ([]Entry, error) {
	td := templateDir{
		Root:       root,
		ReadDir:    os.ReadDir,
		BasePrefix: DefaultBasePrefix,
	}

	return td.Entries()
}

func EntriesFS(f embed.FS, root string) ([]Entry, error) {
	td := templateDir{
		Root:       root,
		ReadDir:    f.ReadDir,
		BasePrefix: DefaultBasePrefix,
	}

	return td.Entries()
}
