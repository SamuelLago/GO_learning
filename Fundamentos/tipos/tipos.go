package main

import (
	"fmt"
	m "math"
	"reflect"
)

func main() {
	//Numeros inteiros
	fmt.Println(1, 2, 50)
	fmt.Println("Literal inteiro é", reflect.TypeOf(50000))

	//Sem sinal (positivos) / uint8 / uint16 / uint32 / uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//Com sinal / int8 / int16 / int32 / int64
	i1 := m.MaxInt64
	fmt.Println("O maximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //int32
	fmt.Println("O rune é", reflect.TypeOf(i2))

	//Numeros reais / float32 / float64
	var x float32 = 49.99 // Se quiser usar float32, declarar antes
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//String
	s1 := "Ola eu sou Samuel"
	fmt.Println(s1 + "!")
	fmt.Println(s1)
	fmt.Println("O tamanho da string é", len(s1))

	//String com multiplas linhas
	s2 := `Ola 
	eu
	sou
	Samuel` //O \n é ignorado
	fmt.Println(s2)
	fmt.Println("O tamanho da string é", len(s2))
}
