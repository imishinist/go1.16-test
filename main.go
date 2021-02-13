package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/imishinist/embet"
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
	defer os.RemoveAll(tmp) // clean up

	fmt.Println(tmp)
	if err := embet.WriteEmbedFiles(dir, "assets", tmp); err != nil {
		log.Fatal(err)
	}
}

//go:embed message.txt
var messageFile embed.FS

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

	if err := embet.WriteEmbedFiles(messageFile, "message.txt", "tmp"); err != nil {
		log.Fatal(err)
	}
}
