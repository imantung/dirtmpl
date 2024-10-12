# Go Directory Template 

Generate go template map based on the directory. 

The conventions are:
- Base templates start with the prefix `_`, which will be inherited by child templates
- Component per directory stored in the `_comp` folder, it will not be inherited by child templates

Use [go standard template](https://pkg.go.dev/html/template), not additional libraries required.

## Features

- [ ] Support OS file system (relative path)
- [ ] Support go fsys (go embed)
- [ ] Support html/template
- [ ] Support text/template
- [ ] Exclude specific file
- [ ] Filter by file extension

## Usages 

WIP

## Examples

- [ ] Print map examples
- [ ] Generate static html file 
- [ ] Web server with go http standard 
- [ ] Web server with echo 
- [ ] Web server with gofiber

## Author

iman.tung@gmail.com