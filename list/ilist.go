package list

import "errors"

type List interface {
	Clear()  // O(1)
	Length() // O(1)

	Append(e int)  // O(1) *Agrega al tail..
	Insert(e int)  // O(n)
	Remove()  // O(n)
	MoveToStart()  // O(1)
	MoveToEnd()  // O(1)
	MoveToPos(pos int)  // Array: O(1) LinkedList: O(n) 
	Prev()  // Array: O(1) LinkedList: O(n) 
	Next()  // O(1)
	CurrentPosition()  // Array: O(1) LinkedList: O(n) 
	CurrentElement()  // O(1)
}

// Errores comunes en las listas
var ErrEmptyList = errors.New("lista vacia")
var ErrFullList = errors.New("lista llena")
var ErrOutOfRangeList = errors.New("fuera de rango")