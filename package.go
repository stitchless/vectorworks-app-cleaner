package main

import (
	"archive/zip"
	"bytes"
	"log"
)

func main() {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	var files = []struct {
		Name, Body string
	} {
		{"rest", "Drugs"},
		{"Rust", "Games"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
