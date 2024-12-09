package main

import (
	"fmt"
	"go-ds/list"
)

func main() {
	// Create a new linked list
	linkedList := list.NewLinkedList()

	// Add some elements to the linked list
	linkedList.AddHead(1)
	linkedList.AddTail(2)
	linkedList.AddHead(3)

	// Print the elements of the linked list
	fmt.Println("Linked List:")
	fmt.Println(linkedList.ToString())
}
