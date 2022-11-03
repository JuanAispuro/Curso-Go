package main
import "fmt"
func main() {

	set:= [7]string{"Leon","Caballo","Vaca","Mariposa","Pajaro","Avioneta","Avion"}
	//Slice de animales
	//Indice inicial e indice final, leon hasta pajaro
	animals := set[0:5]
	//Toma desde la posición 0 hasta la 5
	fly := set[3:7]
	fly[0] = "Dragon"
	//Los slice no poseen datos, son apuntadores de un array, por lo que si cambia
	// el slice de mariposa al slice de dragon.
	fmt.Println("Array: ",set)
	fmt.Println("Animales: ", animals)
	fmt.Println("Voladores: ", fly[0])

}
