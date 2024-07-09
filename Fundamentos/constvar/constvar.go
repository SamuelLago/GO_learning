package main

import (
	"fmt"
	m "math" //Pode renomear o import
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * m.Pow(raio, 2) //:= forma simploficada

	fmt.Println("A area da circuferencia é", area)

	const (
		a = 1
		b = 2
	)
	//Sempre tem que utilizar as variaveis
	var (
		c = 3
		d = a + b
	)

	fmt.Println("A variavel `C` vale", c, ",soma do a e b é", d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!" //Atribui os valores de acordo com a ordem
	fmt.Println(g, h, i)
}
