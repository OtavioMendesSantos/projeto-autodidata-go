package main

import "fmt"

// Crie um programa que utilize a declaração switch, sem switch statement especificado.

func main() {
	age := 17
	switch {
	case age >= 18:
		fmt.Println("maior de idade, pode dirigir")
	case (age < 18 && age >= 16):
		fmt.Println("menor de idade, pode dirigir")
	default:
		fmt.Println("pode n ta doido ?")
	}
}
