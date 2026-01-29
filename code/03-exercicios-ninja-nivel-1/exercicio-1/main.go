package main

import "fmt"

// Na prática: exercício #1

// Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".

// 42
// "James Bond"
// true
// Agora demonstre os valores nestas variáveis, com:

// Uma única declaração print
// Múltiplas declarações print
// Solução: https://play.golang.org/p/yYXnWXIQNa

func main() {
	x, y, z := 42, "James Bond", true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
