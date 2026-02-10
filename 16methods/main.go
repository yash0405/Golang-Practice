package main

import "fmt"

func main() {
	// No inheritance in golang: No super or parent

	yash := User{"Yash", "yg@gmail.com", true, 21}
	fmt.Println(yash)

	fmt.Printf("Details of Yash are: %+v\n", yash)

	fmt.Printf("Name is %v and email is %v\n", yash.Name, yash.Email)

	yash.GetStatus()
	yash.NewMAil()    // 

	fmt.Printf("Name is %v and email is %v\n", yash.Name, yash.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMAil() {                // pass by copy only chnges the email within the function scope not the original value outside it
	u.Email = "test@gmail.com"
	fmt.Println("Email.of this user is: ", u.Email)
}


