package main

import "fmt"

func main() {
	type course struct {
		Name      string
		Profesor string
		Country   string
	}
	//Forma 1
	db := course{
		Name:      "Base de datos",
		Profesor: "Juan",
		Country:   "Mexico",
	}
	//Forma 2
	 /*git := course{"GIT","Armando","Bolivia"}
	css := course{Profesor: "Pedro"} //Con esto nomas enviamos el nombre.
	fmt.Printf("%+v\n",git)
	fmt.Printf("%+v\n",css)
	*/	
	//%+v Para imprimir el nombre de los campos.
	p := &db //almacenara un puntero en la base de datos,
	//Operador de desreferencia se usa entre parentesis y con el . el cambio a actualizar.
	p.Profesor = "Juan"
	//No es necesario desreferenciar el punto simplemente se pone el lugar con punto.
	fmt.Printf("%+v\n",db)
	fmt.Printf("%+v\n",p)
	
}
