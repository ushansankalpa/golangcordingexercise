package main

import "fmt"

func main() {
	name := "Hello World"
	fmt.Println(name)

	fmt.Printf("String: %s\n", name)
	printBytes(name)

	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)

	string1 := "Go"
	string2 := "Go"
	compareStrings(string1, string2)

	string3 := "hello"
	string4 := "world"
	compareStrings(string3, string4)
}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

//String comparison
func compareStrings(str1 string, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}
