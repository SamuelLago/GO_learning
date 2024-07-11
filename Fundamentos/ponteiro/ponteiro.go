package main

import "fmt"

func main() {
	i := 1

	var p *int = nil //Ponteiro (nil == null)

	p = &i // Pega o endereco da variavel (onde esta alocado na memoria)

	*p++ //Desrefenciando (pegando o valor)
	i++

	// GO nao tem aritmetica de  ponteiros

	fmt.Println(p, *p, i, &i) //0xc000096068 3 3 0xc000096068
}
