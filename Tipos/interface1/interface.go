package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type product struct {
	nome  string
	price float64
}

func (p pessoa) toString() string {
	return fmt.Sprintf("%s %s", p.nome, p.sobrenome)
}

func (p product) toString() string {
	return fmt.Sprintf("%s %.2f", p.nome, p.price)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	//Cria uma variavel chamada coisa e usando polimorfismo
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	coisa = product{"Meia", 12.00}
	fmt.Println(coisa.toString())

	//usa o imprimir direto
	imprimir(product{"Camisa", 54.000})
	imprimir(pessoa{"João", "Silva"})

}
