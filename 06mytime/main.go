package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to study time package in Golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05 "))

	createdDate := time.Date(2005, time.February, 04, 23, 23, 0, 0, time.Local)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
