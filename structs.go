package main

import (
	"fmt"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
}

func main() {
	firstName := getUserData("Please enter your first name:")
	lastName := getUserData("please enter your last name:")
	birthDate := getUserData("Please enter yout birtdate")

	fmt.Println(firstName, lastName, birthDate)

}

func outputUserDetails(lastName, firstName, birthDate string) {

	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value

}
