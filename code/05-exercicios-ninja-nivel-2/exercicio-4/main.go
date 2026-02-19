package main

import "fmt"

// Crie um programa que:
// - Atribua um valor int a uma variável
// - Demonstre este valor em decimal, binário e hexadecimal
// - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// - Demonstre esta outra variável em decimal, binário e hexadecimal

func main() {
	x:= 10
	fmt.Printf("Em decimal:%d\n", x)
	fmt.Printf("Em binário:%#b\n", x)
	fmt.Printf("Em hexadecimal:%#x\n", x)

	y:= x << 1
	fmt.Printf("Em decimal:%d\n", y)
	fmt.Printf("Em binário:%#b\n", y)
	fmt.Printf("Em hexadecimal:%#x\n", y)
}
