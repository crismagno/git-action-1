package main

import "fmt"

func main() {
	a, b := 3, 5
	result := Soma(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, result)
}

func Soma(a, b int) int {
	return a + b
}
