package main

import "fmt"

func main(){
	/*
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	*/
	//Agrega las defer hasta la pila de ejecución y la envia a la parte superior de la pila
	//Por eso ejecuta desde la ultima hasta la primera.
	a := 5
	defer fmt.Println("Defer: ", a)

	a = 10
	fmt.Println(a)
}
