package dirtmpl

import (
	"embed"
	htmlTmpl "html/template"
	"os"
	txtTmpl "text/template"
)

func Entries(root string) ([]Entry, error) {
	td := templateDir{
		Root:    root,
		ReadDir: os.ReadDir,
	}
	return td.Entries()
}

func EntriesFS(f embed.FS, root string) ([]Entry, error) {
	td := templateDir{
		Root:    root,
		ReadDir: f.ReadDir,
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
	entries, err := EntriesFS(f, root)
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

func HTMLTemplates(root string) (map[string]*htmlTmpl.Template, error) {
	entries, err := Entries(root)
	if err != nil {
		return nil, err
	}
	var m map[string]*htmlTmpl.Template = make(map[string]*htmlTmpl.Template)
	for _, entry := range entries {
		m[entry.Key], err = htmlTmpl.ParseFiles(entry.Filenames...)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

func HTMLTemplatesFS(f embed.FS, root string) (map[string]*htmlTmpl.Template, error) {
	entries, err := EntriesFS(f, root)
	if err != nil {
		return nil, err
	}
	var m map[string]*htmlTmpl.Template = make(map[string]*htmlTmpl.Template)
	for _, entry := range entries {
		m[entry.Key], err = htmlTmpl.ParseFS(f, entry.Filenames...)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}
