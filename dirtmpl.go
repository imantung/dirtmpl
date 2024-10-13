package dirtmpl

import (
	"embed"
	"os"
	txtTmpl "text/template"
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

func TextTemplates(root string) (map[string]*txtTmpl.Template, error) {
	entries, err := Entries(root)
	if err != nil {
		return nil, err
	}
	var m map[string]*txtTmpl.Template = make(map[string]*txtTmpl.Template)
	for _, entry := range entries {
		m[entry.Key], err = txtTmpl.ParseFiles(entry.Filenames...)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

func TextTemplatesFS(f embed.FS, root string) (map[string]*txtTmpl.Template, error) {
	entries, err := Entries(root)
	if err != nil {
		return nil, err
	}
	var m map[string]*txtTmpl.Template = make(map[string]*txtTmpl.Template)
	for _, entry := range entries {
		m[entry.Key], err = txtTmpl.ParseFS(f, entry.Filenames...)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
