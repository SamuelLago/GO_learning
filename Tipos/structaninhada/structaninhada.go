package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, quantidade: 4, preco: 12.10},
			item{2, 1, 5.50},
			item{3, 6, 6.0},
		},
	}
	fmt.Printf("Valor total do pedido: R$ %2.f", pedido.valorTotal())
}
