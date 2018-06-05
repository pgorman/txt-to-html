// Package txt-to-html formats txt or md files in a directory as HTML.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// anchor turns a file name into an HTML link/anchor.
func anchor(s string) string {
	return strings.Join([]string{`<li><a href="`, s, `">`, s, "</a></li>"}, "")
}

// header returns the start of an HTML document.
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

// footer returns the end of an HTML document.
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
	makeIndex := flag.Bool("i", false, "Generate index.html that lists the directory contents.")
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

	indexLinks := make([]string, 0, 100)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println("failed to read contents of input/output directory")
	}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") || strings.HasPrefix(f.Name(), "index.") || strings.HasPrefix(f.Name(), "HEAD") || strings.HasPrefix(f.Name(), "FOOT") {
			continue
		}
		if *makeIndex && f.IsDir() {
			indexLinks = append(indexLinks, anchor(strings.Join([]string{f.Name(), "/"}, "")))
			continue
		}
		if strings.HasSuffix(f.Name(), ".txt") || strings.HasSuffix(f.Name(), ".md") {
			b, err := ioutil.ReadFile(path.Join(dir, f.Name()))
			if err != nil {
				log.Println("error reading input file: ", err)
			}
			body := string(blackfriday.Run(b, blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs)))
			o, err := os.Create(path.Join(dir, strings.Join([]string{f.Name(), "html"}, ".")))
			if err != nil {
				log.Println("error creating output file: ", err)
			}
			_, err = o.WriteString(strings.Join([]string{header(dir), body, footer(dir)}, "\n"))
			if err != nil {
				log.Println("error writing to output file: ", err)
			}
			o.Close()
		}
		if *makeIndex && !strings.HasSuffix(f.Name(), ".txt") && !strings.HasSuffix(f.Name(), ".md") {
			indexLinks = append(indexLinks, anchor(f.Name()))
		}
	}

	if *makeIndex {
		sort.Strings(indexLinks)
		indexFile, err := os.Create(path.Join(dir, "index.html"))
		if err != nil {
			log.Println("error creating index.html file: ", err)
		}
		fmt.Fprintf(indexFile, header(dir))
		fmt.Fprintln(indexFile, `<ul class="filelist">`)
		for _, a := range indexLinks {
			fmt.Fprintln(indexFile, a)
		}
		fmt.Fprintln(indexFile, "</ul>")
		fmt.Fprintf(indexFile, footer(dir))
		indexFile.Close()
	}
}
