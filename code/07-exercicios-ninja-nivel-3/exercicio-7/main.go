package main

import "fmt"

//Utilizando a solução anterior, adicione as opções else if e else.

func main() {
	x := 10
	y := 10
	if x > y {
		fmt.Println("x é maior q y")
	} else if x < y {
		fmt.Println("x é menor q y")
	} else {
		fmt.Println("x é igual a y")
	}
}
