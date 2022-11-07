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
	//siempre se ejecute cuando se imprima o haya un panic
	defer func (){
		if r:= recover(); r != nil {
			//osea que esto pasara cuando haya un Panic
			fmt.Println("Recuperandose del panic", r)
			//controlamos el error y no se rompe el programa
		}
	}()
	validarDivisor(divisor)
	fmt.Println(diviendo / divisor)
}

func validarDivisor(divisor int){
	if divisor == 0 {
		panic("Panic")
	}
}
