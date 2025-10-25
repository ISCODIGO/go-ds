package main

import (
	"fmt"
	"go-ds/list"
)

func main() {
	// Create a new linked list
	linkedList := list.NewLinkedList()

	// Test new interface methods
	fmt.Println("=== Testing New List Interface ===")

	// Append elements
	linkedList.Append(10)
	linkedList.Append(20)
	linkedList.Append(42)

	fmt.Printf("List after appending: %s\n", linkedList.ToString())
	fmt.Printf("Length: %d\n", linkedList.Length())

	// Navigate through the list
	linkedList.MoveToStart()
	val, _ := linkedList.GetValue()
	fmt.Printf("Current element at start: %d\n", val)

	linkedList.Next()
	val, _ = linkedList.GetValue()
	fmt.Printf("Next element: %d\n", val)

	linkedList.MoveToEnd()
	val, _ = linkedList.GetValue()
	fmt.Printf("Element at end: %d\n", val)

	// Insert at current position
	linkedList.MoveToPos(1)
	linkedList.Insert(99)
	fmt.Printf("After inserting at position 1: %s\n", linkedList.ToString())

	// Remove current element
	removed, _ := linkedList.Remove()
	fmt.Printf("Removed element: %d\n", removed)
	fmt.Printf("List after removal: %s\n", linkedList.ToString())

	fmt.Printf("Current position: %d\n", linkedList.CurrPos())
	fmt.Printf("Is at end: %t\n", linkedList.IsAtEnd())

	// Test compatibility with old methods
	fmt.Println("\n=== Testing Compatibility ===")
	linkedList.AddHead(5)
	linkedList.AddTail(100)
	fmt.Printf("Final list: %s\n", linkedList.ToString())
}
