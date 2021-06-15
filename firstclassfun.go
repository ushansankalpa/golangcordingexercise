package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")

	var b add = func(a int, b int) int {
		return a + b
	}
	s := b(5, 6)
	fmt.Println("Sum", s)

	a1 := appendStr()
	b1 := appendStr()
	fmt.Println(a1("World"))
	fmt.Println(b1("Everyone"))

	fmt.Println(a1("Gopher"))
	fmt.Println(b1("!"))

	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s0 := []student{s1, s2}
	f := filter(s0, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
}

type add func(a int, b int) int

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

//Practical use of first class functions
type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}
