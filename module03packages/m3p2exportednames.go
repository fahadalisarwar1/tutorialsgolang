package main

import "fmt"

func main() {
	fmt.Println("Exported packages")
	// one thing to note here is that exported packages should have a name wiht a Capital letter otherwise they wont
	// be exportable this is a go convention that must be strictly followed

	// also for documentation the exported packages should have a comment on top of that that will Also start with the same
	// name as the exported function which will tell what this function does.
}
