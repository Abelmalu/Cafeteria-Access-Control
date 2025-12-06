package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	gofakeit.Seed(0)

	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.Name())
	fmt.Println(gofakeit.Email())
	fmt.Println(gofakeit.Phone())
	fmt.Println(gofakeit.City())
	fmt.Println(gofakeit.Password(true, true, true, true, false, 12))
}
