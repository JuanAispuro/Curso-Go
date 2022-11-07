package main

import "fmt"

func main() {
	hello := "Hello"
	//Devuelve el indice y el valor, pero queremos que imprema los caracteres	
	//Imprime numeros por que son los bytes asignados a esas letras
	//Se utiliza el casting encerrarlo en un string
	for _,v := range hello {
		fmt.Println(string(v))
	}
}