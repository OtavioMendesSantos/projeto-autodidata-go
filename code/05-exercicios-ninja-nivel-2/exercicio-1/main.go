package main

import "fmt"

// Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

func main() {
	x := 10
	fmt.Printf("Em decimal:%d\n", x)
	fmt.Printf("Em binário:%#b\n", x)
	fmt.Printf("Em hexadecimal:%#x\n", x)
}
