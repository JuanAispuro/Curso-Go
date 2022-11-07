package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "JuAn ArMANdo"
	miusc,mayusc := convert(texto)
	fmt.Println(miusc,mayusc)
}

func convert(text string) (string, string) {
	//Primero mayusculas luego minusculas
	min := strings.ToLower(text)
	may := strings.ToUpper(text)
	return min, may
}
