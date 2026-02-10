package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome onboard on the Golang journey"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age:")

	// comma ok // err

	input, _ := reader.ReadString('\n') // use err instead of _ if you want to store the err at the time of input
	fmt.Println("Thanks for entering your age, ", input)
	fmt.Printf("Type of the age is: %T21", input)

}
