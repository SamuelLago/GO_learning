package main

import "math"

// Iniciando com a letra maiuscula e PUBLICO (visivel dentro e fora do pacote)
// Iniciando com a letra minuscula e PRIVADO (visivel apenas dentro do pacote)

// Ponto - Publico
// ponto ou _Ponto - Privado

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
