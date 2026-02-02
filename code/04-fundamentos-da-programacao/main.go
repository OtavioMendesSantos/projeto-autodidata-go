package main

import (
	"fmt"
	"runtime"
)

var x bool

// Tipos Inteiros
// - uint8: Conjunto de todos os inteiros sem sinal de 8 bits (0 a 255).
// - uint16: Conjunto de todos os inteiros sem sinal de 16 bits (0 a 65.535).
// - int8: Conjunto de todos os inteiros com sinal de 8 bits (-128 a 127).
// - int16: Conjunto de todos os inteiros com sinal de 16 bits (-32.768 a 32.767).
// - int64: Conjunto de todos os inteiros com sinal de 64 bits (faixa longa: -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807).

// Tipos Float e Complexos
// - float32: Conjunto de todos os números de ponto flutuante IEEE-754 de 32 bits.
// - complex64: Conjunto de todos os números complexos com partes real e imaginária float32.
// - byte: Aliás para int8.
// - rune: Aliás para int32.​

func main() {
	// valuesBool()
	// valuesByte()
	valuesInt()
	// myComputer()
}

func valuesBool() {
	fmt.Println("Zero value of bool:", x)
	x = true
	fmt.Println("New value of bool:", x)
	x = 10 < 100
	y := 10 == 100
	z := 10 > 100
	w := 10 != 100
	fmt.Println(x, y, z, w)
}

func valuesByte() {
	a := "e"
	b := "é"
	c := "是"
	fmt.Printf("%v\t%v\t%v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v\t%v\t%v\n", d, e, f)
}

func valuesInt() {
	var x uint16
	// x = 65536 -> cannot use 65536 (untyped int constant) as uint16 value in assignment (overflows)
	// Go detecta overflow na COMPILAÇÃO para atribuições literais

	x = 65535
	fmt.Printf("%v, %T\n", x, x)
	
	x++ // Incrementa: 65535 + 1 = 65536 → OVERFLOW! Volta para 0 (wrap around)
	fmt.Printf("%v, %T\n", x, x)
	
	x++ // Incrementa novamente: 0 + 1 = 1
	fmt.Printf("%v, %T\n", x, x)
	
	y := 10.0
	fmt.Printf("%v, %T\n", y, y)
}

func myComputer() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
}
