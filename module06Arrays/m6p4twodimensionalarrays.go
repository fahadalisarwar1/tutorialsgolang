package main

import "fmt"

func main() {
	fmt.Println("2D arrays ")
	var grid [2][3]int // an array of 2 rows and three columns
	grid[1][2] = 33    // by default all values are assigned zero, so we modify the 2nd row: hence 1 index and
	// 3rd column hence 2nd index since indexes start at zero

	// TODO looping over 2D arrays

	fmt.Println(grid)
}
