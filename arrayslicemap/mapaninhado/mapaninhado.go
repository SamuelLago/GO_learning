package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"J": {
			"Joao Vitor":      1412.76,
			"Joaquim Pereira": 6324.52,
		},
		"R": {
			"Rafael Monteiro": 123.45,
		},
		"P": {
			"Pedro Silva": 123.45,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}
}
