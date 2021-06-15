package main

import "fmt"

func main() {
	employeeSalary := make(map[string]int)
	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary)

	employeeSalary2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	employeeSalary2["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary2)

	//Retrieving value for a key from a map
	retrievingvalue()

	iteratingover()

	mapofstructs()

	lengthofmap()
}

//Retrieving value for a key from a map
func retrievingvalue() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	employee := "jamie"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)
}

//Iterate over all elements in a map
func iteratingover() {
	employeeSalary := map[string]int{
		"steve": 1200,
		"jamie": 1500,
		"mike":  3000,
	}

	fmt.Println("Contents of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}
}

//Deleting items from a map
func deletingitem() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)
}

//Map of structs

type employee struct {
	salary  int
	country string
}

func mapofstructs() {
	emp1 := employee{
		salary:  1200,
		country: "srilanka",
	}

	emp2 := employee{
		salary:  14000,
		country: "india",
	}
	emp3 := employee{
		salary:  15000,
		country: "canada",
	}

	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Jamie": emp2,
		"Mike":  emp3,
	}

	for name, infor := range employeeInfo {
		fmt.Printf("Employee: %s salary:%d country:%s\n", name, infor.salary, infor.country)
	}
}

//length of map
func lengthofmap() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	fmt.Println("length is", len(employeeSalary))
}
