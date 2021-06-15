package main

import "fmt"

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

//Slice arguments vs Variadic arguments

func sliceandvariadic() {
	find2(89, []int{89, 90, 95})
	find2(45, []int{56, 67, 45, 90, 109})
	find2(78, []int{38, 56, 98})
	find2(87, []int{})
}

func find2(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

//Gotcha

func gotcha() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
func change(s ...string) {
	s[0] = "Go"
}
