package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste " + string(97)) //O numero representa pela forma Unicode / ASCII

	//int para String
	fmt.Println("Teste " + strconv.Itoa(123))

	//String para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1") // 0 é False, 1 é True
	if b {
		fmt.Println("True")
	}
}
