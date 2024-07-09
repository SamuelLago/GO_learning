package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println("Mesma") //Quebra a linha
	fmt.Println(" linha.")

	x := 3.141516

	// fmt.Println("O valor de x é" + x) Nesse caso da erro

	//Primeira solucao
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	//Segunda solucao
	fmt.Println("O valor de x é", x)

	//Terceira solucao
	fmt.Printf("O valor de x é %.2f", x) // %f para usar com float, .2 é o numero de casas do numero

	a := 1
	b := 1.9999
	c := false
	d := "Opa"
	fmt.Printf("\n%d %f %.2f %t %s", a, b, b, c, d) //%d numero / %f float / %t boolean / %s string
	fmt.Printf("\n%v %v %v %v", a, b, c, d)         //%v é mais generico
}
