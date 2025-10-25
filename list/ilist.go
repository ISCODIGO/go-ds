package list

import "errors"

type List interface {
	Clear()        // O(1)
	Size() int     // O(1)
	IsEmpty() bool // O(1)

	AddHead(e int) *Node      // O(1)
	AddTail(e int) *Node      // O(1)
	RemoveHead() (int, error) // O(1)
	RemoveTail() (int, error) // O(n)
	Remove(e int) error       // O(n)
}

// Errores comunes en las listas
var ErrEmptyList = errors.New("lista vacia")
var ErrFullList = errors.New("lista llena")
var ErrOutOfRangeList = errors.New("fuera de rango")
