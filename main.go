package main

import (
	"fmt"
	"path/filepath"

	// // "io"
	_ "embed"
	// "os"
)

//go:embed sql/migrations/example.txt
var example string

func main() {

	p := filepath.Join("users", "abel", "docs")
	fmt.Println(p)

	absolutePath, _ := filepath.Abs("save.txt")

	fmt.Println(absolutePath)

}
