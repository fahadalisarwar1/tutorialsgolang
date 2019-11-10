package main

import "fmt"

func main() {
	// fmt.Println("Nested Loops")
	// for i := 1; i <= 10; i++ {
	// 	for j := 5; j >= 0; j-- {
	// 		fmt.Printf("%d %d\n", i, j)
	// 	}
	// }
	n := 5
	for i := 0; i <= n; i++ {
		for j := 5; j >= 0; j-- {
			fmt.Printf("i = %d j = %d ,", i, j)
		}
		fmt.Println()
	}
	for {
		fmt.Println("HI")
	}
}
