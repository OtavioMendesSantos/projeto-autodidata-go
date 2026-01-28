package main

import "fmt"

var globalVariable string

// Variável declarada em um code block é undefined em outro
// Para variáveis com uma abrangência maior, package level scope, utilizamos var
// Funciona em qualquer lugar
// Caso seja inicializado com package scope e sem valor, a atribuição somente será feita no codeblock

var a int
var b float64
var c string
var d bool

func main() {
	// := | Declara ao menos uma nova variável e atribui valor | Se usa apenas na primeira vez criando a variável
	// =  | Apenas reatribui valor a variável existente        | Se usa apenas caso a variável já foi declarada

	globalVariable = "Essa variável tem escopo de pacote"

	x := 10 // Declara int x = 10
	y := 20 // Declara int y = 20
	x = 15  // atribui novo valor
	// x = 10.5 -> Error - Go é uma linguagem de tipagem estática, por isso é impossivel trocar o tipo
	// de uma variável, caso eu queira trocar o tipo, eu preciso declarar uma nova variável com o novo tipo

	z := x + y               // z = 30 (int)
	xIsGreaterThanY := x > y // xIsGreaterThanY = false (bool)

	fmt.Printf("%d %v, %T\n", z, z, z)
	// %d = formato de int, %s = formata string, %v = valor padrão, \n = quebra linha

	fmt.Println("x > y:", xIsGreaterThanY)

	// Go prioriza segurança de tipo acima de "comportamento esperado":
	// int / int -> int (sempre inteiro)
	// float / float -> float (sempre decimal)
	// Filosofia Go: "Se você quer float, diga explicitamente."
	part := 12.0
	total := 185.0
	percentagem := (part / total) * 100
	fmt.Printf("%.2f%% de %.0f é %.0f\n", percentagem, total, part)

	otherFunc()

	fmt.Println("Valores 0:")
	// Valores 0 aparecem quando uma variável é declarada e não é inicializada
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	//String literals
	literal:= "Oi \ncomo vai?\nespero que \"que\" tudo bem!"
	result := fmt.Sprint(literal)
	fmt.Println(result)
}
