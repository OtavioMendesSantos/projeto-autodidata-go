package main

import "fmt"

// Põe na tela: todos os números de 1 a 10000.

func main() {
	fmt.Println("====== Iniciando contagem de 1 a 10000 ======")
	for i := 0; i <= 10000; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n====== Finalizando contagem======\n")
}
