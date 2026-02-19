package main

import "fmt"

// Crie um programa que demonstre o funcionamento da declaração if.

func main() {
	x := 10
	y := 10
	if x > y {
		fmt.Println("x é maior q y")
	} else if x < y {
		fmt.Println("x é menor q y")
	} else {
		fmt.Println("x é igual a y")
	}
}
