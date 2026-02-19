package main

import "fmt"

// Crie um loop utilizando a sintaxe: for {}
// Utilize-o para demonstrar os anos desde que você nasceu.

func main() {
	rightYear := 2005
	year := 1995
	for {
		year++
		isRightYear := rightYear == year
		fmt.Println(year, isRightYear)
		if isRightYear {
			break
		}
	}
}
