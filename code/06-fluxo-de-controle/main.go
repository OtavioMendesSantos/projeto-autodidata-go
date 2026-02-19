package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for x := 0; x < 10; x++ {fmt.Println(x)}; -> for loop can be modernized using range over int
	// Go não tem while

	total := 0
	for i := range 100 {
		for x := range 10 {
			fmt.Printf("Primeiro loop %v, Segundo loop %v\n", i, x)
			total++
		}
	}
	fmt.Printf("===== Finalizou com %v repetições =====\n", total)

	// Infinite loop
	// for {
	// 	fmt.Println("Executando...")
	// }

	x := 0
	y := 0

	fmt.Println("Iniciando processamento desnecessário...")
	for {
		if x >= 10 {
			break
		}
		random := rand.Intn(51) + 1
		y++
		if random > 50 {
			fmt.Printf("Entrando +1 para a conta (%v na tentativa %v)\n", x, y)
			x++
		}
	}
	fmt.Printf("===== Após %v tentativas, o for acabou =====\n", y)
}
