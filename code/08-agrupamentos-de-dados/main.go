package main

import "fmt"

func main() {
	// dadosArray()
	dadosSlice()
}

// Arrays possuem um tamanho fixo e fazem parte do tipo
var x [5]int // tipo [5]int, tipos diferentes

func dadosArray() {
	x[0] = 10
	x[1] = 20
	// x[2] = começa com o valor zero (zero value) -> 0
	x[4] = 50
	// x[5] = 50 -> invalid argument: index 5 out of bounds [0:5]

	fmt.Printf("%T - %v\n", x, x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%v - %v\n", i, x[i])
	}

	y := [6]int{0, 1, 2, 3, 4, 5}
	for i, v := range y {
		fmt.Printf("%v - %v\n", i, v)
	}
}

// Slice é um array dinâmico -> coleção de tipos
func dadosSlice() {
	slice := []int{0, 1, 2, 3}
	// fmt.Println(append(slice, 4, 5, 6))-> [0 1 2 3 4 5 6]
	// fmt.Println(slice) -> [0 1 2 3]
	slice = append(slice, 4, 5, 6)
	
	// Mesmo que a capacidade seja maior, você não pode acessar além do len.
	// slice[11] = 11 -> panic: runtime error: index out of range [0:11] with length 4

	for i, v := range slice {
		fmt.Printf("%v - %v\n", i, v)
	}
}
