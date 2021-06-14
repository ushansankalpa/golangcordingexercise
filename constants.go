package main

import (
	"fmt"
)

func main() {
	const a = 50
	fmt.Println(a)

	const (
		name    = "Ushan"
		age     = 21
		country = "Srilanka"
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	stringConst()
	booleanConst()
	numericConst()
	numericExper()

}

func stringConst() {
	const m = "ushan"
	var name = m

	fmt.Printf("type %T value %v", name, name)
}

func booleanConst() {
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed

	fmt.Println(defaultBool, customBool)
	//defaultBool = customBool //not allowed
}

func numericConst() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type is %T, f's type is %T, c's type is %T", i, f, c)

	const b = 5
	var intVar2 int = b
	var int32Var2 int32 = b
	var float64Var2 float64 = b
	var complex64Var2 complex64 = b
	fmt.Println("intVar", intVar2, "\nint32Var", int32Var2, "\nfloat64Var", float64Var2, "\ncomplex64Var", complex64Var2)
}

func numericExper() {
	var a = 5.9 / 8
	fmt.Printf("a's type is %T and value is %v ", a, a)
}
