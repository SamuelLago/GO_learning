package main

import "fmt"

func main() {
	//Escrever o tipo quando o valor nao for atribuido
	var a int
	var b float64
	var c bool
	var d string
	var e *int //Ponteiro de inteiro

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
