package main

import "fmt"

func main() {

	emp1 := Employee{
		firstName: "ushan",
		age:       22,
		salary:    2000,
		lastName:  "sankalpa",
	}

	emp2 := Employee{"ushan", "sankalpa", 22, 2000}
	fmt.Println("Employee 1 :", emp1)
	fmt.Println("Employee 2 :", emp2)

	//Creating anonymous structs
	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	//Accessing individual fields of a struct

	emp6 := Employee{
		firstName: "Ushan",
		lastName:  "Adhikari",
		age:       55,
		salary:    6000,
	}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)
	emp6.salary = 6500
	fmt.Printf("New Salary: $%d", emp6.salary)

	//Anonymous fields
	p1 := Person{
		string: "naveen",
		int:    50,
	}
	fmt.Println(p1.string)
	fmt.Println(p1.int)

	//Nested structs
	p2 := Person2{
		name: "Naveen",
		age:  50,
		address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}

	fmt.Println("Name:", p2.name)
	fmt.Println("Age:", p2.age)
	fmt.Println("City:", p2.address.city)
	fmt.Println("State:", p2.address.state)

	//Structs Equality
	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name4 := name{
		firstName: "Steve",
	}

	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

//Anonymous fields
type Person struct {
	string
	int
}

//Nested structs
type Address struct {
	city  string
	state string
}

type Person2 struct {
	name    string
	age     int
	address Address
}

//Structs Equality
type name struct {
	firstName string
	lastName  string
}
