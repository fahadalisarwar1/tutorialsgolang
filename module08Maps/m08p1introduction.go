package main

import "fmt"

func main() {
	fmt.Println("Maps introductions")
	// A map is a builtin type in Go which
	// associates a value to a key. The value can be retrieved using the corresponding key.

	// Syntax   mapVar = make(map[keyType]valueType)

	students := make(map[int]string)
	students[0] = "Fahad"
	students[1] = "Ali"
	students[2] = "Sarwar"

	fmt.Println(students)
	for key, value := range students {
		fmt.Println("ID ", key, " Name: ", value)

	}

	dictionary := map[string]int{
		"apples":  22,
		"oranges": 66,
	}

	// the values of the map can be directly retrived by
	fmt.Println(dictionary["apples"])

	// to check if a key is present in the map or not
	_, ok := dictionary["oranges"]
	fmt.Println(ok)

	// to add a new entry to the map
	dictionary["bananas"] = 23

	for fruit, quantity := range dictionary {
		fmt.Println("fruit: ", fruit, " quantity: ", quantity)

	}

	// to delete an entry we can use a builtin tool
	delete(dictionary, "bananas")
	fmt.Println(dictionary)

	// how to find the length of the map
	fmt.Println(len(dictionary))

	// Maps are reference types this means they are referenced by their memory address not by their value
	// copying them into a new variable and changing the value of the new variable will alter the origianl map too
	newFruits := dictionary
	newFruits["apples"] = 99
	for key, fruit := range newFruits {
		fmt.Println(key, " ", fruit)
	}

	for fruit, quantity := range dictionary {
		fmt.Println("fruit: ", fruit, " quantity: ", quantity)
		// this shows that the original dictionary has also been modified
	}
	// finally maps can't be equated
	// map1 != map2 directly, if you want ot compare two maps you should compare their values individually

}
