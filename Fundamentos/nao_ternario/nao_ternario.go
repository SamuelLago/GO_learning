package main

import "fmt"

func obterResultado(nota float64) string {
	if nota <= 6{
		return "Aprovado"
	}
	return "Reprovado"
}

func main(){
	fmt.Printf(obterResultado(7.5))
}