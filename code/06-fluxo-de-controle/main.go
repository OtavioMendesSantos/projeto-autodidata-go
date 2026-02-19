package main

import "fmt"

func main() {
	// for x := 0; x < 10; x++ {fmt.Println(x)}; -> for loop can be modernized using range over int
	total := 0
	for i := range 10 {
		for x := range 10 {
			fmt.Printf("Primeiro loop %v, Segundo loop %v\n", i, x)
			total++
		}
	}
	fmt.Printf("===== Finalizou com %v repetições =====\n", total)
}
