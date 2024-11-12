package main

import (
	"errors"
	"fmt"
	"math"

	"example.com/demo_go/power"
	"github.com/Pallinder/go-randomdata"
)

func pointerChecker() {

	fmt.Println(randomdata.Email())
	var num float64 = 10
	var numPointer *float64

	pointerUser()

	fmt.Println("-------------------")
	numPointer = &num
	fmt.Printf(`this is the value for %v**2 : `, num)
	fmt.Print(power.GetPower(numPointer, 2), "\n")
	changeValueOf(numPointer)
	fmt.Printf(`this is the value for %v**2 : `, num)
	fmt.Print(power.GetPower(numPointer, 2), "\n")
	fmt.Println("-------------------")

	// mainFunc()
	// printEven(120)
	// whileLoop()

}

// this type of return method is recomended
func Calculate(invesment float64, years float64) (string, float64) {
	result := invesment * years
	message := "this is the final value"
	return message, result

}

func Calculateer(invesment float64, years float64) (message string, result float64) {
	result = invesment * years
	message = "this is the final value"
	// return //this ,method of returning is not recomended as it is confusing for debuggimg
	return message, result

}

func getUserInput(message string) (float64, error) {
	var userInput float64
	fmt.Print(message)
	fmt.Scan(&userInput)
	// var input float64;

	if userInput <= 0.0 {
		return 0, errors.New("value must be a positive number")
	}
	return userInput, nil
}

func mainFunc() {
	const message string = "Hello, World!! enter the values"
	fmt.Println(message)
	// var invesment float64 = 10.34
	// var years float64 = 20
	//user input

	invesment, err := getUserInput("Enter Invesment Amount : ")
	if err != nil {
		// fmt.Println(err)
		// return
		panic(err)
	}

	years, err := getUserInput("Enter Years : ")

	if err != nil {
		// fmt.Println(err)
		// return
		panic(err)

	}
	var result float64 = invesment * years
	fmt.Printf("This is the result : %v \n 2**10 = %v \n", result, math.Pow(2, 10))

	fmt.Println(Calculate(invesment, years))
	fmt.Println(Calculateer(invesment, years))
}

// func print

func printEven(num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}

func whileLoop() {
	var status bool = true
	var count int = 0
	for status {
		if count >= 10 {
			status = false
			return
		}
		count += 1
		fmt.Println(count)

	}
}

func pointerUser() {
	var age float64 = 34

	var agePointer *float64
	agePointer = &age

	fmt.Println(agePointer)
	fmt.Println(*agePointer)

	fmt.Println(*agePointer)

	fmt.Println(power.GetPower(agePointer, 3))

}

func changeValueOf(num *float64) {
	*num += 1
}
