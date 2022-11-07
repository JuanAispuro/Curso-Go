package main

import "fmt"

func main() {
	total := sum(4, 7)
	fmt.Println(total)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}