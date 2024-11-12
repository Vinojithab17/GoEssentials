package main

import (
	"fmt"

	"example.com/demo_go/user"
)

func mainDec() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// appUser = &user.User{
	// 	FirstName: "Max",
	// 	LastName:  "braint",
	// 	BirthDate: "12/23/4143",
	// }
	adminUser := user.NewAdmin()

	adminUser.PrintUserFullname()
	// ... do something awesome with that gathered data!
	appUser.PrintUserFullname()
	appUser.ClearUserName()
	appUser.PrintUserFullname()

	fmt.Println(adminUser.GetEmail_Mac())
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
