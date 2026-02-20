package main

import "fmt"

// Arrays possuem um tamanho fixo
var x [5]int
var y [6]int

func main() {
	x[0] = 10
	x[1] = 20
	// x[2] = 30
	x[3] = 40
	x[4] = 50

	fmt.Printf("%T - %v\n", x, x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%v - %v\n", i, x[i])
	}
}
