package main

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	defer fmt.Println("deferred call in main")
	firstName2 := "Elon"
	fullName2(&firstName2, nil)
	fmt.Println("returned normally from main")
}

//Recovering from a Panic

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName2(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}
