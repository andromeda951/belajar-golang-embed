package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

// generate untuk mengcopy file di luar folder main ke dalam folder main
// lalu kita bisa gunakan untuk di embed di file ini
// file new_version.txt, new_logo.png, dan new_files/*.txt awalnya belum ada di folder main
// setelah kita generate maka akan ada filenya yang bersumber dari luar folder main dan kita bisa embed di file ini

// ada error vs code dan bisa hilang sendiri

//go:generate cp -r ../version.txt ./new_version.txt
//go:embed new_version.txt
var version string

//go:generate cp -r ../logo.png ./new_logo.png
//go:embed new_logo.png
var logo []byte

//go:generate cp -r ../files ./new_files
//go:embed new_files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	// logo_new_main.png muncul diluar folder main jika di run diluar main
	// jika di run di dalam folder main maka logo_new_main.png akan mulcul di dalam folder main
	err := os.WriteFile("logo_new_main.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("new_files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("new_files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
