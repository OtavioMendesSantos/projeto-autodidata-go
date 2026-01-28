package main

import "fmt"

var globalVariable = "Essa variável tem escopo de pacote"
// Variável declarada em um code block é undefined em outro
// Para variáveis com uma abrangência maior, package level scope, utilizamos var
// Funciona em qualquer lugar

func main() {
	// := | Declara ao menos uma nova variável e atribui valor | Se usa apenas na primeira vez criando a variável
	// =  | Apenas reatribui valor a variável existente        | Se usa apenas caso a variável já foi declarada

	x := 10                  // Declara int x = 10
	y := 20                  // Declara int y = 20
	z := x + y               // z = 30 (int)
	xIsGreaterThanY := x > y // xIsGreaterThanY = false (bool)

	fmt.Printf("%d %v, %T\n", z, z, z)
	// %d = formato de int, %s = formata string, %v = valor padrão, \n = quebra linha

	fmt.Println(xIsGreaterThanY)

	otherFunc();
}
