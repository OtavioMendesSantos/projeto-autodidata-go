package main

import "fmt"

// Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
// - ==
// - !=
// - <=
// - <
// - =
// Demonstre estes valores.

func main() {
	a := (5 == 10)
	a = (5 == 10)
	b := (5 != 10)
	c := (5 <= 10)
	d := (5 >= 10)
	e := (5 < 10)
	f := (5 > 10)
	fmt.Printf("a: %v\nb: %v\nc: %v\nd: %v\ne: %v\nf: %v\n", a, b, c, d, e, f)
}
