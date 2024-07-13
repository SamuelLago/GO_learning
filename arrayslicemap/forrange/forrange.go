package main

import "fmt"

func main() {
	numeros := [...]int{7, 2, 6, 4, 5, 8} //Compilador conta!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
