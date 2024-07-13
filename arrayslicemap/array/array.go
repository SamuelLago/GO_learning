package main

import "fmt"

func main() {
	// Homogenea (mesmo tipo em toda a array) e estatica (fixo, mesmo tamanho)
	var nota [3]float64
	var a = "ola"

	fmt.Println(nota)

	nota[0], nota[1], nota[2] = 7.8, 4.3, 9.1

	fmt.Println(nota)

	total := 0.0

	for i := 0; i < len(nota); i++ { //len == Tamanho da array
		total += nota[i]
	}
	media := total / float64(len(nota))

	fmt.Println(len(a)) //len == Tamanho da String

	fmt.Printf("Media da nota:%.2f", media)
}
