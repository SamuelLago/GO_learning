package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	slice1 = append(slice1, 4, 5, 6) // adiciona no slice
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) //2 ==> tamanho do slice
	copy(slice2, slice1)
	fmt.Println(slice2)
}
