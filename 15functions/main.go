package main

import "fmt"

func main() {
	fmt.Println("Functions Class")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes := proAdder(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Result of proAdder is: ", proRes)
}

// Functions in golag can return multiple datatypes
func adder(a int, b int) int { // int outside the brackets is denoting the return type of function
	return a + b
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Hi I'm Yash Gupta")
}
