package main

import "fmt"

var x bool

func main() {
	fmt.Println("Zero value of bool:", x)
	x = true
	fmt.Println("New value of bool:", x)
	x = 10 < 100
	y := 10 == 100
	z := 10 > 100
	w := 10 != 100
	fmt.Println(x, y, z, w)
}
