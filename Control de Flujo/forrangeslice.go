package main

import "fmt"

func main() {
	nums := []uint8{2,4,6,8}
	for i :=range nums{
		//iterar el slice y se ponga la palabra edteam por cada numero
		nums[i] *= 2
	}	
	fmt.Println(nums)
	
}