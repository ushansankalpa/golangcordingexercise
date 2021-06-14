package main

import (
	"fmt"
	"math"
)

func main() {
	//variable declaration
	var age int
	fmt.Println("My age is ", age)

	age = 22
	fmt.Println("My age is ", age)

	age = 54
	fmt.Println("My age is ", age)

	withInitial()

	withoutType()

	multipleVar()
	multipleVar2()
	multipleVar3()

	shortHandDec()
	shortHandDec2()
	shortHandDec3()
	shortHandDec4()

}

//variable declaration with initial value
func withInitial() {
	var age int = 22
	fmt.Println("My age is ", age)

}

//type will be inferred
func withoutType() {

	var age = 22
	fmt.Println("My age is ", age)

}

//declaring multiple variable
func multipleVar() {

	var width, height int = 100, 50
	fmt.Println("Width is ", width, "Height is ", height)

	var wi, hei = 100, 50
	fmt.Println("Width is ", wi, "Height is ", hei)
}

func multipleVar2() {
	var width, height int
	fmt.Println("Width is ", width, "Height is ", height)

	width = 200
	height = 100
	fmt.Println("New Width is ", width, " New Height is ", height)
}

func multipleVar3() {
	var (
		name   = "Ushan"
		age    = 21
		height int
	)

	fmt.Println("My name is ", name, ", age is ", age, ", and height is ", height)
}

func shortHandDec() {
	count := 10
	fmt.Println("Count ", count)
}

func shortHandDec2() {
	name, age := "Ushan", 29
	fmt.Print("My name is ", name, "My age is ", age)
}

func shortHandDec3() {
	a, b := 20, 30
	fmt.Println("a is ", a, "b is ", b)

	b, c := 40, 50
	fmt.Println("b is ", b, "c is ", c)

	b, c = 80, 90
	fmt.Println("New a is ", b, "New b is ", c)

}

func shortHandDec4() {
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Minimum value is ", c)
}
