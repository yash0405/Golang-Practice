package main

import "fmt"

func main() {
	// No inheritance in golang: No super or parent

	yash := User{"Yash", "yg@gmail.com", true, 21}
	fmt.Println(yash)

	fmt.Printf("Details of Yash are: %+v\n", yash)

	fmt.Printf("Name is %v and email is %v", yash.Name, yash.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
