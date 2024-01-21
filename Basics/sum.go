package main

import (
	"fmt"
)

func sum(a float64, b float64) float64 {
	c := a + b
	return c
}

func main() {
	fmt.Print("Enter the first number: ")
	var num1 float64
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	var num2 float64
	fmt.Scan(&num2)

	num3 := sum(num1, num2)

	fmt.Printf("Sum of %.2f and %.2f is %.2f\n", num1, num2, num3)
}
