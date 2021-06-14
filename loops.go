package main

import "fmt"

func main() {

	secondLoop()
	fmt.Println("\n")
	thirdLoop()
	fmt.Println("\n")
	starpattern()
	fmt.Println("\n")
	lables()
	fmt.Println()
	lable2()
	fmt.Println()
	lable3()

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d", i)
	}

}

func secondLoop() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
}

func thirdLoop() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

func starpattern() {
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func lables() {
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
		}

	}
}

func lable2() {
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}

	}
}

func lable3() {
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}

	}
}

func lable4() {
	i := 0
	// initialisation and post are omitted
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}
}

func lable5() {
	//multiple initialisation and increment
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
