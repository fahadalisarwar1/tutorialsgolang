package main

import "fmt"

// Add intro here
// A defer statement defers the execution of a function until the surrounding function returns.
// usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.
// Deferred calls are executed even when the function panics:
func main() {
	defer fmt.Println("World") // ensure that some action is performed even if everything doesnt go well
	// means for ensuring something, for example closing the file if we are writing to it
	// If there are several deferred function calls, they are executed in last-in-first-out order.	LIFO stack


	panic("Stop")
	fmt.Println("Hello")
}
