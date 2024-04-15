package main

import (
	"fmt"

	"user/validate"
)



func main() {
	var name string
	var age int
	var email string

	fmt.Print("Enter Name:")
	fmt.Scanln(&name)

	fmt.Print("Enter Age:")
	fmt.Scanln(&age)

	fmt.Print("Enter Email:")
	fmt.Scanln(&email)

	user := user{Name: name, Age: age, Email: email}
	errors := user.Validate(name,age,email)

	if len(errors) > 0 {
		fmt.Println("\nErrors:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("successfully.")
	}
}
