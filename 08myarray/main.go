package main

import "fmt"

func main() {

	var fruits [4]string

	fruits[0] = "apple"
	fruits[1] = "papaya"
	// fruits[2] = "grapes"
	fruits[3] = "mango"

	fmt.Println("Fruits List is: ", fruits)

	var vegies = [3]string{"potato", "beans", "carrot"}
	fmt.Println("Vegy List is: ", vegies)
}
