package dirtmpl

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
)

type (
	TemplateDir struct {
		Root       string
		BasePrefix string
		ReadDir    ReadDirFn
	}

	ReadDirFn func(name string) ([]fs.DirEntry, error)

	Entry struct {
		Key       string
		Filenames []string
	}
)

func (t *TemplateDir) validate() error {
	if t.Root == "" {
		return errors.New("'Root' can't be empty")
	}
	if t.ReadDir == nil {
		return errors.New("'ReadDir' can't be nil")
	}
	if t.BasePrefix == "" {
		t.BasePrefix = "_"
	}
	return nil
}

func (t *TemplateDir) Entries() ([]Entry, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	entries := []Entry{}
	t.walk(&entries, t.Root, []string{})
	return entries, nil
}

func (t *TemplateDir) walk(e *[]Entry, parent string, bases []string) {
	entries, _ := t.ReadDir(parent)

	var comps []string
	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(parent, name)

		if strings.HasPrefix(name, t.BasePrefix) {
			if entry.IsDir() {
				entries, _ := t.ReadDir(path)
				for _, entry := range entries {
					if !entry.IsDir() {
						comps = append(comps, filepath.Join(parent, name, entry.Name()))
					}
				}
			} else {
				bases = append(bases, path)
			}
		} else {
			if entry.IsDir() {
				t.walk(e, path, bases)
			} else {
				includes := make([]string, len(bases))
				copy(includes, bases)
				includes = append(includes, comps...)
				*e = append(*e, Entry{
					Key:       path[len(t.Root)+1:],
					Filenames: append(includes, path),
				})
			}
		}
	}
}
