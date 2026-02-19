package main

import "fmt"

// Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
// Por exemplo: 65 U+0041 'A' U+0041 'A' U+0041 'A' 66 U+0042 'B' U+0042 'B' U+0042 'B' ...e por aí vai.

func main() {
	for x := 1; x <= 26; x++ {
		i := 64 + x
		letter := string(i)
		fmt.Printf("=== Imprimindo letra %v (3 vezes) ===\n", letter)
		for y := 0; y < 3; y++ {
			fmt.Printf("%v.%v - Letra %v %U \n", x, y, letter, i)
		}
	}
}
