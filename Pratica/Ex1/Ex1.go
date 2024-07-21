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

func verificarIdade(idade int) string { //Ex.5
	if idade < 0 {
		return "Idade invalida"
	} else {
		if idade > 0 && idade < 12 {
			return "Criança"
		} else if idade >= 12 && idade < 18 {
			return "Adolescente"
		} else if idade >= 18 && idade < 64 {
			return "Adulto"
		} else {
			return "Idoso"
		}
	}
}

func calculadora(num1, num2 float64, caracter byte) float64 { //Ex.6
	switch caracter {
	case '+': //case aceita caracter
		return num1 + num2
	case '-': //case aceita caracter
		return num1 - num2
	case '*': //case aceita caracter
		return num1 * num2
	case '/': //case aceita caracter
		return num1 / num2
	default:
		return 0
	}
}

func anoBissexto(ano int) bool {
	if ano%4 == 0 && (ano%100 != 0 || ano%400 == 0) {
		return true
	} else {
		return false
	}
}

func main() {
	f.Printf("A temperatura em celsius %2.f°C\n", converteCelsius(91))
	f.Println("O total da area do retangulo:", areaRetangulo(10, 5))
	f.Println(operacoesBasicas(4, 2))
	f.Println(verificarParImpar(3))
	f.Println(verificarIdade(19))
	f.Println(calculadora(3, 6, '+'))
	f.Println(anoBissexto(2020))
}
