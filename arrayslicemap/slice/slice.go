package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3] // Comeca no primeiro numero e termina 1 antes do segundo numero, alem disso nao cria um novo array, mas define uma parte dele
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(a2, s3)

	s4 := s2[:1] // slice de um slice
	fmt.Println(s2, s4)
}
