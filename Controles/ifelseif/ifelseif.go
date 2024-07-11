package main

func notaParaConceito(n float64) string {
	if n <= 10 && n > 9 {
		return "A"
	} else if n <= 9 && n > 7 {
		return "B"
	} else if n <= 7 && n > 5 {
		return "C"
	} else if n <= 5 && n > 3 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	println(notaParaConceito(9.1))
	println(notaParaConceito(5.1))
	println(notaParaConceito(1.1))
}
