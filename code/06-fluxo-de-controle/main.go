package main

import "fmt"

func main() {
	// for x := 0; x < 10; x++ {fmt.Println(x)}; -> for loop can be modernized using range over int
	for i := range 10 {
		fmt.Println(i)
	}

}
