package main

import (
	"fmt"
	"math"
)

func main() {

	emp1 := Employee{
		name:     "Ushan sankalpa",
		salary:   5000,
		currency: "lkr",
	}

	emp1.displaySalary()
	displaySalary(emp1)

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area())

	//Pointer Receivers vs Value Receivers
	e := Employee2{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)

}

type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func displaySalary(e Employee) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//Pointer Receivers vs Value Receivers
type Employee2 struct {
	name string
	age  int
}

//Method with value receiver

func (e Employee2) changeName(newName string) {
	e.name = newName
}

//Method with pointer receiver

func (e *Employee2) changeAge(newAge int) {
	e.age = newAge
}
