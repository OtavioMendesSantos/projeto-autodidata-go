package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// loops()
	// desafio()
	condicional()
}

func loops() {
	// for x := 0; x < 10; x++ {fmt.Println(x)}; -> for loop can be modernized using range over int
	// Go não tem while

	total := 0
	for i := range 100 {
		for x := range 10 {
			fmt.Printf("Primeiro loop %v, Segundo loop %v\n", i, x)
			total++
		}
	}
	fmt.Printf("===== Finalizou com %v repetições =====\n", total)

	// Infinite loop
	// for {
	// 	fmt.Println("Executando...")
	// }

	x := 0
	y := 0

	fmt.Println("Iniciando processamento desnecessário...")
	for {
		if x >= 10 {
			break
		}
		random := rand.Intn(51) + 1
		y++
		if random > 50 {
			fmt.Printf("Entrando +1 para a conta (%v na tentativa %v)\n", x, y)
			x++
		}
	}
	fmt.Printf("===== Após %v tentativas, o for acabou =====\n", y)

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue // pula a iteração atual
		}
		fmt.Println(i)
	}
}

// Loops: Desafio surpresa utilizando ascii!
// Format printing:
// - Decimal %d
// - Hexadecimal %#x
// - Unicode %#U
// - Tab \t
// - Linha nova \n
// - Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.
func desafio() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v %U - %v\n", i, i, string(i))
		// o Go interpreta esse número como um code point Unicode e converte para o caractere correspondente.
		// O que está acontecendo tecnicamente
		// - Um int pode representar um rune
		// - rune é um alias para int32
		// - rune representa um Unicode code point
	}
}

func condicional() {
	x := 10
	if x > 5 {
		fmt.Println("Maior que 5")
	} else {
		fmt.Println("Menor que 5")
	}
	// if err := doSomething(); err != nil {
	// 	fmt.Println(err)
	// }

	switch {
	case x < 5:
		fmt.Println("Menor que 5")
	case x > 5:
		fmt.Println("Maior que 5")
	default:
		fmt.Println("Igual a 5")
	}

	fruta := "banana"
	switch fruta {
	case "maçã":
		fmt.Println("Maçã")
	case "banana":
		fmt.Println("Banana")
		fallthrough // cair através
	case "granola":
		fmt.Println("com granola é massa")
	default:
		fmt.Println("Fruta desconhecida")
	}

	equipe := "Azul"
	switch equipe {
	case "Azul", "Vermelho", "Amarelo":
		fmt.Println("Equipe Azul, Vermelho ou Amarelo")
	case "Verde":
		fmt.Println("Equipe Verde")
	default:
		fmt.Println("Equipe desconhecida")
	}
}
