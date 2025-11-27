package main

import (
	"fmt"
	"io"
	"os"

	// // "io"
	_ "embed"
	// "os"
)

//go:embed sql/migrations/example.txt
var example string

func main() {

	// _, err := os.Create("c:\\users\\abell\\desktop\\mydesktop.txt")
	// //file := os.Create("")

	// if err != nil {

	// 	fmt.Println(err)
	// }

	fmt.Println(example)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("this is using os.Open")
	file, _ := os.Open("sql/migrations/example.txt")
	os.Create("save.txt")

	fileData, _ := io.ReadAll(file)

	fmt.Println(string(fileData))

}
