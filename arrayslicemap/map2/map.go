package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"João":  1000.31,
		"Maria": 1500.34,
		"Pedro": 800.79,
	}

	funcsESalarios["Samuel"] = 1624.67
	delete(funcsESalarios, "Inexistente") // Mesmo nao existindo nao tem problema

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
