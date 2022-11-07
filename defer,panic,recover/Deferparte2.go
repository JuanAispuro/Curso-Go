package main 

import ("os"
"fmt")

func main(){
	//creamos un archivo hello.txt y nos regresa un puntero de error o file.
	file, err := os.Create("Hello.txt")
	if err != nil {
		fmt.Printf("Ocurrio un error al crear el archivo: %v", err)
		return
	}
	defer file.Close()
	//Escribir en el archivo
	//Le reasignamos un valor a error 
	_,err = file.Write([]byte ("Hello There!"))
	if err != nil {
		//si no agarra cerramos el archivo ya no sera necesario por el defer
		// file.Close()
		fmt.Printf("Ocurrio un error escribir el archivo: %v", err)
		return
	}
}