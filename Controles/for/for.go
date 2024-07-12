package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 { //Nao tem while
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { // Inicio; Fim; Passos
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d par\n", i)
		} else {
			fmt.Printf("%d impar\n", i)
		}
	}

	for { //Laço infinito
		fmt.Println("Loop infinito")
		time.Sleep(time.Second * 3)
	}
}
