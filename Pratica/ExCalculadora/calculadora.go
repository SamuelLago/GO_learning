package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("Escolha uma operação:")
		fmt.Println("1. Adição")
		fmt.Println("2. Subtração")
		fmt.Println("3. Multiplicação")
		fmt.Println("4. Divisão")
		fmt.Println("5. Módulo")
		fmt.Println("6. Sair")

		var opcao int
		fmt.Scan(&opcao)

		if opcao == 6 {
			fmt.Println("Saindo...")
			break
		}

		var num1, num2 float64
		fmt.Print("Digite o primeiro número: ")
		fmt.Scan(&num1)
		fmt.Print("Digite o segundo número: ")
		fmt.Scan(&num2)

		switch opcao {
		case 1:
			fmt.Printf("Resultado da Adição: %.2f\n", num1+num2)
		case 2:
			fmt.Printf("Resultado da Subtração: %.2f\n", num1-num2)
		case 3:
			fmt.Printf("Resultado da Multiplicação: %.2f\n", num1*num2)
		case 4:
			if num2 != 0 {
				fmt.Printf("Resultado da Divisão: %.2f\n", num1/num2)
			} else {
				fmt.Println("Erro: Divisão por zero não é permitida.")
			}
		case 5:
			if int(num2) != 0 {
				fmt.Printf("Resultado do Módulo: %d\n", int(num1)%int(num2))
			} else {
				fmt.Println("Erro: Módulo por zero não é permitido.")
			}
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
