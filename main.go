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

func WriteEmbedFiles(dir embed.FS, prefix, dest string) error {
	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dpath := filepath.Join(dest, path)
		if d.IsDir() {
			if err := os.Mkdir(dpath, DirPerm); err != nil {
				return err
			}
			return nil
		}

		// if file
		data, err := fs.ReadFile(dir, path)
		if err != nil {
			return err
		}
		if err := os.WriteFile(dpath, data, FilePerm); err != nil {
			return err
		}
		return nil
	}
	if err := fs.WalkDir(dir, prefix, walkFunc); err != nil {
		return err
	}
	return nil
}

func writeEmbedFS(dir embed.FS) {
	tmp, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmp) // clean up

	fmt.Println(tmp)
	if err := WriteEmbedFiles(dir, "assets", tmp); err != nil {
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
