package main

import (
	"fmt"
)

const LoginToken string = "grshsytyytyjdtr" // first letter capital indicates that this is the Public variable

func main() {
	var username string = "yash"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username) //%T is used to tell the data type of the variable

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.1235466585464345678765
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and aliases

	var anotherVariable int //default value will be zero not any garbage value in it
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var randomString string //default value will be empty string not any garbage value in it
	fmt.Println(randomString)
	fmt.Printf("Variable is of type: %T \n", randomString)

	// implicite type

	var website = "www.yg.com" //lexer will give it the data type type on its own according to the value
	fmt.Println(website)

	// no var style  (only allowed inside the method)

	numberOfUsers := 1500 // := is called the walarus operator
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
