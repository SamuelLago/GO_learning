package main

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	case n == 1:
		return 1
	default:
		return n * fatorial(n-1)

	}
}

func main() {
	println(fatorial(5))
}
