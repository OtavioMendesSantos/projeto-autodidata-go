package main

import "fmt"

// Na prática: exercício #5
// Utilizando a solução do exercício anterior:
// Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
// Na função main:
// Isto já deve estar feito:
// Demonstre o valor da variável "x"
// Demonstre o tipo da variável "x"
// Atribua 42 à variável "x" utilizando o operador "="
// Demonstre o valor da variável "x"
// Agora faça tambem:
// Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
// Demonstre o valor de "y"
// Demonstre o tipo de "y"
// Solução: https://play.golang.org/p/uq81T_fasP
// Na prática: exercício #6
// Prova!
// Link: https://goo.gl/forms/s9y91iVSPvA4iahj1
// Se você deu pausa e fez todos os exercícios anteriores você mesmo, e só viu a resposta depois... e se você der pausa agora e fizer a prova inteira por conta própria, e só assistir as respostas depois... sabe o que isso quer dizer? Que você é ninja. Ninja nível 1. Tá no caminho certo pra ser um developer ninja mestre.

type CustomType int

var x CustomType
var y int

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%v %T\n", x, x)

	y = int(x)
	fmt.Printf("%v %T\n", y, y)
	
}
