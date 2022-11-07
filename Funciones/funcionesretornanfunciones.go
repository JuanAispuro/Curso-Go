package main 
import "fmt"

func main () {
	/*
		nums := []int {1,4,70,5,67,90,2}
		//filter devuelva los numeros menores o iguales a 10.
		result := filter(nums,func(num int) bool {
			if num >=10 {
				return true
			}
			return false
		})
		fmt.Println(result)
	*/
	x := hello ("Juan")
	fmt.Println(x("Aispuro"))
}
	func hello(name string) func(string)string {
		return func(text string) string {
			return "Hello " + name + " " + text
		}
	}

	/*
	//slice de enteros, la funcion y un slice de enteros(el filtrado)
	func filter(nums []int, callback func(int) bool) []int {
		result := []int{}
		//for indice valor.
		for _,v := range nums {
			//llamamamos funcion callback con el valor. 
			callback(v)
			//retorna un booleano.
			if callback(v) {
				result = append(result, v)//agregamos el numero
			}
			//sino pues no se hace.
		}
		return result
	}
	*/