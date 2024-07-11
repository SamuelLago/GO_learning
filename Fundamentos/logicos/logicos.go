package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2    //Ou exclusivo
	comprarSorvete := trab1 || trab2 //Ou inclusivo

	return comprarTV50, comprarTV32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}
