package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

// generate untuk mengcopy file di luar folder main ke dalam folder main
// ada error vs code dan bisa hilang sendiri

//go:generate cp -r ../version.txt ./version.txt
//go:embed version.txt
var version string

//go:generate cp -r ../logo.png ./logo.png
//go:embed logo.png
var logo []byte

//go:generate cp -r ../files ./files
//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("logo_new_main.png", logo, fs.ModePerm) // logo_new.png munculnya di luar folder main
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
