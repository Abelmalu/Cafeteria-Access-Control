package main

import (
	"fmt"
	"time"
)

func main() {
	var currentTime time.Time

nextYear := "18:30:00"
currentTime = time.Now()

newTime,err := time.Parse(time.TimeOnly,nextYear)
if err != nil{

	fmt.Println(err)
}

	fmt.Println(currentTime)

	fmt.Println("time from string")
	fmt.Println(newTime)

}