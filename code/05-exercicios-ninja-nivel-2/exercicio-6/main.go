package main

import "fmt"

// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// Demonstre estes valores.

const (
	_ = iota + 2026
	year1
	year2
	year3
	year4
)

func main() {
	fmt.Println(year1, year2, year3, year4)
}
