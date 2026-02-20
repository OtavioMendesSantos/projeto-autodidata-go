package main

import "fmt"

func main() {
	// dadosArray()
	// dadosSlice()
	// fatiandoFatias()
	anexandoSlices()
}

// Arrays possuem um tamanho fixo e fazem parte do tipo
var x [5]int // tipo [5]int, tipos diferentes

func dadosArray() {
	x[0] = 10
	x[1] = 20
	// x[2] = começa com o valor zero (zero value) -> 0
	x[4] = 50
	// x[5] = 50 -> invalid argument: index 5 out of bounds [0:5]

	fmt.Printf("%T - %v\n", x, x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%v - %v\n", i, x[i])
	}

	y := [6]int{0, 1, 2, 3, 4, 5}
	for i, v := range y {
		fmt.Printf("%v - %v\n", i, v)
	}
}

// Slice é uma estrutura que referencia um array interno.
// Ele possui tamanho variável (len) e capacidade (cap).
func dadosSlice() {
	// Cria um slice com 4 elementos.
	// O array subjacente tem tamanho 4 neste caso.
	slice := []int{0, 1, 2, 3}

	// append adiciona elementos após o último índice válido (len-1).
	// Ele pode reutilizar o array interno ou criar um novo,
	// dependendo da capacidade disponível.
	// Por isso é necessário atribuir o retorno.
	slice = append(slice, 4, 5, 6)

	// Mesmo que a capacidade (cap) seja maior,
	// não é permitido acessar índices >= len(slice).
	// Exemplo: slice[10] = 10 -> panic: index out of range

	for i, v := range slice {
		fmt.Printf("%v - %v\n", i, v)
	}

	// make cria um slice com len=3 e cap=6.
	// O array interno tem 6 posições (todas inicializadas com zero).
	s := make([]int, 3, 6) // len=3 cap=6

	fmt.Printf("%v %T\n", s, s)

	// len -> quantidade de elementos acessíveis
	// cap -> tamanho total do array interno

	// Ao expandir até a capacidade,
	// passamos a enxergar todas as posições do array interno.
	fmt.Println(s[:cap(s)])
}

func fatiandoFatias() {
	sabores := []string{"chocolate", "morango", "limão", "baunilha"}
	fmt.Println(sabores[1:3])
	fmt.Println(sabores[:2])
	fmt.Println(sabores[2:])
	fmt.Println(sabores[:])
	sabores = append(sabores, "abacaxi")
	fmt.Println(append(sabores[:1], sabores[3:]...))
}

func anexandoSlices() {
	// append -> anexa itens a uma slice - faz parte do package built-in
	sabores := []string{"chocolate", "morango", "limão", "baunilha"}
	sabores = append(sabores, "abacaxi")
	fmt.Println(sabores)
}
