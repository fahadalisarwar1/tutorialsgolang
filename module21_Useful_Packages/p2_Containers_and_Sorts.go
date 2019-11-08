package main

import (
	"container/list"
	"fmt"
)

func main(){
	// in addition to lists and maps, Go has several more types
	// Containers and lists
	// The container/list package implements a doubly linked list

	// each node in the linked list contains a value and a pointer to the next node.
	// In case of doubly linked link list there will be a pointer to previous node too.
	var ll list.List
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)
	for ptr := ll.Front(); ptr != nil; ptr = ptr.Next(){
		fmt.Println(ptr.Value.(int))
	}

}
