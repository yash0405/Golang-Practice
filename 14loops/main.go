package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ { //pre increment (++d) is not allowed in Golang
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and the value is %v\n", index, day)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

}
