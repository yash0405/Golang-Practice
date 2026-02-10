package main

import "fmt"

func main() {
	// var ptr *int        // if a pointer is started without any value then it will store <nil>

	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer to myNumber is: ", ptr)
	fmt.Println("Value of actual pointer to myNumber is: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}
