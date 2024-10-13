# Go Directory Template 

Generate go template map based on the directory. 

The conventions are:
- Base template use `__base` filename
- Other base components (inherited) use the prefix `_` filename
- Other template components (not inherited) stored in the `_comp` folder

Use [go standard template](https://pkg.go.dev/html/template), not additional libraries required.

## Features

- [x] Support nested template
- [x] Support OS file system (relative path)
- [x] Support go fsys (go embed)
- [ ] Support html/template
- [x] Support text/template
- [ ] Custom prefix
- [ ] Exclude specific file
- [ ] Filter by file extension

## Usages 

Print entries
```go
entries, err := dirtmpl.Entries("samples/simpletxt") // or use `dirtmpl.EntriesFS()` for go embed
if err != nil {
    log.Fatal(err)
}

for _, entry := range entries {
    fmt.Println(entry.Key, entry.Filenames)
}
```

Execute template
```go
m, err := dirtmpl.TextTemplate("samples/simpletxt") // or use `dirtmpl.TextTemplateFS()` for go embed
if err != nil {
    log.Fatal(err)
}

err = m["section_b/subsection_b1.md"].Execute(os.Stdout, nil)
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