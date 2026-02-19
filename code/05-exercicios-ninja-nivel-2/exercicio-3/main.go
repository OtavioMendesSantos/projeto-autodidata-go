package main

import "fmt"

//Crie constantes tipadas e não-tipadas.
// Demonstre seus valores.

func main() {
	const w int = 100
	const y = 300

	fmt.Printf("Const tipados: %v, %T\n", w, w)
	fmt.Printf("Const não tipados: %v, %T\n", y, y)
}
