package main

import "fmt"

// Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

func main() {
	for x := 10; x <= 100; x++ {
		rest := x % 4 
		fmt.Printf("O resto da divisão de %v por %v é %v\n", x, 4, rest)
	}
}
