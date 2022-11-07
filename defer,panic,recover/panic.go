package main

import "fmt"

func main() {
	division(10,2)
	division(40,3)
	division(12,0)
	division(20,4)
	//Se ejecutan las primeras 2 pero el panic se efectua en la tercera
}

func division(diviendo, divisor int) {
	validarDivisor(divisor)
	fmt.Println(diviendo / divisor)
}

func validarDivisor(divisor int){
	if divisor == 0 {
		panic("Panic")
	}
}
