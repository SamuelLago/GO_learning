package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[54363732576] = "Maria"
	aprovados[79134723656] = "João"
	aprovados[02346572543] = "Pedro"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados { // cpf == range, nome == aprovados
		fmt.Printf("%v (CPF: %v)\n", nome, cpf)
	}

	fmt.Println(aprovados[54363732576])
	delete(aprovados, 54363732576)
	fmt.Println(aprovados[54363732576])

}
