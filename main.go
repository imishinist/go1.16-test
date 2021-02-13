package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

//go:embed message.txt
var message string

//go:embed assets/html/*
var assets embed.FS

var (
	DirPerm = os.FileMode(0755)
	FilePerm = os.FileMode(0644)
)

func writeEmbedFS(dir embed.FS) {
	tmp, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmp) // clean up

	fmt.Println(tmp)
	err = fs.WalkDir(dir, "assets", func(path string, d fs.DirEntry, err error) error {
		cpath := filepath.Join(tmp, path)
		if d.IsDir() {
			if err := os.Mkdir(cpath, DirPerm); err != nil {
				return err
			}
			return nil
		}

		// if file
		data, err := fs.ReadFile(dir, path)
		if err != nil {
			return err
		}
		if err := os.WriteFile(cpath, data, FilePerm); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println(message)

	htmlfiles, err := fs.ReadDir(assets, "assets/html")
	if err != nil {
		log.Fatal(err)
	}
	for _, html := range htmlfiles {
		fmt.Printf("%q\n", html.Name())
	}
	writeEmbedFS(assets)
}
