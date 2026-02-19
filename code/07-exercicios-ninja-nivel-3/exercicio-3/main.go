package main

import "fmt"

// Crie um loop utilizando a sintaxe: for condition {}
// Utilize-o para demonstrar os anos desde que você nasceu.

func main() {
	rightYear := 2005
	year := 1995
	for rightYear != year {
		year++
		fmt.Println(year, rightYear == year)
	}
}
