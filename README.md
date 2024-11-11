# Go Directory Template 

Generate go template map based on the directory. 

The conventions are:
- Base template file at `_base/__base.*` 
- Components template file at `_private/*` (will not inherit)

Use [go standard template](https://pkg.go.dev/html/template), not additional libraries required.

## Features

- [x] Support nested template
- [x] Support OS file system (relative path)
- [x] Support go fsys (go embed)
- [x] Support html/template
- [x] Support text/template
- [ ] Custom base template folder
- [ ] Custom components folder
- [ ] Exclude specific file
- [ ] Filter by file extension

## Usages 

Get entries using `dirtmpl.Entries()` or `dirtmpl.EntriesFS()`
```go
entries, err := dirtmpl.Entries("samples/simpletxt")
if err != nil {
    log.Fatal(err)
}

for _, entry := range entries {
    fmt.Println(entry.Key, entry.Filenames)
}
```

Get templates using `dirtmpl.TextTemplates()`, `dirtmpl.TextTemplatesFS()`, `dirtmpl.HTMLTemplates()`, or `dirtmpl.HTMLTemplatesFS()`
```go
m, err := dirtmpl.TextTemplates("samples/simpletxt")
if err != nil {
    log.Fatal(err)
}

tmpl, ok := m["section_b/subsection_b1.md"]
if !ok {
	log.Fatal("template not found")
}

err = tmpl.Execute(os.Stdout, nil)
if err != nil {
    log.Fatal(err)
}
```

## Examples

- [ ] Generate static html file 
- [ ] Web server with go http standard 
- [ ] Web server with echo 
- [ ] Web server with gofiber

## Author

iman.tung@gmail.com