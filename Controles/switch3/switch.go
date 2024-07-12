package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float32, float64:
		return "float"
	case func():
		return "func"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(tipo(5))
	fmt.Println(tipo("Samuel"))
	fmt.Println(tipo(7.1))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
