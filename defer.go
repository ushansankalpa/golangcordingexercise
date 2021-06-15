package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)

	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")

	name := "Ushan"
	fmt.Printf("Original String: %s\n", string(name))
	fmt.Printf("Reversed String: ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}

}

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

//Practical use of defer
type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
}
