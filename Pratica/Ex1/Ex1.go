package main

import f "fmt" // Pode mudar o nome do import

func converteCelsius(temperaturaF float64) float64 { //Ex.1
	return (temperaturaF - 32) / 1.8
}

func main() {
	f.Printf("A temperatura em celsius %2.f°C\n", converteCelsius(91))
}
