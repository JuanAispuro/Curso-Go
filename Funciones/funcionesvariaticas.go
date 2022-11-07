package main

import "fmt"

func main() {
	fmt.Println(sum(1,2,6,3,6))
}

// Con esto se puede pasar un numero indefinido de parametros (funcion variatica).
func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}
	return total
}