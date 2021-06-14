package main

import (
	"fmt"
)

func main() {

	ifandelse()
	elseif()
	ifwithAssignment()
	idiomatic()

	num := 10
	if num%2 == 0 {
		fmt.Println("The number ", num, "is even")
		return
	}

	fmt.Println("The number is ", num, "is odd")

}

func ifandelse() {
	num := 10
	if num%2 == 0 {
		fmt.Println("the number is ", num, "is even'")
	} else {
		fmt.Println("the number is ", num, "is odd'")
	}
}

func elseif() {
	num := 99
	if num <= 50 {
		fmt.Println(num, "is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println(num, "is between 51 and 100")
	} else {
		fmt.Println(num, "is greater than 100")
	}
}

func ifwithAssignment() {
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

func idiomatic() {
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
		return
	}
	fmt.Println(num, "is odd")
}
