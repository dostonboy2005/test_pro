package user

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (member *User) Validate(name string,age int,email string) []error {
	var errors []error

	if len(name) < 6 {
		errors = append(errors,fmt.Errorf("length of Name should be 6 or above"))
	}

	str := strconv.Itoa(age)
		if str=="0"{
			errors=append(errors, fmt.Errorf("age: can not be empty"))
		}
	if age < 0 || age> 120{
		errors=append(errors, fmt.Errorf("age: should be between 0 and 120"))
	}
	if email == "" {
		errors = append(errors, fmt.Errorf("email: cannot be empty"))
	} else {
		at := strings.Index(email, "@")
		dot := strings.LastIndex(email, ".")
		if at == -1 || dot == -1 || dot < at {
			errors = append(errors, fmt.Errorf("email: Wrong format"))

		}
	}

	return errors
}