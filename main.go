package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed message.txt
var message string

//go:embed assets/html/*
var assets embed.FS

func main() {
	fmt.Println(message)

	htmlfiles, err := fs.ReadDir(assets, "assets/html")
	if err != nil {
		log.Fatal(err)
	}
	for _, html := range htmlfiles {
		fmt.Printf("%q\n", html.Name())
	}
}
