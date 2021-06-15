package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	b := [3]int{12, 78, 50}
	fmt.Println(b)

	c := [...]int{12, 78, 50}
	fmt.Println(c)

	arrayValueType()

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("Befor passing to function ", num)
	changeLocal(num)
	fmt.Println("after passing to function", num)

	//length of array
	s := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(s))

	//iterating array
	iteratingArray()

	//iteraing arrays2
	iteratingArray2()

	//multidimensional array
	multidimensionalArray()

	//creating slices
	creatingslices()

	//modifying slice
	modifyingslice()

	//modifying slice2
	modifyingslice2()

	//length and capacity of slices
	lengthandcap()
	lengthandcap2()

	//appending to slice
	appendingtoslice()

	//binding two array
	bindingArrays()

	//multidiemensional slices
	multidiemSlice()
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("Inside function", num)
}

func arrayValueType() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a
	b[0] = "singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

}

func iteratingArray() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
}

func iteratingArray2() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}

//multidimensional Array
func multidimensionalArray() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printarray(a)

	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
		fmt.Printf("\n")
	}
}

//creating slices
func creatingslices() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)

	//creates and array and returns a slice reference
	c := []int{6, 7, 8}
	fmt.Println(c)
}

//modifying slice
func modifyingslice() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}

//modifying slice2
func modifyingslice2() {
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
}

//length and capacity of slices
func lengthandcap() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	//length of fruitslice is 2 and capacity is 6
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))
}

func lengthandcap2() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
	fruitslice = fruitslice[:cap(fruitslice)]
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
}

//appending to slice
func appendingtoslice() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
}

//binding two arrays
func bindingArrays() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}

//multidiemensional slices
func multidiemSlice() {
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
