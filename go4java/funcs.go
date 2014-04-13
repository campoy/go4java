package main

import "fmt"

func adder(c int) func(int) int {
	return func(x int) int {
		return x + c
	}
}

func main() {
	// START OMIT
	addTwo := adder(2)
	for i := 0; i < 10; i++ {
		fmt.Println(addTwo(i))
	}
	// END OMIT
}
