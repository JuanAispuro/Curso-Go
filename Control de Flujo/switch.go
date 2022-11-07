package main

import "fmt"

func main() {
	/* Estructura del switch
	switch expression {
	case condition:
	case condition:
	default:
	}
	*/
	emoji := "🍎"
	switch  {

	case emoji == "🐱" || emoji == "🐶":
		fmt.Println("Es un gato")
	
	case emoji == "🍎" || emoji == "🍌":
		fmt.Println("Es una fruta")

	default:
		fmt.Println("No es un animal ni una fruta")
	}
}
