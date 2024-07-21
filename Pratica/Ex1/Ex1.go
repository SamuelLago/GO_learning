package main

import f "fmt" // Pode mudar o nome do import

func converteCelsius(temperaturaF float64) float64 { //Ex.1
	return (temperaturaF - 32) / 1.8
}
func areaRetangulo(largura, altura float64) float64 { //Ex.2
	return largura * altura
}

func operacoesBasicas(num1, num2 float64) string { //Ex.3
	soma := num1 + num2
	sub := num1 - num2
	div := num1 / num2
	mult := num1 * num2

	return f.Sprintf("Soma:%2.f Subtracao:%2.f Divisao:%2.f Multiplicacao:%2.f", soma, sub, div, mult)
}

func verificarParImpar(num1 int) string { //Ex.4
	if num1%2 == 0 {
		return "par"
	} else {
		return "impar"
	}
}

func main() {
	f.Printf("A temperatura em celsius %2.f°C\n", converteCelsius(91))
	f.Println("O total da area do retangulo:", areaRetangulo(10, 5))
	f.Println(operacoesBasicas(4, 2))
	f.Println(verificarParImpar(3))
}
