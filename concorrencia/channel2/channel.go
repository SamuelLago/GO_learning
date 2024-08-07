package main

import (
	"fmt"
	"time"
)

//Channel - Forma de comunicacao entre goroutines
//Channel - um tipo, como int float32 etc...

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // Enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // Recebendo dados do canal
	fmt.Println(a, b)

	fmt.Println(<-c)
}
