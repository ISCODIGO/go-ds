package list

import "errors"

type List interface {
	// Remove all contents from the list, so it is once again empty
	Clear()

	// Insert "it" at the current location
	// The client must ensure that the list's capacity is not exceeded
	Insert(it int) bool

	// Append "it" at the end of the list
	// The client must ensure that the list's capacity is not exceeded
	Append(it int) bool

	// Remove and return the current element
	Remove() (int, error)

	// Set the current position to the start of the list
	MoveToStart()

	// Set the current position to the end of the list
	MoveToEnd()

	// Move the current position one step left, no change if already at beginning
	Prev()

	// Move the current position one step right, no change if already at end
	Next()

	// Return the number of elements in the list
	Length() int

	// Return the position of the current element
	CurrPos() int

	// Set the current position to "pos"
	MoveToPos(pos int) bool

	// Return true if current position is at end of the list
	IsAtEnd() bool

	// Return the current element
	GetValue() (int, error)

	// Return true if the list is empty
	IsEmpty() bool
}

// Errores comunes en las listas
var ErrEmptyList = errors.New("lista vacia")
var ErrFullList = errors.New("lista llena")
var ErrOutOfRangeList = errors.New("fuera de rango")
var ErrNoSuchElement = errors.New("no such element")
