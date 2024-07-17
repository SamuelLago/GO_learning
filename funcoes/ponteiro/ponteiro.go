package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	//Desferencia, tem acesso ao valor no qual o ponteiro mostra
	*n++
}

func main() {
	n := 1
	inc1(n) // Por valor
	fmt.Println(n)

	inc2(&n) // Por referencia
	fmt.Println(n)
}
