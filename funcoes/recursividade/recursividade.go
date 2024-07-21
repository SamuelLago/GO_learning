package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return 0, fmt.Errorf("Numero invalido")
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1) //Chama os fatoriais ate o return do 0
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(3)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
