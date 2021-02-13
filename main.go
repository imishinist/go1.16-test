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

func writeEmbedFS(dir embed.FS) {
	tmp, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	// defer os.RemoveAll(tmp) // clean up

	fmt.Println(tmp)

	fs.WalkDir(dir, "assets", fs.WalkDirFunc(func(path string, d fs.DirEntry, err error) error {
		cpath := filepath.Join(tmp, path)
		if d.IsDir() {
			os.Mkdir(cpath, 0755)
		} else {
			os.WriteFile(cpath, []byte(""), 0644)
		}
		return nil
	}))
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
