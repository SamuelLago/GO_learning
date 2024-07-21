package main

import f "fmt" // Pode mudar o nome do import

//Fundamentos

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

//Estruturas de Controle

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

func anoBissexto(ano int) bool { //Ex.7
	if ano%4 == 0 && (ano%100 != 0 || ano%400 == 0) {
		return true
	} else {
		return false
	}
}

func numeroPrimo(num1 int) bool { //Ex.8
	for i := 2; i < num1; i++ {
		if num1%i == 0 {
			return false
		}
	}
	return true
}

// Array Slice Map

func mediaNumeros(numeros [5]int) float64 { //Ex.9
	total := 0
	for i := 0; i < len(numeros); i++ {
		total += numeros[i]
	}
	return float64(total) / float64(len(numeros))
}

func contarNumerosPositivos(s1 []int) int { //Ex.10
	total := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] > 0 {
			total++
		}
	}
	return total
}

func buscaArray(array1 [10]int, num int) string { //Ex.11
	for i := 0; i < len(array1); i++ {
		if array1[i] == num {
			return "Encontrado"
		}
	}
	return "Nao encontrado"
}

func inverteArray(array1 [10]int) [10]int { //Ex.12
	var array2 [10]int
	for i := 0; i < len(array1); i++ {
		array2[len(array1)-1-i] = array1[i]
	}
	return array2
}

func main() {
	f.Printf("A temperatura em celsius %2.f°C\n", converteCelsius(91))
	f.Println("O total da area do retangulo:", areaRetangulo(10, 5))
	f.Println(operacoesBasicas(4, 2))
	f.Println(verificarParImpar(3))
	f.Println(verificarIdade(19))
	f.Println(calculadora(3, 6, '+'))
	f.Println(anoBissexto(2020))
	f.Println(numeroPrimo(11))
	f.Println(mediaNumeros([5]int{10, 3, 3, 4, 5}))
	f.Println(contarNumerosPositivos([]int{1, 2, -1, 3, -2, -2, 1}))
	f.Println(buscaArray([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7))
	f.Println(inverteArray([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
