package main

import (
	"fmt"
	"reflect"
)

func main() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(createQuery(o))

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Science Park Road, Singapore",
		salary:  90000,
		country: "Singapore",
	}
	fmt.Println(createQuery1(e))

	createQuery2(o)

	createQuery3(o)
	createQuery3(e)
	i := 90
	createQuery3(i)

}

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery1(o employee) string {
	i := fmt.Sprintf("insert into order values(%s %d %s %d %s)", o.name, o.id, o.address, o.salary, o.country)
	return i
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func createQuery2(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}

func createQuery3(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}
