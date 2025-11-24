package main

import (
	"fmt"
	"time"
)

func expensiveCall() {

	fmt.Println("this will be delayed")
}

func main() {

	fmt.Println("hellow ")

	time.Sleep(10 * time.Second)

	expensiveCall()

}
