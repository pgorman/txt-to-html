// Package txt-to-html formats txt or md files in a directory as HTML.
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

func header(dir string) string {
	var header string
	f := path.Join(dir, "HEADER.html")
	if _, err := os.Stat(f); os.IsNotExist(err) {
		header = `<!DOCTYPE html>
<html lang="en-us">
<head>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1" />
<link rel="stylesheet" href="default.css" />
</head>
<body>`
	} else {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			log.Println("error reading header file: ", err)
		}
		header = string(b)
	}
	return header
}

func footer(dir string) string {
	var footer string
	f := path.Join(dir, "FOOTER.html")
	if _, err := os.Stat(f); os.IsNotExist(err) {
		footer = `</body>
</html>`
	} else {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			log.Println("error reading footer file: ", err)
		}
		footer = string(b)
	}
	return footer
}

func main() {
	var dir string
	flag.Parse()
	switch len(flag.Args()) {
	case 1:
		dir = flag.Args()[0]
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			log.Fatal("error using input directory: ", err)
		}
	default:
		log.Fatal("please supply a directory for input/output")
	}

	txt, err := filepath.Glob(path.Join(dir, "*.txt"))
	if err != nil {
		log.Fatal("error finding .txt input files: ", err)
	}
	md, err := filepath.Glob(path.Join(dir, "*.md"))
	if err != nil {
		log.Fatal("error finding .md input files: ", err)
	}

	files := append(txt, md...)
	for _, f := range files {
		o, err := os.Create(path.Join(strings.Join([]string{f, "html"}, ".")))
		if err != nil {
			log.Println("error creating output file: ", err)
		}
		body := blackfriday.MarkdownCommon([]byte("foo **bar** bat"))
		_, err = o.WriteString(strings.Join([]string{header(dir), string(body), footer(dir)}, "\n"))
		if err != nil {
			log.Println("error writing to output file: ", err)
		}
		o.Close()
	}
}
