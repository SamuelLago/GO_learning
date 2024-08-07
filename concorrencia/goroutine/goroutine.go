package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Porque nao fala comigo?", 3)
	//fale("Pedro", "Porque voce comeu meu pastel", 1)

	//go fale("Maria", "Ei...", 500)
	//go fale("Joao", "Opa...", 500)

	go fale("Maria", "Entendi!!!", 10)
	fale("Pedro", "Parabens", 3)
}
