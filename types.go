package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	b := false

	fmt.Println("a :", a, "b :", b)

	c := a && b
	fmt.Println("c :", c)

	d := a || b
	fmt.Println("d :", d)

	ranges()

	floatingPoint()

	complexType()
	numericType()
	TypeConv()
}

func ranges() {
	var a int = 89
	b := 95

	fmt.Println("Value of a is", a, "Value of b is ", b)
	fmt.Printf("type of a is %T , size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T , size of b is %d", a, unsafe.Sizeof(b))

}

func floatingPoint() {
	a, b := 5.67, 8.97
	fmt.Printf("Type of a %T b %T\n", a, b)

	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}

func complexType() {
	c1 := complex(5, 7)
	c2 := 8 + 27i

	cadd := c1 + c2
	fmt.Println("sum ", cadd)

	cmul := c1 * c2
	fmt.Println("product : ", cmul)
}

func numericType() {
	first := "Ushan "
	last := "Adhikari"
	name := first + " " + last
	fmt.Println("My name is ", name)
}

func TypeConv() {
	i := 65
	j := 67.8
	sum := i + int(j)
	fmt.Println(sum)

	k := 10
	var l float64 = float64(k)
	fmt.Println("l", l)
}
