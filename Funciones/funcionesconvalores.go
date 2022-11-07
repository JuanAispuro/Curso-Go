package main
import "fmt"

func main() {
	hello("Juan", "Aispuro")
	emoji := "😎"
	change(&emoji)
	fmt.Println(emoji)
}

func hello(firstName string, lastName string){
	fmt.Println("Hello",firstName,lastName)
}
//Recibira un valor y lo cambiara
//Funcion que recibe parametro a base por valor.

	/*func Change(value string){
	value = "😂"
	}
	*/

//funcion que recibe parametro a base por referencia
//Ponemos un puntero y ocuparemos el valor de ampersand (&)
func change(value *string){
	*value = "😂"
}