package main

import "fmt"

// Crie um programa que utilize a declaração switch, onde o switch
// statement seja uma variável do tipo string com identificador "esporteFavorito".

func main() {
	esporteFavorito := "futebol de botão"
	switch esporteFavorito {
	case "futebol":
		fmt.Println("futebol é o esporte favorito")
	case "futebol de botão":
		fmt.Println("futebol de botão é o esporte favorito")
	default:
		fmt.Println("vish, tem o esport nao")
	}
}
