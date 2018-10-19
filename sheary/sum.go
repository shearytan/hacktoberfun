package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	var num1 int
	var num2 int
	fmt.Print("Enter the 1st number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the 2nd number: ")
	fmt.Scan(&num2)
	fmt.Println("The sum of these 2 numbers is: ", sum(num1, num2))
}
