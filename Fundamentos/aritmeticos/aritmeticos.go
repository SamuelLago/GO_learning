package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("subtração =>", a-b)
	fmt.Println("multiplicação =>", a*b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Módulo =>", a%b)

	//bitwise
	fmt.Println("E =>", a&b)   //11 & 10 == 10
	fmt.Println("OU =>", a|b)  // 11 | 10 == 11
	fmt.Println("XOR =>", a^b) // 11 ^ 10 == 01

	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b))) //Maior numero
	fmt.Println("Menor =>", math.Min(c, d))                   //Menor numero
	fmt.Println("Exponenciação =>", math.Pow(c, d))           //(a(numero), b(potencia))
}
