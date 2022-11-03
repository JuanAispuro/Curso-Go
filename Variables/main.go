package main
//fmt es un paquete que nos permite imprimir en la consola
import "fmt"
func main() {
    fmt.Println(4 != 5)
    var age = 30
    fmt.Println("adulto: ",  age >= 18 && age <= 60)
    //El otro operador es inverso de que si la edad es para un niño o tercera edad
    fmt.Println("niño o viejo: ",  age < 18 || age > 60)
    //El unitario inverte el valor booleano
    fmt.Println(!(4 == 4 ))
}