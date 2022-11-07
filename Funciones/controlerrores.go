package main

import (
	"fmt"
	"errors"
	//"strings"
)

func main() {
	
	/*	llamar paquete ioutil con leer archivo.
		Devuelve el contenido que lee o el error.
		content, err := ioutil.ReadFile("archivo.txt")
		La función no lee por lo que se imprime ocurrio el error y el contenido
		del error.

			content, err := ioutil.ReadFile("archivo.txt")
			if err != nil {
				fmt.Printf("Ocurrio un error: %v", err)
				return Para que finalice el programa, por lo que no sera necesario eusar el else
			}
			fmt.Println("El contenido es: %v", string(content))
	*/
	//resultado y error.
	result, err := division(10,0)
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return 
	}
	fmt.Println(result)
}

//Propia funcion para que retorne un error.
func division (dividendo, divisor int) (result int, err error) {
	if divisor == 0{
		err = errors.New("No puedes dividir por cero.")
		return
	}
	result = dividendo / divisor
	return 
}