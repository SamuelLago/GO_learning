package main

import (
	"time"
)

func main() {
	t := time.Now()
	switch { //switch true
	case t.Hour() < 12:
		println("Good Morning")
	case t.Hour() < 17:
		println("Good Afternoon")
	case t.Hour() < 20:
		println("Good Evening")
	default:
		println("Good Night")
	}
}
