package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed version.txt
// var version string

//go:embed logo.png
// var logo []byte

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

//go:embed files/*.txt
var path embed.FS

func main() {
	// fmt.Println(version)

	// err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	// if err != nil {
	// 	panic(err)
	// }

	// a, _ := files.ReadFile("files/a.txt")
	// fmt.Println(string(a))

	// b, _ := files.ReadFile("files/b.txt")
	// fmt.Println(string(b))

	// c, _ := files.ReadFile("files/c.txt")
	// fmt.Println(string(c))

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content: ", string(content))
		}
	}
}
