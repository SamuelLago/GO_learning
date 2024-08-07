package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //Tambem suportado pelo switch
		fmt.Println(i)
		fmt.Println("Ganhou!!!")

	} else {
		fmt.Println(i)
		fmt.Println("Perdeu!")
	}
}
