package main

import "fmt"

func recursiveDivide(n int) {
	if n == 1 {
		return
	}
	recursiveDivide(n / 2)
	fmt.Println(n)
}

func main() {
	input1 := 9
	input2 := 17
	input3 := 31

	fmt.Println("input:1")
	recursiveDivide(input1)

	fmt.Println("input:2")
	recursiveDivide(input2)

	fmt.Println("input:3")
	recursiveDivide(input3)
}
