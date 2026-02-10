package main

import "fmt"

func main() {
	// defer works on Last In First Out (LIFO)
	defer fmt.Println("World") // defer keyword removes this line from here and puts it at the last of hte main function (just before "}")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hi")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
