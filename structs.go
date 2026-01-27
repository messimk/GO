package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name:")
	lastName := getUserData("please enter your last name:")
	birthDate := getUserData("Please enter yout birtdate")

	outputUserDetails(firstName, lastName, birthDate)

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
