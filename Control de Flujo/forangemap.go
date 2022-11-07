package main

import "fmt"

func main() {
	sports := map[string]string{"basketball": "🏀","soccer:": "⚽️"}

	for key, v := range sports {
		fmt.Println("Key: ",key,"Valor: ", v)
	}
}
