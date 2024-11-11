package dirtmpl

import (
	"io/fs"
	"path/filepath"
)

type (
	templateDir struct {
		Root    string
		ReadDir ReadDirFn
	}

	ReadDirFn func(name string) ([]fs.DirEntry, error)

	Entry struct {
		Key       string
		Filenames []string
	}
)

func (t *templateDir) Entries() ([]Entry, error) {
	entries := []Entry{}
	t.walk(&entries, t.Root, []string{})
	return entries, nil
}

func (t *templateDir) walk(e *[]Entry, parent string, bases []string) {
	entries, _ := t.ReadDir(parent)

	var comps []string
	for _, entry := range entries {
		name := entry.Name()
		path := filepath.Join(parent, name)

		if entry.IsDir() {
			if name == "_base" {
				entries, _ := t.ReadDir(path)
				for _, entry := range entries {
					if !entry.IsDir() {
						bases = append(bases, filepath.Join(parent, name, entry.Name()))
					}
				}

			} else if name == "_private" {
				entries, _ := t.ReadDir(path)
				for _, entry := range entries {
					if !entry.IsDir() {
						comps = append(comps, filepath.Join(parent, name, entry.Name()))
					}
				}
			} else {
				t.walk(e, path, bases)
			}
		} else {
			bases = append(bases, path)
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
